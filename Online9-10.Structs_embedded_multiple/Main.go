package main

import "fmt"

type movie struct {
	name  string
	actor string
}

func (m movie) fullInfo() string {
	return m.name + " " + m.actor
}

func main() {
	m1 := movie{"For Gump", "Tom Hanks"}
	m2 := movie{"the godfather", "Marlon Brando"}

	fmt.Println(m1.fullInfo())
	fmt.Println(m2.fullInfo())
}
