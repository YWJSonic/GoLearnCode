package main

import "fmt"

type player struct {
	string
	int
}

func main() {
	player1 := player{"ali", 99}
	fmt.Println("player1:", player1)
	fmt.Printf("playerint=%d player1.string=%s\n", player1.int, player1.string)

	player2 := player{
		int:    36,
		string: "Fed",
	}
	fmt.Println("player2:", player2)
}
