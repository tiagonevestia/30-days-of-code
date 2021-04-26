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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32
	var arrNew []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
		// exemplo do livro
		// for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		// 	arr[i], arr[j] = arr[j], arr[i]
		// }
	}

	for x := len(arr) - 1; x > 0; x-- {
		arrNew = append(arrNew, (arr[x]))
	}

	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(arrNew), " ", " ", -1), "[]"))

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
