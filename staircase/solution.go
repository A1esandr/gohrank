package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func staircase(n int32) {
	m := int(n)
	for i := 1; i <= m; i++ {
		var sb strings.Builder
		sb.Grow(m)
		for {
			if sb.Len() == m-i {
				break
			}
			sb.WriteByte(' ')
		}
		for {
			if sb.Len() == m {
				break
			}
			sb.WriteByte('#')
		}
		fmt.Println(sb.String())
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
