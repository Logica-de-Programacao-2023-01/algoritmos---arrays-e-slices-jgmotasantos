package main

import "fmt"

func main() {
	//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
	var tamanho int
	fmt.Println("Qual o tamanho do seu slice?")
	fmt.Scan(&tamanho)

	numeros := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Println("Quais sao seus elemetos ?")
		fmt.Scan(&numeros[i])
	}
	fmt.Printf("Seu slice eh: ", numeros)
	var soma int
	for _, numeros := range numeros {
		soma += numeros
	}
	fmt.Println("A soma dos elementos foi:", soma)
}
