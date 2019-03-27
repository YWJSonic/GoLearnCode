package main

import "fmt"

type generalInfo struct {
	conutry, hairColor string
}

type player struct {
	name, sport string
	age         int
	info        generalInfo
}

func main() {
	var player1 player
	player1.name = "way"
	player1.age = 57
	player1.sport = "Hocyey"
	player1.info.conutry = "Canada"
	player1.info.hairColor = "Brown"

	fmt.Println("player1:", player1)

	info2 := generalInfo{
		conutry:   "USA",
		hairColor: "black",
	}

	player2 := player{
		name:  "Le",
		sport: "Bacsk",
		age:   36,
		info:  info2,
	}

	fmt.Println("player2:", player2)
}
