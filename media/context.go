package media

type Context interface {
	Update() map[Button]bool

	Plot(x, y int32, col Colour)

	Present()

	SetWindowTitle(title string)
	WindowSize() (int32, int32)

	Destroy()
}
