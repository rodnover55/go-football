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

		// TODO: Добавить обработку эвентов:
		// - завершение игры (с перезагрузкой)
		// - загрузка игры
		if launcher.scene == nil {
			// TODO: Грузить сцену с надписью Loading, а в фоне игру. По окончанию загрузки запускать эту сцену
			launcher.scene = scenes.NewGame(win)
		}

		for !win.Closed() {
			win.Clear(colornames.Black)
			launcher.scene.Update()
			win.Update()
		}
	})
}
