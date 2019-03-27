package main

import (
	"fmt"
)

var i = 0

func main() {
	defer func() {
		errMsg := recover()
		fmt.Println("panic error")
		fmt.Println(errMsg)
	}()

	var x map[string]int
	x["key"] = 10
	fmt.Println(x)

	panic("BOO!!!")
	fmt.Println("this. line won't be reached")
}
