package ansictrls

import "fmt"

type Color interface {
	// Foreground returns the SGR parameter for setting the foreground color.
	Foreground() string
	// Background returns the SGR parameter for setting the background color.
	Background() string
}

var (
	defaultForegroundColor Color
	defaultBackgroundColor Color
)

// 8-bit ANSI color.
//
// See: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
type ANSIColor uint8

// Foreground returns the SGR parameter for setting the foreground color.
func (c ANSIColor) Foreground() string {
	return fmt.Sprintf("38;5;%d", c)
}

// Background returns the SGR parameter for setting the background color.
func (c ANSIColor) Background() string {
	return fmt.Sprintf("48;5;%d", c)
}

// ANSI colors
const (
	Black ANSIColor = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

type rgbColor struct {
	R, G, B uint8
}

// Foreground returns the SGR parameter for setting the foreground color.
func (c rgbColor) Foreground() string {
	return fmt.Sprintf("38;2;%d;%d;%d", c.R, c.G, c.B)
}

// Background returns the SGR parameter for setting the background color.
func (c rgbColor) Background() string {
	return fmt.Sprintf("48;2;%d;%d;%d", c.R, c.G, c.B)
}

// RGBColor returns a 24-bit color.
//
// See: https://en.wikipedia.org/wiki/ANSI_escape_code#24-bit
func RGBColor(r, g, b uint8) Color {
	return rgbColor{r, g, b}
}

// SetForegroundColor sets the default foreground color.
func SetForegroundColor(color Color) {
	defaultForegroundColor = color
	fmt.Printf("\x1b[%sm", color.Foreground())
}

// SetBackgroundColor sets the default background color.
func SetBackgroundColor(color Color) {
	defaultBackgroundColor = color
	fmt.Printf("\x1b[%sm", color.Background())
}

// ResetForegroundColor resets the foreground color to the terminal's default.
func ResetForegroundColor() {
	defaultForegroundColor = nil
	fmt.Printf("\x1b[39m")
}

// ResetBackgroundColor resets the background color to the terminal's default.
func ResetBackgroundColor() {
	defaultBackgroundColor = nil
	fmt.Printf("\x1b[49m")
}
