package main

import (
	"fmt"

	"github.com/kassybas/cellchi/bl"
	"github.com/kassybas/cellchi/view"
	"github.com/veandco/go-sdl2/sdl"
)

func initWindowAndSurface(title string, w, h int) (*sdl.Window, *sdl.Surface, error) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, err
	}

	surface, err := window.GetSurface()
	if err != nil {
		window.Destroy()
		return nil, nil, err
	}

	return window, surface, nil
}

func main() {
	//Window has to be in main thread
	window, surface, err := initWindowAndSurface("02", int(view.W_WIDTH), int(view.W_HEIGHT))
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	go bl.Start()
	go view.Start(window, surface)
	var input string
	fmt.Scanln(&input)
	sdl.Quit()
}
