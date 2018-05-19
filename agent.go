package main

import (
	"fmt"
	"math/rand"
)

// Agents implements a common interface for all agents
type Agents interface {
	act(int) int
	feedback(float64)
	newEpisode()
}

// random is a agent who chooses his actions at random (1-3)
type random struct {
	actions int
}

func (r random) act(state int) int {
	return rand.Intn(r.actions)
}

func (r *random) feedback(reward float64) {}

func (r *random) newEpisode() {}

// user player is human input
type user struct {
	name string
}

func (u *user) act(state int) int {
	var move int
	fmt.Printf("State: %d \n%s, it's your turn. Input: ", state, u.name)

	_, err := fmt.Scan(&move)
	if err != nil || move < 1 || move > 3 {
		fmt.Println("Wrong user inpud")
		return 0
	}
	return move - 1
}

func (u *user) feedback(reward float64) {}

func (u *user) newEpisode() {}

// monteCarlo is an agent playing and learning according to the moMonte Carlo algorithm
type monteCarlo struct {
	actions int
	learn   bool
	eps     float32

	reward map[int]map[int]float64
	count  map[int]map[int]float64

	stateActionBuffer [][2]int
}

func argMax(args map[int]float64) int {
	max := -10000000.0
	maxInd := 0

	for key, arg := range args {
		if arg > max {
			max = arg
			maxInd = key
		}
	}
	return maxInd
}

func (mc *monteCarlo) act(state int) int {
	var action int

	if mc.learn && rand.Float32() < mc.eps {
		action = rand.Intn(mc.actions)
	} else {
		action = argMax(mapDivision(mc.reward[state], mc.count[state]))
	}

	mc.stateActionBuffer = append(mc.stateActionBuffer, [2]int{state, action})
	return action
}

func (mc *monteCarlo) feedback(reward float64) {
	if mc.learn == true {
		for _, stateAction := range mc.stateActionBuffer {
			state := stateAction[0]
			action := stateAction[1]

			if _, ok := mc.reward[state]; !ok {
				m1 := make(map[int]float64)
				m2 := make(map[int]float64)

				for i := 0; i < mc.actions; i++ {
					m1[i] = 0
					m2[i] = 0.00001
				}

				m1[action] = reward
				m2[action] = 1

				mc.reward[state] = m1
				mc.count[state] = m2
			} else {
				mc.reward[state][action] += reward
				mc.count[state][action]++
			}
		}
	}
}

func (mc *monteCarlo) newEpisode() {
	mc.stateActionBuffer = [][2]int{}
}

func (mc *monteCarlo) init() {
	mc.reward = make(map[int]map[int]float64)
	mc.count = make(map[int]map[int]float64)
}

type perfect struct{}

func (p perfect) act(state int) int {
	if (state-1)%4 == 0 {
		return 0
	}

	action := state - (4*((state-1)/4) + 1)
	if action >= 1 && action <= 3 {
		return action - 1
	}

	return -1
}

func (p *perfect) feedback(reward float64) {}

func (p *perfect) newEpisode() {}
