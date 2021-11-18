package media

type Button uint8

const (
	BUTTON_NONE Button = iota
	BUTTON_QUIT
	BUTTON_UP
	BUTTON_DOWN
	BUTTON_LEFT
	BUTTON_RIGHT
	BUTTON_A
	BUTTON_B
	BUTTON_SELECT
	BUTTON_START
)

func (b Button) String() string {
	switch b {
	case BUTTON_UP:
		return "↑"
	case BUTTON_DOWN:
		return "↓"
	case BUTTON_LEFT:
		return "←"
	case BUTTON_RIGHT:
		return "→"
	case BUTTON_A:
		return "A"
	case BUTTON_B:
		return "B"
	case BUTTON_SELECT:
		return "select"
	case BUTTON_START:
		return "start"
	}
	return ""
}
