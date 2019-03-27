package main

import (
	"fmt"
	"strings"

	athletes "./types"
)

func changeAthleteName1(p athletes.Player) {
	p.Name = "123456"
}
func changeAthleteName2(p *athletes.Player) {
	p.Name = "123456"
	p.Country = strings.ToUpper(p.Country)
}
func main() {
	// player1 := athletes.Player{"AS", "MMA", 43, athletes.Info{"Brazil", "Black"}}
	info1 := athletes.Info{Country: "Brazil", HairColor: "Black"}
	player1 := athletes.Player{Name: "AS", Sport: "MMA", Age: 43, Info: info1}
	fmt.Println("(1) player1:", player1)

	changeAthleteName1(player1)
	fmt.Println("(2) player1:", player1)

	changeAthleteName2(&player1)
	fmt.Println("(3) player1:", player1)

	fmt.Println("(4) player1:", *player1.ToLowercase())
}
