package main

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func main() {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Enter {
			fmt.Println("Enter key pressed script will terminate.")
			return true, nil
		}

		fmt.Println("\r" + key.String())
		return false, nil
	})
}
