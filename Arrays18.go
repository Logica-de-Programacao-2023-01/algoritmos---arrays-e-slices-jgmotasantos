package main

import "fmt"

func main() {
	var n int

	fmt.Print("D	igite o valor de n: ")
	fmt.Scan(&n)

	var primos []int

	for i := 2; i <= n; i++ {
		var primo bool = true
		for x := 2; x < i; x++ {
			if i%x == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}

	fmt.Println(primos)
}
