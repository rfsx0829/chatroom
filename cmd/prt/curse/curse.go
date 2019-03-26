package curse

import "fmt"

func ClearScreen() (n int, err error) {
	return fmt.Printf("\033[2J")
}

func MoveUp(x int) (n int, err error) {
	return fmt.Printf("\033[%dA", x)
}

func MoveDown(x int) (n int, err error) {
	return fmt.Printf("\033[%dB", x)
}

func MoveLeft(x int) (n int, err error) {
	return fmt.Printf("\033[%dD", x)
}

func MoveRight(x int) (n int, err error) {
	return fmt.Printf("\033[%dC", x)
}

func MoveTo(x, y int) (n int, err error) {
	return fmt.Printf("\033[%d;%dH", x, y)
}

func ResetCurse() (n int, err error) {
	return fmt.Printf("\033[H")
}

func HideCurse() (n int, err error) {
	return fmt.Printf("\033[?25l")
}

func ShowCurse() (n int, err error) {
	return fmt.Printf("\033[?25h")
}

func HighLight() (n int, err error) {
	return fmt.Printf("\033[7m")
}

func UnHighLight() (n int, err error) {
	return fmt.Printf("\033[27m")
}
