package media

type Context interface {
	Update() []Button
	SetTitle(title string)
	Destroy()
}
