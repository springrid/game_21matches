import {
	"fmt"
}

func main() {
        if err := run(); err != nil {
                fmt.Printf("Error %s running game", err)
        }
}
