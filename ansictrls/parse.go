package ansictrls

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	attrDelimiterStart    = "#["
	attrDelimiterEnd      = "]"
	attrDelimiterStartLen = 2
	attrDelimiterEndLen   = 1
	attrDelimiterRE       = compileRegexp()
)

func compileRegexp() *regexp.Regexp {
	return regexp.MustCompile(regexp.QuoteMeta(attrDelimiterStart) +
		"(.*?)" + regexp.QuoteMeta(attrDelimiterEnd))
}

// SetAttributeDelimiters sets the attribute delimiters.
//
// Defaults: "#[", "]"
func SetAttributeDelimiters(start, end string) {
	attrDelimiterStart = start
	attrDelimiterEnd = end
	attrDelimiterStartLen = len(attrDelimiterStart)
	attrDelimiterEndLen = len(attrDelimiterEnd)
	attrDelimiterRE = compileRegexp()
}

// Text returns just the text without any control sequences.
//
// See: Parsing in the package comment.
func Text(s string) string {
	return attrDelimiterRE.ReplaceAllString(s, "")
}

// Parse returns the text with embedded control sequences.
//
// See: Parsing in the package comment.
func Parse(s string) string {
	return attrDelimiterRE.ReplaceAllStringFunc(s, attributes)
}

func attributes(s string) string {
	s = s[attrDelimiterStartLen : len(s)-attrDelimiterEndLen]
	builder := strings.Builder{}
	for _, x := range strings.Split(s, ";") {
		a := strings.Split(x, " ")
		if l := len(a); l == 1 && (a[0] == "" || a[0] == "reset") {
			builder.WriteString("0;")
			if defaultForegroundColor != nil {
				builder.WriteString(defaultForegroundColor.Foreground())
				builder.WriteRune(';')
			}
			if defaultBackgroundColor != nil {
				builder.WriteString(defaultBackgroundColor.Background())
				builder.WriteRune(';')
			}
		} else if l == 1 || l == 2 && a[0] == "not" {
			if style := mapStyle(a); style != nil {
				if l == 1 {
					builder.WriteString(style.Enable())
				} else {
					builder.WriteString(style.Disable())
				}
				builder.WriteRune(';')
			}
		} else if l == 2 && (a[0] == "fg" || a[0] == "bg") && a[1] == "default" {
			if a[0] == "fg" {
				if defaultForegroundColor == nil {
					builder.WriteString("39")
				} else {
					builder.WriteString(defaultForegroundColor.Foreground())
				}
			} else {
				if defaultBackgroundColor == nil {
					builder.WriteString("49")
				} else {
					builder.WriteString(defaultBackgroundColor.Background())
				}
			}
			builder.WriteRune(';')
		} else if (l == 2 || l == 3) && (a[0] == "fg" || a[0] == "bg") {
			if color := getColor(a); color != nil {
				if a[0] == "fg" {
					builder.WriteString(color.Foreground())
				} else {
					builder.WriteString(color.Background())
				}
				builder.WriteRune(';')
			}
		}
	}
	bs := builder.String()
	if len(bs) > 0 {
		bs = fmt.Sprintf("\x1b[%sm", strings.TrimRight(bs, ";"))
	}
	return bs
}

func mapStyle(a []string) *Style {
	var st string
	if len(a) == 1 {
		st = a[0]
	} else {
		st = a[1]
	}
	switch st {
	case "bd":
		return &Bold
	case "ft":
		return &Faint
	case "it":
		return &Italic
	case "ul":
		return &Underlined
	case "bk":
		return &Blink
	case "iv":
		return &Inverse
	case "co":
		return &Crossedout
	case "ol":
		return &Overlined
	default:
		return nil
	}
}

func getColor(a []string) Color {
	if len(a) == 2 {
		if a[1][0] == '#' {
			s := a[1][1:]
			if len(s) == 3 {
				s = fmt.Sprintf("%[1]c%[1]c%[2]c%[2]c%[3]c%[3]c", s[0], s[1], s[2])
			}
			if len(s) == 6 {
				return convRGB([]string{s[:2], s[2:4], s[4:]}, 16)
			}
		} else if strings.ContainsRune(a[1], ',') {
			if rgb := strings.Split(a[1], ","); len(rgb) == 3 {
				return convRGB(rgb, 10)
			}
		} else {
			return mapColor(false, a[1])
		}
	} else if a[1] == "bright" {
		return mapColor(true, a[2])
	}
	return nil
}

func convRGB(rgb []string, base int) Color {
	r, err := strconv.ParseUint(rgb[0], base, 8)
	if err != nil {
		return nil
	}
	g, err := strconv.ParseUint(rgb[1], base, 8)
	if err != nil {
		return nil
	}
	b, err := strconv.ParseUint(rgb[2], base, 8)
	if err != nil {
		return nil
	}
	return RGBColor(uint8(r), uint8(g), uint8(b))
}

func mapColor(b bool, s string) Color {
	var color Color
	if b {
		switch s {
		case "black":
			color = BrightBlack
		case "red":
			color = BrightRed
		case "green":
			color = BrightGreen
		case "yellow":
			color = BrightYellow
		case "blue":
			color = BrightBlue
		case "magenta":
			color = BrightMagenta
		case "cyan":
			color = BrightCyan
		case "white":
			color = BrightWhite
		}
	} else {
		switch s {
		case "black":
			color = Black
		case "red":
			color = Red
		case "green":
			color = Green
		case "yellow":
			color = Yellow
		case "blue":
			color = Blue
		case "magenta":
			color = Magenta
		case "cyan":
			color = Cyan
		case "white":
			color = White
		}
	}
	return color
}
