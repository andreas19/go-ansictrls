package ansictrls

import (
	"fmt"
	"strconv"
)

// Style attribute.
type Style [2]uint64

// Enable returns the SGR parameter for enabling the style attribute.
func (style Style) Enable() string {
	return strconv.FormatUint(style[0], 10)
}

// Disable returns the SGR parameter for disabling the style attribute.
func (style Style) Disable() string {
	return strconv.FormatUint(style[1], 10)
}

// Styles
var (
	Bold       = Style{1, 22}
	Faint      = Style{2, 22}
	Italic     = Style{3, 23}
	Underlined = Style{4, 24}
	Blink      = Style{5, 25}
	Inverse    = Style{7, 27}
	Crossedout = Style{9, 29}
	Overlined  = Style{53, 55}
)

// Reset resets all attributes.
func Reset() {
	fmt.Printf("\x1b[0m")
}
