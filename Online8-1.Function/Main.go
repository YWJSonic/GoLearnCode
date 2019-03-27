package main

import (
	"fmt"
)

func main() {
	prinMsg("on fun print")

	showMsg := func(msg string) {
		fmt.Println(msg)
	}

	showMsg("tmp func print")
	fmt.Printf("%T\n", showMsg)

	func(msg string) {
		fmt.Println(msg)
	}("one time print")
}
func prinMsg(msg string) {
	fmt.Println(msg)
}
