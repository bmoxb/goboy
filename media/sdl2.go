package media

import "github.com/veandco/go-sdl2/sdl"

type sdlContext struct {
	window       *sdl.Window
	surface      *sdl.Surface
	keyButtonMap map[sdl.Keycode]Button
	buttonsDown  map[Button]bool
}

func NewSDL2(title string, width, height int32, keyButtonMap map[sdl.Keycode]Button) (sdlContext, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return sdlContext{}, err
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Quit()
		return sdlContext{}, err
	}

	surface, err := window.GetSurface()
	if err != nil {
		window.Destroy()
		sdl.Quit()
		return sdlContext{}, err
	}

	buttonsDown := map[Button]bool{}

	return sdlContext{window, surface, keyButtonMap, buttonsDown}, nil
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

	c.window.UpdateSurface()

	return c.buttonsDown
}

func (c sdlContext) SetTitle(title string) {
	c.window.SetTitle(title)
}

func (c sdlContext) Destroy() {
	c.window.Destroy()
	sdl.Quit()
}
