package main

import (
	"fmt"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/go-vgo/robotgo"
)

func autoClicker(running *bool) {
	for {
		if *running {
			robotgo.Click("left", true)
			fmt.Println("Simulated double-click mouse event!")

			fmt.Println("next click after 3-sec")
			time.Sleep(3 * time.Second)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	running := false

	go autoClicker(&running)

	keyboard.Listen(
		func(key keys.Key) (stop bool, err error) {
			if key.String() == "s" {
				fmt.Println("Hitting 's' STARTS the auto-click")
				running = true
			} else if key.String() == "t" {
				fmt.Println("Hitting 't' TERMINATES the auto-click")
				running = false
			} else if key.String() == "e" {
				fmt.Println("Ended Script by pressing 'e'")
				return true, nil
			} else {
				fmt.Println("\r" + key.String())
				fmt.Println("Unrelated key pressed!")
			}
			return false, nil
		})
}
