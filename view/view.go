package view

import (
	"time"

	"fmt"

	"github.com/kassybas/cellchi/bl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	W_WIDTH  int32 = 800
	W_HEIGHT int32 = 600
)

func Start(window *sdl.Window, surface *sdl.Surface) {

	for bl.Items == nil {
		fmt.Println("Items still not initialized")
		//TODO: correct racecondition
		time.Sleep(time.Second)
	}
	mainLoop(surface, window)
}

func mainLoop(surface *sdl.Surface, window *sdl.Window) {
	clear := sdl.Rect{X: 0, Y: 0, W: W_WIDTH, H: W_HEIGHT}
	running := true
	//var event sdl.Event

	for running {
		//get events TODO:work with channels
		// for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		// 	switch t := event.(type) {
		// 	case *sdl.QuitEvent:
		// 		running = false
		// 		fmt.Println("Quit")
		// 	case *sdl.KeyDownEvent:
		// 		fmt.Printf("Sym:%c\tMod:%d\n", t.Keysym.Sym, t.Keysym.Mod)
		// 	}
		// }

		surface.FillRect(&clear, 0)
		for i := range bl.Items {
			itemRect := bl.Items[i].Poll()
			rect := sdl.Rect{
				X: itemRect.PosX,
				Y: itemRect.PosY,
				W: itemRect.Size,
				H: itemRect.Size,
			}
			surface.FillRect(&rect, itemRect.Color)
		}
		sdl.Delay(10)
		window.UpdateSurface()
	}
}
