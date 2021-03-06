package main

import (
	"log"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	for k := range robotgo.Keycode {
		robotgo.EventHook(hook.KeyDown, []string{k}, func(e hook.Event) {
			//robotgo.TypeStr(string(e.Keychar))
			//fmt.Printf(string(e.Keychar))
			log.Println(e.String())
		})
	}

	log.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
