package main

import (
	"fmt"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

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

func autoClickerGUI() {
	a := app.New()
	w := a.NewWindow("")

	start_ac_key_label := widget.NewLabel("Start:")
	stop_ac_key_label := widget.NewLabel("Stop:")
	interval_label := widget.NewLabel("Interval:")
	start_ac_key_input := widget.NewEntry()
	stop_ac_key_input := widget.NewEntry()
	interval_input := widget.NewEntry()
	start_button := widget.NewButton("START", func() {})
	stop_button := widget.NewButton("STOP", func() {})

	w.SetContent(container.NewVBox(
		start_ac_key_label,
		start_ac_key_input,
		stop_ac_key_label,
		stop_ac_key_input,
		interval_label,
		interval_input,
		start_button,
		stop_button,
	))

	w.ShowAndRun()
}

func main() {
	running := false

	go autoClicker(&running)

	go keyboard.Listen(
		func(key keys.Key) (stop bool, err error) {
			if key.String() == "s" {
				fmt.Println("Hitting 's' STARTS the auto-click")
				running = true
			} else if key.String() == "t" {
				fmt.Println("Hitting 't' TERMINATES the auto-click")
				running = false
			} else {
				fmt.Println("\r" + key.String())
				fmt.Println("Unrelated key pressed!")
			}
			return false, nil
		})

	autoClickerGUI()
}
