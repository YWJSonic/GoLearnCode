package main

import "fmt"

type generalInfo struct {
	conutry, hairColor string
}

type player struct {
	name, sport string
	age         int
	generalInfo
}

func main() {
	var player1 player
	player1.name = "way"
	player1.age = 57
	player1.sport = "Hocyey"
	player1.conutry = "Canada"
	player1.hairColor = "Brown"

	fmt.Println("player1:", player1)
}
