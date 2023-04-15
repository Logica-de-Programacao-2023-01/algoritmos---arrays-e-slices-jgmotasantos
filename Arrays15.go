package main

import "fmt"

func main() {

	arr := [10]float64{1.2, 2.3, 2.7, 3.5, 4.9, 5.1, 5.8, 7.5, 9.2, 10.5}
	slice := []float64{}

	for _, i := range arr {
		if i > 5 {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)
}

