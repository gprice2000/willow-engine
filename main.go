package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	//TODO add startup logic (what should happen when engine starts?)

	//Initialize subsystems (video,audio,etc)
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		panic(err)
	}
	defer sdl.Quit()	

	//Attempt to create a window
	sdl.CreateWindow()
}