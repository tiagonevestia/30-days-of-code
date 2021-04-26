package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		inputs = append(inputs, input)
	}

	listSize := inputs[0]
	T, err := strconv.ParseInt(listSize, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	agendaTelefonica := make(map[string]int64, T)
	sharedName := inputs[T+1:]
	nomesENumeros := inputs[1 : T+1]

	for _, val := range nomesENumeros {
		str := strings.Split(val, " ")
		nome := str[0]
		nums := str[1]

		number, err := strconv.ParseInt(nums, 10, 64)
		checkError(err)

		agendaTelefonica[nome] = number
	}

	for i, val := range sharedName {
		if numero, ok := agendaTelefonica[sharedName[i]]; !ok {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%d\n", val, numero)
		}
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
