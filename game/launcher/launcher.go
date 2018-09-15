package launcher

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/rodnover55/go-football/game"
	"github.com/rodnover55/go-football/game/scenes"
	"github.com/rodnover55/go-football/game/window"
	"golang.org/x/image/colornames"
)

type Launcher struct {
	scene game.Scene
}

func New() *Launcher {
	return &Launcher{}
}

func (launcher *Launcher) Run() {
	pixelgl.Run(func() {
		win := window.New()

		if launcher.scene == nil {
			launcher.scene = scenes.MakeGame(win)
		}

		for !win.Closed() {
			win.Clear(colornames.Black)
			launcher.scene.Update()
			win.Update()
		}
	})
}
