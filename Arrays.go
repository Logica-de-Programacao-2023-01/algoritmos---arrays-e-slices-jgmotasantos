package main

import "fmt"

func main() {
	numeros := [3]int{1, 2, 3}
	numero1 := numeros[0]
	numero2 := numeros[1]
	numero3 := numeros[2]
	Resultado := numero1 + numero2 + numero3
	fmt.Println("A soma dos tres numeros eh:", Resultado)

}
