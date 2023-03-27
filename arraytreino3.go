package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	numeros = append(numeros[:2], numeros[3:]...)
	fmt.Println("Sem o terceiro termo o slice ficaria:", numeros)

	//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e
	//imprima o Slice resultante.

}
