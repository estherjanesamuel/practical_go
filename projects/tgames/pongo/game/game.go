package game

import (
	"pongo/ball"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Ball ball.Ball
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	g.Screen.SetStyle(defStyle)

	for {
		width, height := g.Screen.Size()
		g.Ball.CheckEdges(width,height)
		g.Screen.Clear()
		g.Ball.Update()
		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		time.Sleep(45 * time.Millisecond)
		g.Screen.Show()
	}
}
