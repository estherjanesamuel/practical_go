package main

import (
	"log"
	"os"
	"time"

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
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)

	go run (screen, defStyle)
	for {

		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func run(screen tcell.Screen, defStyle tcell.Style)  {
	x := 0
	for {
		screen.Clear()
		screen.SetContent(x, 10, 'H', nil, defStyle)
		screen.SetContent(x+1, 10, 'i', nil, defStyle)
		screen.SetContent(x+2, 10, '!', nil, defStyle)

		screen.Show()
		x++

		time.Sleep(40 * time.Millisecond)
	}
}
