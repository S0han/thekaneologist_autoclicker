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
			fmt.Println("Simulated mouse-click event!")
			time.Sleep(10 * time.Second)
			fmt.Println("next click after 10-sec")
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
			if key.String() == "S" {
				fmt.Println("Hitting 'S' STARTS the auto-click")
				running = true
			} else if key.String() == "T" {
				fmt.Println("Hitting 'T' TERMINATES the auto-click")
				running = false
			} else if key.String() == "E" {
				fmt.Println("Ended Script by pressing 'E'")
				return true, nil
			} else {
				fmt.Println("\r" + key.String())
				fmt.Println("Unrelated key pressed")
			}
			return false, nil
		})
}
