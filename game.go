package main

import "fmt"


func main() {
	player := 0
	matches := 21

	for matches >= 0 {
		var move int
		fmt.Printf("Player %d, it's your turn. There are %d matches left, choose how many matches to take (1-3): ", player, matches)
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("Wrong user inpud")
			return
		}
		if move > 0 && move <= 3 {
			matches -= move	
		}
		
		if matches <= 0 {
			fmt.Printf("Player %d has lost the game. :( \n", player)
			return

		}

		player = (player + 1) % 2
		
	}
}
