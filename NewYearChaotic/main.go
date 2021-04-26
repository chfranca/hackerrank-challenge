package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the substrCount function below.
func minimumBribes(q []int32) {
	length, bribes, gains, chaotic := len(q), 0, 0, false

	for i := 0; i < length; i++ {
		// fmt.Println(q[i], i, length)
		value := int(q[i] - 1)
		if value > i {
			distance := int(math.Abs(float64(value - i)))
			if distance > 2 {
				chaotic = true
				break
			} else {
				gains += distance
				bribes += distance
			}
		} else if value < i {
			distance := int(math.Abs(float64(i - value)))
			if (gains - distance) >= 0 {
				gains -= distance
				if distance > 2 {
					bribes += (distance - 3)
				}
			} else {
				chaotic = true
				break
			}
		}
	}

	if chaotic {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribes)
	}
}

func main() {

	f, err := os.Open("./NewYearChaotic/case2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 10, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	scanner.Scan()
	read := scanner.Text()
	tTemp, err := strconv.ParseInt(strings.TrimSpace(read), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		scanner.Scan()
		read := scanner.Text()
		nTemp, err := strconv.ParseInt(strings.TrimSpace(read), 10, 64)
		checkError(err)
		n := int32(nTemp)

		scanner.Scan()
		qTemp := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
