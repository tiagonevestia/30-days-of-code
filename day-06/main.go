package main

import (
	"fmt"
)

func main() {

	var T int
	var s string
	var str string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&s)
		resPar := []string{}
		resImpar := []string{}

		for i, val := range s {
			if i%2 == 0 {
				resPar = append(resPar, string(val))
			} else {
				resImpar = append(resImpar, string(val))
			}
		}
		for _, valPar := range resPar {
			str += valPar
		}
		str += " "
		for _, valImp := range resImpar {
			str += valImp
		}
		str += "\n"
	}

	fmt.Println(str)
}
