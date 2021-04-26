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
	length := int(n)

	groups := 0

	for i := 0; i < length; i++ {
		pivot := s[i]
		diffChar := -1

		for j := i + 1; j < length; j++ {
			curChar := s[j]
			if pivot == curChar {
				if diffChar == -1 || (j-diffChar) == (diffChar-i) {
					groups++
				}
			} else {
				if diffChar == -1 {
					diffChar = j
				} else {
					break
				}
			}
		}
	}

	return int64(groups)
}

func main() {

	f, err := os.Open("./SpecialStringAgain/case2.txt")

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

	scanner.Scan()
	out := scanner.Text()
	temp, err := strconv.ParseInt(out, 10, 64)
	checkError(err)
	expectedOutput := int32(temp)

	if expectedOutput == int32(result) {
		fmt.Println("Passou")
	} else {
		fmt.Println("Rejeitado")
	}

	fmt.Println("Encontrou =>", result, "esperava =>", expectedOutput)
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
