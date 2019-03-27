package main

import "fmt"

func main() {
	nextPos1 := getPositiveIng()

	fmt.Println(nextPos1())
	fmt.Println(nextPos1())
	fmt.Println(nextPos1())
	fmt.Println(getPositiveIng()())

	fmt.Println()
	nextPos2 := getPositiveIng()
	fmt.Println(nextPos2())

	fmt.Printf("'%T' '%T' \n", nextPos1, nextPos1())
	fmt.Printf("%x %x \n", &nextPos1, &nextPos2)

}
func getPositiveIng() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
