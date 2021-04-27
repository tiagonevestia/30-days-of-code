package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)
	arrNew := []string{}

	for i := 0; i < n; i++ {

		num, _ := strconv.ParseInt(readLine(reader), 10, 64)

		if primo(num) {
			arrNew = append(arrNew, "Prime")
		} else {
			arrNew = append(arrNew, "Not prime")
		}
	}

	for _, v := range arrNew {
		fmt.Println(v)
	}

}

func primo(n int64) bool {
	var d int64
	if n <= 1 {
		return false
	}
	for d = 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
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
