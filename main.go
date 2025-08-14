package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/go-vgo/robotgo"
)

func autoClicker(running *bool, interval *time.Duration, mu *sync.Mutex) {
	for {
		mu.Lock()
		copy_running := *running
		mu.Unlock()
		if copy_running {

			robotgo.Click("left", true)
			fmt.Println("Simulated mouse click event!")

			fmt.Printf("next click after %v\n", interval)
			time.Sleep(*interval)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func autoClickerGUI(running *bool, s *string, t *string, interval *time.Duration, mu *sync.Mutex) {
	a := app.New()
	w := a.NewWindow("")

	//start input
	start_ac_key_label := widget.NewLabel("Start:")
	start_ac_key_input := widget.NewEntry()
	start_ac_key_input.OnChanged = func(text string) {
		if len(text) == 1 {
			mu.Lock()
			*s = text
			mu.Unlock()
			return
		}
	}

	//stop input
	stop_ac_key_label := widget.NewLabel("Stop:")
	stop_ac_key_input := widget.NewEntry()
	stop_ac_key_input.OnChanged = func(text string) {
		if len(text) == 1 {
			mu.Lock()
			*t = text
			mu.Unlock()
			return
		}
	}

	//time interval
	interval_label := widget.NewLabel("Interval:")
	interval_input := widget.NewEntry()
	interval_input.OnChanged = func(text string) {
		//convert the string input to integer
		sec, err := strconv.Atoi(text)
		if err != nil {
			return
		}

		mu.Lock()
		*interval = time.Duration(sec) * time.Second
		mu.Unlock()
	}

	//start & stop buttons
	start_button := widget.NewButton("START", func() {
		mu.Lock()
		fmt.Println("Hitting 'START' button STARTS the auto-click!")
		*running = true
		mu.Unlock()
	})
	stop_button := widget.NewButton("STOP", func() {
		mu.Lock()
		fmt.Println("Hitting 'STOP' button TERMINATES the auto-click!")
		*running = false
		mu.Unlock()
	})

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
	var running bool = false
	var mu sync.Mutex
	var interval time.Duration = 3 * time.Second
	s := ""
	t := ""

	go autoClicker(&running, &interval, &mu)

	go keyboard.Listen(
		func(key keys.Key) (stop bool, err error) {
			mu.Lock()
			copy_s := s
			copy_t := t
			mu.Unlock()

			if key.String() == copy_s {
				mu.Lock()
				fmt.Printf("Hitting '%v' STARTS the auto-click\n", copy_s)
				running = true
				mu.Unlock()
			} else if key.String() == copy_t {
				mu.Lock()
				fmt.Printf("Hitting '%v' TERMINATES the auto-click\n", copy_t)
				running = false
				mu.Unlock()
			} else {
				fmt.Println("\r" + key.String())
				fmt.Println("Unrelated key pressed!")
			}
			return false, nil
		})
	autoClickerGUI(&running, &s, &t, &interval, &mu)
}
