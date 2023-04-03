package main

import "fmt"

func main() {
	numeros := []float64{4.43, 5.63, 3.33, 8.321}
	numero1 := numeros[0]
	numero2 := numeros[1]
	numero3 := numeros[2]
	numero4 := numeros[3]
	Produto := numero1 * numero2 * numero3 * numero4
	fmt.Println("O produto entre os elementos eh:", Produto)
}
