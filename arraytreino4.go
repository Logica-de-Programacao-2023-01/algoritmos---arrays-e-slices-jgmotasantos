package main

import "fmt"

func main() {
	arrayFloats := []float64{2.232132, 4.429836, 9.943653945639, 5.634192469, 7.7777777, 4.72148214}
	fmt.Print("Seus 6 elementos são", arrayFloats)
	elemento1 := float64(2.232132)
	elemento2 := float64(4.429836)
	elemento3 := float64(9.943653945639)
	elemento4 := float64(5.634192469)
	elemento5 := float64(7.7777777)
	elemento6 := float64(4.72148214)

	resultado := float64(elemento1 + elemento2 + elemento3 + elemento4 + elemento5 + elemento6)

	fmt.Println("A medias dos 6 elementos é:")
	fmt.Println(resultado / 6)

}

//Crie um Array de floats com 6 elementos e calcule a média dos valores armazenados no
//Array.