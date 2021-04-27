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
	numberOfSwaps := 0
	arrNew := []int64{}

	arrTemp := strings.Split(readLine(reader), " ")

	for _, v := range arrTemp {
		num, _ := strconv.ParseInt(v, 10, 64)
		arrNew = append(arrNew, num)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arrNew[j] > arrNew[j+1] {
				swap(arrNew, j, j+1)
				numberOfSwaps++
			}
		}
		if numberOfSwaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numberOfSwaps)
	fmt.Printf("First Element: %d\n", arrNew[0])
	fmt.Printf("Last Element: %d\n", arrNew[n-1])
}

func swap(arr []int64, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
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
