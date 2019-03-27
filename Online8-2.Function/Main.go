package main

import (
	"fmt"
)

func main() {
	factor := 2
	mult := func(i, j int) int {
		fmt.Println("factor1 ", factor)
		factor = 3
		fmt.Println("factor2 ", factor)
		factor := 4
		fmt.Println("factor3 ", factor)
		return i * j * factor
	}
	fmt.Println("factor4 ", factor)
	fmt.Println(mult(3, 4))
	fmt.Println("factor5 ", factor)
}
