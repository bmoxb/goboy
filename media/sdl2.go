package media

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type sdlContext struct {
	window            *sdl.Window
	renderer          *sdl.Renderer
	screenScaleFactor int32
	colourMap         map[Colour][3]uint8
	keyButtonMap      map[sdl.Keycode]Button
	buttonsDown       map[Button]bool
}

func NewSDL2(title string, width, height, screenScaleFactor int32, colourMap map[Colour][3]uint8, keyButtonMap map[sdl.Keycode]Button) (sdlContext, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return sdlContext{}, err
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width*screenScaleFactor, height*screenScaleFactor, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Quit()
		return sdlContext{}, err
	}
	log.Println("SDL2 window created")

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		sdl.Quit()
		return sdlContext{}, err
	}
	log.Println("SDL2 renderer created")

	buttonsDown := map[Button]bool{}
	return sdlContext{window, renderer, screenScaleFactor, colourMap, keyButtonMap, buttonsDown}, nil
}

func (c sdlContext) Update() map[Button]bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return map[Button]bool{BUTTON_QUIT: true}

		case *sdl.KeyboardEvent:
			key := t.Keysym.Sym
			button, present := c.keyButtonMap[key]

			if present {
				if t.State == sdl.PRESSED {
					c.buttonsDown[button] = true
				} else if t.State == sdl.RELEASED {
					delete(c.buttonsDown, button)
				}
			}
		}
	}

	return c.buttonsDown
}

func (c sdlContext) Plot(x, y int32, col Colour) {
	rgb := c.colourMap[col]
	c.renderer.SetDrawColor(rgb[0], rgb[1], rgb[2], 255)

	rect := sdl.Rect{X: x * c.screenScaleFactor, Y: y * c.screenScaleFactor, W: c.screenScaleFactor, H: c.screenScaleFactor}

	c.renderer.FillRect(&rect)
}

func (c sdlContext) Present() {
	c.renderer.Present()
}

func (c sdlContext) SetWindowTitle(title string) {
	c.window.SetTitle(title)
}

func (c sdlContext) WindowSize() (int32, int32) {
	return c.window.GetSize()
}

func (c sdlContext) Destroy() {
	c.renderer.Destroy()
	c.window.Destroy()
	sdl.Quit()
}
