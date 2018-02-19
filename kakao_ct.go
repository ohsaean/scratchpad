package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/ohsaean/gogpd/lib"
	"time"
	"strings"
	"math"
)

func main() {
	//

	//ret := secretMap(5,
	//	[]int{9, 20, 28, 18, 11},
	//	[]int{30, 1, 21, 17, 28},
	//)
	//fmt.Println(ret)
	//
	//ret = secretMap(6,
	//	[]int{46, 33, 33, 22, 31, 50},
	//	[]int{27, 56, 19, 14, 14, 10},
	//)
	//fmt.Println(ret)
	//
	//sum := dartGame("1D2S3T*")
	//fmt.Println(sum)

	//ret := GetCacheExecutionTime(3, 	[]string{"Jeju","Pangyo","Seoul","NewYork","LA","Jeju","Pangyo","Seoul","NewYork","LA"})
	ret := GetCacheExecutionTime(3, 	[]string{"Jeju","Pangyo","Seoul","Jeju","Pangyo","Seoul", "Jeju","Pangyo","Seoul"})
	fmt.Println(ret)
}

func secretMap(n int, arr1 []int, arr2 []int) []string {
	//1. 비밀 지도(난이도: 하)
	//
	//네오는 평소 프로도가 비상금을 숨겨놓는 장소를 알려줄 비밀지도를 손에 넣었다. 그런데 이 비밀지도는 숫자로 암호화되어 있어 위치를 확인하기 위해서는 암호를 해독해야 한다. 다행히 지도 암호를 해독할 방법을 적어놓은 메모도 함께 발견했다.
	//	지도는 한 변의 길이가 n인 정사각형 배열 형태로, 각 칸은 “공백”(“ “) 또는 “벽”(“#”) 두 종류로 이루어져 있다.
	//	전체 지도는 두 장의 지도를 겹쳐서 얻을 수 있다. 각각 “지도 1”과 “지도 2”라고 하자. 지도 1 또는 지도 2 중 어느 하나라도 벽인 부분은 전체 지도에서도 벽이다. 지도 1과 지도 2에서 모두 공백인 부분은 전체 지도에서도 공백이다.
	//“지도 1”과 “지도 2”는 각각 정수 배열로 암호화되어 있다.
	//	암호화된 배열은 지도의 각 가로줄에서 벽 부분을 1, 공백 부분을 0으로 부호화했을 때 얻어지는 이진수에 해당하는 값의 배열이다.

	var completeMap []int
	for i := 0; i < n; i++ {
		overlappedLine := arr1[i] | arr2[i]
		completeMap = append(completeMap, overlappedLine)
	}

	var convertedMap []string

	for _, line := range completeMap {
		lineStr := ""

		for i := n; 0 < i; i-- {
			mask := 1 << uint32(i-1)
			masked := line & mask
			ret := masked >> uint32(i-1)
			if ret == 1 {
				lineStr += "#"
			} else {
				lineStr += " "
			}
		}

		convertedMap = append(convertedMap, lineStr)
	}

	return convertedMap
}

func dartGame(input string) int {

	const (
		Single = "S"
		Double = "D"
		Triple = "T"
		Star   = "*"
		Ahcha  = "#"
	)

	//calcSum := 0

	//finishCalc := true

	sum := 0
	doubleBonus := 1
	minusBonus := 1
	power := 1
	number := -1
	minusBonusLimit := 0
	doubleBonusLimit := 0
	fmt.Println(len(input))
	for i := len(input) - 1; 0 <= i; i-- {
		str := string(input[i])

		switch str {
		case Single:
			power = 1
			number = -1
			break
		case Double:
			power = 2
			number = -1
			break
		case Triple:
			power = 3
			number = -1
			break
		case Star:
			doubleBonus *= 2
			number = -1
			doubleBonusLimit = 2
			break
		case Ahcha:
			minusBonusLimit = 1
			minusBonus *= -1
			number = -1
			break
		default:
			// 숫자
			if number == 0 {
				number = 10
			} else {
				number = lib.Atoi(str)
			}
			break
		}

		if number > 0 {
			shoot := 1
			for j := 0; j < power; j++ {
				shoot *= number
			}

			if doubleBonusLimit > 0 {
				shoot *= doubleBonus
			}

			if minusBonusLimit > 0 {
				shoot *= minusBonus
			}

			sum += shoot

			if doubleBonus > 1 {
				doubleBonusLimit--
			}
			if minusBonus < 0 {
				minusBonusLimit--
			}
		}

	}

	return sum
}

type CacheItem struct {
	Key       string
	Value     string
	cacheTime int64
	hits      int
}

type CacheServer struct {
	MaxBucketSize int
	Items         map[string]*CacheItem
}

func makeTimestamp() int64 {
	return time.Now().UnixNano()
}

func NewCacheServer(size int) *CacheServer {
	return &CacheServer{
		MaxBucketSize: size,
		Items:         map[string]*CacheItem{},
	}
}

func (c *CacheServer) Get(key string) *CacheItem {
	if value, ok := c.Items[key]; ok {
		value.cacheTime = makeTimestamp()
		value.hits++
		return value
	}
	return nil
}

func (c *CacheServer) Set(key string, item *CacheItem) {
	if len(c.Items) >= c.MaxBucketSize {
		c.eviction()
	}

	if len(c.Items) > c.MaxBucketSize {
		return
	}

	c.Items[key] = item
	item.cacheTime = makeTimestamp()
	item.hits = 0
	time.Sleep(time.Nanosecond)
}

func (c *CacheServer) eviction() {
	// LRU
	var minCacheTime int64
	minCacheTime = math.MaxInt64

	leastUsedKey := ""

	for key, val := range c.Items {
		if val.cacheTime < minCacheTime {
			minCacheTime = val.cacheTime
			leastUsedKey = key
		}
	}

	if len(leastUsedKey) > 0 {
		delete(c.Items, leastUsedKey) // 삭제
	}
}

func GetCacheExecutionTime(cacheSize int, cities []string) int {
	if cacheSize < 0 || cacheSize > 30 {
		log.Fatal("invalid cache size")
	}

	citiesSize := len(cities)
	if citiesSize < 0 {
		log.Fatal("empty cities")
	}

	if citiesSize > 100000 {
		log.Fatal("max cities")
	}

	cache := NewCacheServer(cacheSize)

	var executeTime int
	executeTime = 0
	for _, val := range cities {
		city := strings.ToLower(val)
		if cache.Get(city) == nil {
			cache.Set(city, &CacheItem{
				Key:city,
				Value:city,
			})
			executeTime += 5
		} else {
			executeTime += 1
		}
	}

	return executeTime
}
