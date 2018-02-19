package main

import (
	"os"
	"bufio"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readLineByte(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return ln, err
}

func findLine(fileName string, targetLineNum int) (ret []byte) {
	file, err := os.Open(fileName)
	checkErr(err)

	fr := bufio.NewReader(file)

	lineBytes, readError := readLineByte(fr)
	checkErr(readError)

	// first line..
	lineNum := 1
	if len(lineBytes) != 0 && targetLineNum == lineNum {
		ret = lineBytes
	}

	// other
	for readError == nil {
		lineNum++
		lineBytes, readError = readLineByte(fr)
		if len(lineBytes) != 0 && targetLineNum == lineNum {
			ret = lineBytes
		}
	}

	return ret
}

func main() {
	fileName := "findline_test.log"
	fmt.Println(string(findLine(fileName, 1)))
}