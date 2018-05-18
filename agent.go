import "math/rand"


type agentRandom struct {
	actions := 2
}

func (a agentRandom) Act() int {
	return rand.Intn(len(actions))
}

type agentUser struct{}

func (a agentUser) act(state) int {
	var move int
	fmt.Printf("State: ", state)
	_, err := fmt.Scan(&move)
	if err != nil || move < 1 || move > 3 {
		fmt.Println("Wrong user inpud")
		return
	}
	return move - 1
}
