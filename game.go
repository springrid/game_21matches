package main

import (
	"fmt"
)

func play(players []Agents, startingPlayer int) int {
	player := startingPlayer
	matches := 21

	for matches >= 0 {
		var move int
		move = players[player].act(matches)
		move++
		matches -= move

		if matches <= 0 {
			// fmt.Printf("Player %d has won the game. :) \n", (player + 1) % 2)
			return (player + 1) % 2
		}

		player = (player + 1) % 2
	}
	return -1
}

func playN(players []Agents, nGames int, plotOn bool) {
	winners := []int{}
	for i := 0; i < nGames; i++ {
		winner := play(players, i%2)

		if winner == 0 {
			players[0].feedback(1)
			players[1].feedback(-1)
		} else {
			players[0].feedback(-1)
			players[1].feedback(1)
		}
		players[0].newEpisode()
		players[1].newEpisode()

		winners = append(winners, winner)
	}

	counter := make(map[int]int)
	for _, row := range winners {
		counter[row]++
	}
	fmt.Println("Winners (n times): ", counter)
	if plotOn {
		plotResult(winners)
	}
}
