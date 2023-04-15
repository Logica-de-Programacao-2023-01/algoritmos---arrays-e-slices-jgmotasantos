package main

import "fmt"

func main() {
	//Variaveis
	array := [7]int{0, 0, 0, 0, 0, 0, 0}
	var primeiraCasa int
	var ultimaCasa int

	//Aplica  as respostas as variaveis
	fmt.Println("Qual numero deve ser adicionado na primeira casa?")
	fmt.Scan(&primeiraCasa)
	fmt.Println("Qual numero deve ser adicionado na ultima casa?")
	fmt.Scan(&ultimaCasa)

	//Atribui o valor indicado pelo usuario a primeira e ultima casa respectivamente
	array[0] = primeiraCasa
	array[6] = ultimaCasa

	//Imprime o array
	fmt.Println(array)

}
