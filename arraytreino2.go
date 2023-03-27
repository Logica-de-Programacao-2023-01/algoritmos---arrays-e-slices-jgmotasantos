package main

import "fmt"

func main() {
	elementos := []string{}
	fmt.Print("Antes de adicionar seu string é:")
	fmt.Println(elementos)
	elementos = append(elementos, "João,Maria,José")
	fmt.Print("Depois de adicionar os nomes seu string é:")
	fmt.Println(elementos)
}

//Crie um Slice de strings vazio, adicione os nomes "João"
//"Maria" e "José" a ele e imprima
//o Slice.
