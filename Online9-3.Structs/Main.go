package main

import "fmt"

func main() {
	player1 := struct {
		name, sport string
		age         int
	}{"player1 name", "player1 sport", 30}

	fmt.Println("player1 =", player1)

	player2 := struct {
		name, sport string
		age         int
	}{
		age:   21,
		sport: "player2 sport",
		name:  "player2 name",
	}

	fmt.Println("player2 =", player2)
}
