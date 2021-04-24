package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	groups := int(n)

	length := int(n) - 1
	sameChar := true

	for i := 0; i < length; i++ {
		if s[i] == s[i+1] {
			groups++
		} else {
			sameChar = false
		}

		if i > 0 && i < length && s[i-1] == s[i+1] {
			groups++
		}
	}

	if sameChar {
		groups++
	}

	return int64(groups)
}

func main() {

	f, err := os.Open("input-example.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 10, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	scanner.Scan()
	read := scanner.Text()
	nTemp, err := strconv.ParseInt(read, 10, 64)
	checkError(err)
	n := int32(nTemp)

	if scanner.Err() != nil {
		fmt.Printf(" > Failed with error %v\n", scanner.Err())
	}

	scanner.Scan()
	s := strings.TrimRight(string(scanner.Text()), "\r\n")

	result := substrCount(n, s)

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
