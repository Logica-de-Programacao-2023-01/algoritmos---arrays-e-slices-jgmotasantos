package main

import "fmt"

func main() {
	n := 5

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Digite os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Println("Digite os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array2[i])
	}
	array3 := make([]int, n)
	for i := 0; i < n; i++ {
		array3[i] = array1[i] + array2[i]
	}

	fmt.Println("O terciero array que é a soma dos dois primeiros é:")
	fmt.Println(array3)
}
