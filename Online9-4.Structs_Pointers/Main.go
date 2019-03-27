package main

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"
)

type employee struct {
	name, job    string
	lastLoggedIn string
	DOB          date.Date
}

func main() {
	var emp employee

	emp.name = "Nick"
	emp.job = "Go Programmer"
	emp.lastLoggedIn = time.Now().Format(time.RFC850)

	emp.DOB = date.Today()
	fmt.Println(emp)

	//-------------------------------------------------------
	p := &emp

	p.name = "Jack"
	emp.job = "Go Expert"
	fmt.Println(*p)

	fmt.Printf("%x %x", &emp.name, &p.name)

}
