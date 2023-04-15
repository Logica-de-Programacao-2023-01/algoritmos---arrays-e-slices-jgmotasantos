package main

import "fmt"

func main() {
	array := [5]int{1, 4, 6, 7, 9}
	var slice []int

	for _, i := range array {
		if i%3 == 0 {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)
}
