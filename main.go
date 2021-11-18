package main

import (
	"log"
	"os"

	"github.com/WiredSound/goboy/media"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH  = 160
	SCREEN_HEIGHT = 144
)

func main() {
	if len(os.Args) != 2 {
		println("Please specify a single command-line argument to indicate the game ROM to load")
		return
	}

	path := os.Args[1]

	log.Printf("Will open ROM at path: %s", path)

	sdlKeyMap := map[sdl.Keycode]media.Button{
		sdl.K_ESCAPE:    media.BUTTON_QUIT,
		sdl.K_UP:        media.BUTTON_UP,
		sdl.K_DOWN:      media.BUTTON_DOWN,
		sdl.K_LEFT:      media.BUTTON_LEFT,
		sdl.K_RIGHT:     media.BUTTON_RIGHT,
		sdl.K_x:         media.BUTTON_A,
		sdl.K_z:         media.BUTTON_B,
		sdl.K_BACKSPACE: media.BUTTON_SELECT,
		sdl.K_RETURN:    media.BUTTON_START,
	}

	context, err := media.NewSDL2("GoBoy", SCREEN_WIDTH, SCREEN_HEIGHT, sdlKeyMap)
	if err != nil {
		panic("Failed to initialise SDL2 context")
	}
	log.Printf("Created multimedia context with window dimensions: %dx%d", SCREEN_WIDTH, SCREEN_HEIGHT)

	defer func() {
		context.Destroy()
		log.Printf("Destroyed multimedia context")
	}()

	log.Printf("Beginning game loop...")

	for buttons := map[media.Button]bool{}; !buttons[media.BUTTON_QUIT]; buttons = context.Update() {
		log.Println(buttons)
	}
}
