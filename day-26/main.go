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

	const (
		taxaDias = 15
		taxaMes  = 500
		taxaAno  = 10000
	)

	dataRetorno := strings.Split(readLine(reader), " ")
	dataVencimento := strings.Split(readLine(reader), " ")

	diaRetorno, _ := strconv.ParseInt(dataRetorno[0], 10, 64)
	mesRetorno, _ := strconv.ParseInt(dataRetorno[1], 10, 64)
	anoRetorno, _ := strconv.ParseInt(dataRetorno[2], 10, 64)

	diaVencimento, _ := strconv.ParseInt(dataVencimento[0], 10, 64)
	mesVencimento, _ := strconv.ParseInt(dataVencimento[1], 10, 64)
	anoVencimento, _ := strconv.ParseInt(dataVencimento[2], 10, 64)

	if anoRetorno > anoVencimento {
		fmt.Printf("%d\n", taxaAno)
		return
	} else if anoRetorno < anoVencimento {
		fmt.Println(0)
		return
	}

	if mesRetorno > mesVencimento {
		fmt.Printf("%d\n", taxaMes*(mesRetorno-mesVencimento))
		return
	}

	if diaRetorno > diaVencimento {
		fmt.Printf("%d\n", taxaDias*(diaRetorno-diaVencimento))
		return
	}

	fmt.Println(0)

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
