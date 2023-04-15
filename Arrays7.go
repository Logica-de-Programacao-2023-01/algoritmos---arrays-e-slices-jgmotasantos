package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 5, 9, 14, 23}

	var x int
	fmt.Println("Qual numero deseja adicionar?")
	fmt.Scan(&x)

	encontrado := false
	for _, i := range slice {
		if x == i {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Println("Seu numero ja esta presente na lista!")
	} else {
		slice = append(slice, x)
		sort.Ints(slice)
		fmt.Println(slice)
	}

}
