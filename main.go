package main

import "fmt"

func main() {
	mc := &monteCarlo{actions: 3, learn: true, eps: 0.00}
	mc.init()
	p := &perfect{}

	// MonteCarlo vs. Perfect -> MC learns to beat perfect in 50% of all games :)
	fmt.Println("\nMonte Carlo vs Perfect")
	players := []Agents{mc, p}
	playN(players, 1000, false)

	// freeze learning of first MonteCarlo agent
	mc.learn = false

	// New MonteCarlo agent learns to beat the first (frozen) one
	fmt.Println("\nOld Monte Carlo (frozen) vs new Monte Carlo")
	mc2 := &monteCarlo{actions: 3, learn: true, eps: 0.01}
	mc2.init()
	players = []Agents{mc, mc2}
	playN(players, 10000, false)

	// MonteCarlo beats random agent
	fmt.Println("\nPre-trained Monte Carlo vs Random")
	r := &random{actions: 3}
	players = []Agents{mc, r}
	playN(players, 10000, false)

	// User vs MonteCarlo
	// u := &user{}
	// players = []Agents{mc, u}
	// playN(players, 3, false)
}
