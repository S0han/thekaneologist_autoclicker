package main

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/go-vgo/robotgo"
)

func main() {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Enter {
			robotgo.Click("left", true)
			fmt.Println("Hitting Enter triggers a simulated mouse-click event!")
			return false, nil
		}

		fmt.Println("\r" + key.String())
		return true, nil
	})
}
