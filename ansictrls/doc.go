/*
Package ansictrls provides functions for working with ANSI control sequences,
i.e. moving the cursor and coloring/styling output in a terminal.

See also

	- Wikipedia: https://en.wikipedia.org/wiki/ANSI_escape_code
	  - CSI (Control Sequence Introducer) sequences
	  - SGR (Select Graphic Rendition) parameters
	- XTerm Control Sequences: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html
	  - Functions using CSI

Parsing

Strings with embedded attributes can be parsed with Parse() which returns a string with
control sequences that produce a formatted output when printed to the terminal screen.
These attributes are delimited by default with "#[" (start) and "]" (end), but this can
be changed with SetAttributeDelimiters(). Attributes are seperated by ";".

Supported style attributes

	attr |  meaning
	-----|------------
	 bd  | bold
	 ft  | faint
	 it  | italic
	 ul  | underlined
	 bk  | blink
	 iv  | inverse
	 co  | crossed-out
	 ol  | overlined

To disable an attribute use "not <attr>".

Color attributes start with "fg" (foreground color) or "bg" (background color) followed
by a color.

	    color     |           description
	--------------|-------------------------------------------------------------------
	    <name>    | one of black, red, green, yellow, blue, magenta, cyan, white
	bright <name> | bright colors (<name> see above)
	    default   | set foreground/background color to the default color
	  <integer>   | index of an 8-bit color in the interval [0..255]
	 <r>,<g>,<b>  | 3 integers in the interval [0..255] for a 24-bit color (no spaces)
	  HTML-style  | 3 or 6 hexadecimal integer values, e.g. #00FF00 or #0F0

To reset all attributes to normal use "reset" or an empty string.

Example

	s := "#[bd;fg red;bg 0,255,0]Hello, #[not bd;fg bright blue]World!#[]"
	Parse(s)
	// "\x1b[1;38;5;1;48;2;0;255;0mHello, \x1b[22;38;5;12mWorld!\x1b[0m"
*/
package ansictrls
