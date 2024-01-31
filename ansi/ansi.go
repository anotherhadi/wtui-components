package ansi

import (
	"fmt"
)

// Escape code
const ESC = "\033"

//
//        Text Styling
//

// Ex: fmt.Println(ansi.Underline + "Underlined text")
const (
	Reset      = ESC + "[0m"
	Bold       = ESC + "[1m"
	Faint      = ESC + "[2m"
	Italic     = ESC + "[3m"
	Underline  = ESC + "[4m"
	Blink      = ESC + "[5m"
	Reverse    = ESC + "[7m"
	CrossedOut = ESC + "[9m"
)

//
//        Foreground color
//

// Ex: fmt.Println(ansi.Blue + "Blue text")
const (
	Black   = ESC + "[30m"
	Red     = ESC + "[31m"
	Green   = ESC + "[32m"
	Yellow  = ESC + "[33m"
	Blue    = ESC + "[34m"
	Magenta = ESC + "[35m"
	Cyan    = ESC + "[36m"
	White   = ESC + "[37m"

	BrightBlack   = ESC + "[90m"
	BrightRed     = ESC + "[91m"
	BrightGreen   = ESC + "[92m"
	BrightYellow  = ESC + "[93m"
	BrightBlue    = ESC + "[94m"
	BrightMagenta = ESC + "[95m"
	BrightCyan    = ESC + "[96m"
	BrightWhite   = ESC + "[97m"
)

// 256 Foreground Color
// Ex: fmt.Println(ansi.Fg256(27) + "Blue text")
func Fg256(color uint8) string {
	return fmt.Sprintf("%s[38;5;%dm", ESC, color)
}

// RGB Foreground Color
// Ex: fmt.Println(ansi.FgRgb(30, 30, 255) + "Blue text")
func FgRgb(red, green, blue uint8) string {
	return fmt.Sprintf("%s[38;2;%d;%d;%dm", ESC, red, green, blue)
}

//
//        Background color
//

// Ex: fmt.Println(ansi.BgBlue + "Text with a blue background")
const (
	BgBlack   = ESC + "[40m"
	BgRed     = ESC + "[41m"
	BgGreen   = ESC + "[42m"
	BgYellow  = ESC + "[43m"
	BgBlue    = ESC + "[44m"
	BgMagenta = ESC + "[45m"
	BgCyan    = ESC + "[46m"
	BgWhite   = ESC + "[47m"

	BgBrightBlack   = ESC + "[100m"
	BgBrightRed     = ESC + "[101m"
	BgBrightGreen   = ESC + "[102m"
	BgBrightYellow  = ESC + "[103m"
	BgBrightBlue    = ESC + "[104m"
	BgBrightMagenta = ESC + "[105m"
	BgBrightCyan    = ESC + "[106m"
	BgBrightWhite   = ESC + "[107m"
)

// 256 Background Color
// Ex: fmt.Println(ansi.Bg256(27) + "Text with a blue background")
func Bg256(color uint8) string {
	return fmt.Sprintf("%s[48;5;%dm", ESC, color)
}

// RGB Background Color
// Ex: fmt.Println(ansi.BgRgb(30, 30, 255) + "Text with a blue background")
func BgRgb(red, green, blue uint8) string {
	return fmt.Sprintf("%s[48;2;%d;%d;%dm", ESC, red, green, blue)
}

//
//        Cursor Movement
//        X: Col   Y : Line

// Move cursor to {col}, {line}
// Ex: fmt.Print(ansi.CursorMove(10,20))
func CursorMove(x, y int) string {
	return fmt.Sprintf("%s[%d;%dH", ESC, y, x)
}

// Move cursor {line} up
// Ex: fmt.Print(ansi.CursorUp(10))
func CursorUpN(line int) string {
	return fmt.Sprintf("%s[%dA", ESC, line)
}

// Move cursor {line} down
// Ex: fmt.Print(ansi.CursorDown(10))
func CursorDownN(line int) string {
	return fmt.Sprintf("%s[%dB", ESC, line)
}

// Move cursor {col} right
// Ex: fmt.Print(ansi.CursorRight(10))
func CursorRightN(col int) string {
	return fmt.Sprintf("%s[%dC", ESC, col)
}

// Move cursor {col} left
// Ex: fmt.Print(ansi.CursorLeft(10))
func CursorLeftN(col int) string {
	return fmt.Sprintf("%s[%dD", ESC, col)
}

// Move cursor to {col}
// Ex: fmt.Print(ansi.CursorCol(10))
func CursorCol(col int) string {
	return fmt.Sprintf("%s[%dG", ESC, col)
}

const (
	CursorHome      = ESC + "[H"
	CursorSave      = ESC + "[s"
	CursorRestore   = ESC + "[u"
	CursorVisible   = ESC + "[?25h"
	CursorInvisible = ESC + "[?25l"
)

//
//        Alternative buffer
//

const (
	AlternativeBufferEnable  = ESC + "[?1049h"
	AlternativeBufferDisable = ESC + "[?1049l"
)

//
//        Screen
//

const (
	ScreenSave      = ESC + "[?47h"
	ScreenRestore   = ESC + "[?47l"
	ScreenClear     = ESC + "[2J"
	ScreenClearUp   = ESC + "[1J"
	ScreenClearDown = ESC + "[0J"
	ScreenClearEnd  = ESC + "[J"
)

//
//        Clear Lines
//

const (
	LineClear      = ESC + "[2K"
	LineClearStart = ESC + "[1K"
	LineClearEnd   = ESC + "[K"
)
