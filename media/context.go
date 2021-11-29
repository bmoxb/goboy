package media

import mapset "github.com/deckarep/golang-set"

type Context interface {
	Update() mapset.Set

	Plot(x, y int32, col Colour)

	Present()

	SetWindowTitle(title string)
	WindowSize() (int32, int32)

	Destroy()
}
