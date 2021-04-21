package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	// O scanner lê a entrada e separa-a em linhas ou palavras, é a maneira
	// mais fácil de processar entradas fornecidas naturalmente em linhas.
	// O scanner lê a entrada-padrão do programa.
	scanner := bufio.NewScanner(os.Stdin)

	var entradaInt int64
	var entradaFloat float64
	var entradaString string
	inputs := []string{}

	// Cada chamada ao método .Scan() lê a próxima linha e remove o caracter
	// de quebra de linha do final.
	for scanner.Scan() {
		// O resultado de .Scan() vamos obter por meio da chamada .Text()
		input := scanner.Text()
		// A função Scan devolve true se houver uma linha, e false quando
		// não houver mais dados de entrada
		if len(input) == 0 {
			break
		}
		// Adicionando uma string ao vetor
		inputs = append(inputs, input)
	}

	entradaInt, _ = strconv.ParseInt(inputs[0], 10, 64)
	entradaFloat, _ = strconv.ParseFloat(inputs[1], 64)
	entradaString = inputs[2]

	fmt.Println(i + uint64(entradaInt))
	fmt.Printf("%.1f\n", (d + entradaFloat))
	fmt.Println(s + entradaString)
}
