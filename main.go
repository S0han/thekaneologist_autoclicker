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

		// logs the key presses
		fmt.Println("\r" + key.String())

		//terminate script with this button
		if key.String() == "P" {
			fmt.Println("STOP key pressed!")
			return true, nil
		}

		fmt.Println("Un-related key pressed")
		return false, nil
	})
}
