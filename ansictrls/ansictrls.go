package ansictrls

import "fmt"

const (
	PackageVersion = "0.1.0"
)

type EraseMode uint

// Erase modes for the functions EraseInLine() and EraseInDisplay().
const (
	EraseToEnd       EraseMode = iota // erase from cursor to end of screen/line
	EraseToBeginning                  // erase from cursor to beginning of screen/line
	EraseAll                          // erase entire screen/line
	EraseLines                        // clear the scrollback buffer (EraseInDisplay only)
)

// MoveCursor moves the cursor by rows/cols cells in the given direction.
//
// If rows/cols < 0 move cursor up/left,
// if rows/cols > 0 move cursor down/right,
// else do nothing.
func MoveCursor(rows, cols int) {
	if rows < 0 {
		fmt.Printf("\x1b[%dA", -rows) // CUU
	} else if rows > 0 {
		fmt.Printf("\x1b[%dB", rows) // CUD
	}
	if cols < 0 {
		fmt.Printf("\x1b[%dD", -cols) // CUB
	} else if cols > 0 {
		fmt.Printf("\x1b[%dC", cols) // CUF
	}
}

// MoveCursorTo moves the cursor to the given position.
//
// If row/col == 0 the row/column will not be changed.
func MoveCursorTo(row, col uint) {
	if row != 0 && col != 0 {
		fmt.Printf("\x1b[%d;%dH", row, col) // CUP
	} else if row != 0 {
		fmt.Printf("\x1b[%dd", row) // VPA
	} else if col != 0 {
		fmt.Printf("\x1b[%d`", col) // HPA
	}
}

// SaveCurrentCursorPostion saves the current cursor postion.
func SaveCurrentCursorPostion() {
	fmt.Print("\x1b[s") // SCP
}

// RestoreSavedCursorPosition restores the last saved cursor position.
func RestoreSavedCursorPosition() {
	fmt.Print("\x1b[u") // RCP
}

// HideCursor hides the cursor.
func HideCursor() {
	fmt.Print("\x1b[?25l")
}

// ShowCursor shows the cursor.
func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

// EraseInDisplay erases part of the screen (cursor position does not change).
//
// This function panics if it gets an unknown erase mode.
func EraseInDisplay(mode EraseMode) {
	switch mode {
	case EraseToEnd, EraseToBeginning, EraseAll, EraseLines:
		fmt.Printf("\x1b[%dJ", mode) // ED
	default:
		panic("unknown erase mode")
	}
}

// EraseInLine erases part of the line (cursor position does not change).
//
// This function panics if it gets an unknown erase mode.
func EraseInLine(mode EraseMode) {
	switch mode {
	case EraseToEnd, EraseToBeginning, EraseAll:
		fmt.Printf("\x1b[%dK", mode) // EL
	default:
		panic("unknown erase mode")
	}
}

// Scroll scrolls the whole page by n lines.
//
// If n < 0 scroll down (new lines are added at the top),
// if n > 0 scroll up (new lines are added at the bottom),
// else do nothing.
func Scroll(n int) {
	if n < 0 {
		fmt.Printf("\x1b[%dT", -n) // SD
	} else if n > 0 {
		fmt.Printf("\x1b[%dS", n) // SU
	}
}

// AlternativeScreenBuffer enables/disables the alternative screen buffer.
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-The-Alternate-Screen-Buffer
func AlternativeScreenBuffer(enable bool) {
	if enable {
		fmt.Print("\x1b[?1049h")
	} else {
		fmt.Print("\x1b[?1049l")
	}
}

// ClearScreen clears the screen and moves the cursor to the top left corner.
func ClearScreen() {
	EraseInDisplay(EraseAll)
	MoveCursorTo(1, 1)
}

// ResetScreen clears the screen and the scrollback buffer and moves the cursor to the top left corner.
func ResetScreen() {
	EraseInDisplay(EraseAll)
	EraseInDisplay(EraseLines)
	MoveCursorTo(1, 1)
}

// Home moves the cursor to the top left corner.
func Home() {
	MoveCursorTo(1, 1)
}
