package main

import (
	"fmt"
)

func main() {
	fmt.Println("=>", factiorial(3))
	fmt.Println("=>", factiorial(4))
	fmt.Println("=>", factiorial(7))
}
func factiorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(n, " ")
	return n * factiorial(n-1)
}
