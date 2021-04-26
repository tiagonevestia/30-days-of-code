package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	// b := fmt.Sprintf("%b", n)
	binary := strconv.FormatInt(int64(n), 2)
	binaryArr := strings.Split(binary, "0")
	// var acc int
	// var total int
	var max float64
	for _, v := range binaryArr {
		// fmt.Println(v)
		num, _ := strconv.ParseFloat(v, 64)
		max = math.Max(num, max)
		// if num == 1 {
		// 	acc += 1
		// 	if acc > total {
		// 		total = acc
		// 	}
		// } else {
		// 	acc = 0
		// }
	}
	// fmt.Println(baseConvertBinary(n))
	// fmt.Printf("%b", n)
	// fmt.Printf("%b\n", n)
	// fmt.Println(acc)
	// fmt.Println(total)
	consecutiveOnes := strconv.Itoa(int(max))
	fmt.Println(len(consecutiveOnes))
}

// LIVRO: Estrutura de dados e algoritmos com JS: Loiane
func baseConvertBinary(num int32) string {
	remStacks := []float64{}
	var number = num
	var baseString string

	for number > 0 {
		rem := float64(int32(number) % 2)
		remStacks = append(remStacks, rem)
		number = int32(float64((number / 2)))
	}

	for len(remStacks) != 0 {
		baseString += fmt.Sprintf("%v", remStacks[len(remStacks)-1])
		remStacks = remStacks[:len(remStacks)-1] // .pop
	}
	return baseString
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
