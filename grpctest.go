package main

import (
	"flag"
	"fmt"

	"runtime"
	"smon_raid_cli_client/lib"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	//address     = "localhost:50051"
	defaultName = "world"
)

type RequesterStats struct {
	TotRespSize    int64
	TotDuration    time.Duration
	MinRequestTime time.Duration
	MaxRequestTime time.Duration
	NumRequests    int
	NumErrs        int
}

var statsAggregator chan *RequesterStats
var goroutines int = 4
var duration int = 10
var wg *sync.WaitGroup
var reqStart time.Time
var host string = ""

func init() {
	flag.IntVar(&goroutines, "c", 10, "Number of goroutines to use (concurrent connections)")
	flag.IntVar(&duration, "d", 10, "Duration of test in seconds")
	flag.StringVar(&host, "host", "", "Host Address")
	flag.Parse()

	maxProcs := runtime.NumCPU() + goroutines

	fmt.Printf("Running %vs test @ %v\nUse %v CPU(s) and %v goroutine(s) running concurrently\n", duration, host, maxProcs, goroutines)

	reqStart = time.Now()
	statsAggregator = make(chan *RequesterStats, goroutines)
	wg = new(sync.WaitGroup) // 대기 그룹 생성
	wg.Add(goroutines)
	runtime.GOMAXPROCS(maxProcs)
}

func main() {

	responders := 0
	aggStats := RequesterStats{MinRequestTime: time.Minute}

	for i := 0; i < goroutines; i++ {
		go request()
	}

	go func() {
		for {
			select {
			case stats := <-statsAggregator:
				aggStats.NumErrs += stats.NumErrs
				aggStats.NumRequests += stats.NumRequests
				aggStats.TotRespSize += stats.TotRespSize
				aggStats.TotDuration += stats.TotDuration
				aggStats.MaxRequestTime = lib.MaxDuration(aggStats.MaxRequestTime, stats.MaxRequestTime)
				aggStats.MinRequestTime = lib.MinDuration(aggStats.MinRequestTime, stats.MinRequestTime)
				responders++
			}
		}
	}()

	wg.Wait()
	if aggStats.NumRequests == 0 {
		fmt.Println("Error: No statistics collected / no requests found")
		return
	}

	avgThreadDur := aggStats.TotDuration / time.Duration(responders) //need to average the aggregated duration
	reqRate := float64(aggStats.NumRequests) / avgThreadDur.Seconds()
	avgReqTime := aggStats.TotDuration / time.Duration(aggStats.NumRequests)
	bytesRate := float64(aggStats.TotRespSize) / avgThreadDur.Seconds()
	fmt.Printf("%v requests avg in %v, %v read\n", aggStats.NumRequests, avgThreadDur, lib.ByteSize{float64(aggStats.TotRespSize)})
	fmt.Printf("Requests/sec:\t\t%.2f\nTransfer/sec:\t\t%v\nAvg Req Time:\t\t%v\n", reqRate, lib.ByteSize{bytesRate}, avgReqTime)
	fmt.Printf("Fastest Request:\t%v\n", aggStats.MinRequestTime)
	fmt.Printf("Slowest Request:\t%v\n", aggStats.MaxRequestTime)
	fmt.Printf("Number of Errors:\t%v\n", aggStats.NumErrs)
}

func request() {

	stats := &RequesterStats{MinRequestTime: time.Minute}

	defer wg.Done()

	// Set up a connection to the server.
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}

	if err != nil {
		fmt.Println(err)
		stats.NumErrs++
		statsAggregator <- stats
		return
	}

	for time.Since(reqStart).Seconds() <= float64(duration) {

		start := time.Now()

		respSize := SayHello(conn)

		reqDur := time.Since(start)
		if respSize > 0 {
			stats.TotRespSize += int64(respSize)
			stats.TotDuration += time.Since(start)
			stats.MaxRequestTime = lib.MaxDuration(reqDur, stats.MaxRequestTime)
			stats.MinRequestTime = lib.MinDuration(reqDur, stats.MinRequestTime)
			stats.NumRequests++
		} else {
			stats.NumErrs++
		}
	}

	statsAggregator <- stats
}

func SayHello(conn *grpc.ClientConn) int {
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	msg := &pb.HelloRequest{Name: name}

	r, err := c.SayHello(context.Background(), msg)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	//fmt.Printf("Greeting: %s", r.Message)
	return len(r.Message)
}
