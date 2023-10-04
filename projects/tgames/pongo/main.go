package main

import (
	"log"
	"os"
	"pongo/ball"
	"pongo/game"
	_ "pongo/game"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal("%+v", err)
	}

	ball := ball.Ball{
		X: 1,
		Y: 1,
		Xspeed: 1,
		Yspeed: 1,
	}

	game := game.Game{
		Screen: screen,
		Ball: ball,
	}
	go game.Run()

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}
