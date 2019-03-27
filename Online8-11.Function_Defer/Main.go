package main

import (
	"fmt"
)

var i = 0

func main() {
	fmt.Println("\t\t\ti(m1)=", i)
	testDefer()
	fmt.Println("\t\t\ti(m2)=", i)
}
func testDefer() {
	fmt.Println("\t\t\ti(t1)=", i)

	defer closeFiles()
	defer closeDBConnections()

	fmt.Println("\t\t\ti(t2)=", i)
	doSomething()
	fmt.Println("\t\t\ti(t3)=", i)
}
func closeFiles() {
	fmt.Println("closeFiles\t\ti(f1)=", i)
	i = 1
}
func closeDBConnections() {
	fmt.Println("closeDBConnections\ti(f2)=", i)
	i += 2
}
func doSomething() {
	fmt.Println("doSomething\t\ti(f3)=", i)
	i = 3
}
