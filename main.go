package main

import (
	"log"
	"os"

	"github.com/WiredSound/goboy/gameboy"
	"github.com/WiredSound/goboy/media"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	GAMEBOY_SCREEN_WIDTH  = 160
	GAMEBOY_SCREEN_HEIGHT = 144
)

func main() {
	if len(os.Args) != 2 {
		println("Please specify a single command-line argument to indicate the game ROM to load")
		return
	}

	path := os.Args[1]

	log.Printf("Will open ROM at path: %s", path)

	sdlColourMap := map[media.Colour][3]uint8{
		media.COLOUR_BLACK: {0, 0, 0},
		media.COLOUR_DARK:  {100, 100, 100},
		media.COLOUR_LIGHT: {200, 200, 200},
		media.COLOUR_WHITE: {255, 255, 255},
	}

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

	context, err := media.NewSDL2("GoBoy", GAMEBOY_SCREEN_WIDTH, GAMEBOY_SCREEN_HEIGHT, 6, sdlColourMap, sdlKeyMap)
	if err != nil {
		panic("Failed to initialise SDL2 context")
	}
	windowWidth, windowHeight := context.WindowSize()
	log.Printf("Created multimedia context with window dimensions: %dx%d", windowWidth, windowHeight)

	defer func() {
		context.Destroy()
		log.Printf("Destroyed multimedia context")
	}()

	context.Plot(GAMEBOY_SCREEN_WIDTH-1, GAMEBOY_SCREEN_HEIGHT-1, media.COLOUR_LIGHT)
	context.Present()

	log.Printf("Beginning game loop")

	gb := gameboy.New()

	for buttons := map[media.Button]bool{}; !buttons[media.BUTTON_QUIT]; buttons = context.Update() {
		gb.Update(context, buttons)
	}
}
