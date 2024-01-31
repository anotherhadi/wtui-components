package numberpicker

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/asciimoji"
	"github.com/anotherhadi/wtui-components/getchar"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func stringToFloat(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}

func roundTo(n float64, decimals int) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func countDigitsAfterDecimal(num float64) int {
	strNum := fmt.Sprintf("%g", num)
	decimalPos := strings.Index(strNum, ".")
	if decimalPos != -1 {
		return len(strNum) - decimalPos - 1
	}
	return 0
}

func Draw(number *string, opts wtopts.Opts) []string {
	boxStyle := utils.GetBoxStyle(opts.Style)
	var strs [3]string
	maxLength := max(len(fmt.Sprintf("%d", int(opts.Maximum))), len(fmt.Sprintf("%d", int(opts.Minimum)))) + opts.Decimals
	if opts.Decimals > 0 { // .
		maxLength++
	}

	strs[0] = ansi.Fg256(opts.SecondaryColor) + "  " + boxStyle.TopLeft + strings.Repeat(boxStyle.Horizontaly, maxLength+2) + boxStyle.TopRight + "  " + ansi.Reset

	strs[1] = ansi.Fg256(opts.SecondaryColor) + asciimoji.Left + " " + boxStyle.Verticaly
	strs[1] += " " + strings.Repeat(" ", maxLength-len(*number))
	strs[1] += ansi.Fg256(opts.AccentColor) + *number
	strs[1] += ansi.Fg256(opts.SecondaryColor) + " " + boxStyle.Verticaly + " " + asciimoji.Right + ansi.Reset

	strs[2] = ansi.Fg256(opts.SecondaryColor) + "  " + boxStyle.BottomLeft + strings.Repeat(boxStyle.Horizontaly, maxLength+2) + boxStyle.BottomRight + "  " + ansi.Reset

	return strs[:]
}

func HandleKey(number *string, key string, opts wtopts.Opts) error {

	if key == "Down" || key == "Left" || key == "h" || key == "j" {
		n := stringToFloat(*number)
		n -= opts.Increment
		if n >= opts.Minimum {
			*number = fmt.Sprint(roundTo(n, opts.Decimals))
		}
	} else if key == "Up" || key == "Right" || key == "l" || key == "k" {
		n := stringToFloat(*number)
		n += opts.Increment
		if n <= opts.Maximum {
			*number = fmt.Sprint(roundTo(n, opts.Decimals))
		}
	} else if key == "0" || key == "1" || key == "2" || key == "3" || key == "4" || key == "5" || key == "6" || key == "7" || key == "8" || key == "9" { // Manually input a number
		n := stringToFloat(*number + key)
		if n >= opts.Minimum && n <= opts.Maximum && countDigitsAfterDecimal(n) <= opts.Decimals {
			if *number == "0" {
				*number = ""
			}
			*number += key
		}
	} else if key == "Backspace" || key == "Backspace2" {
		if len(*number) > 1 {
			*number = (*number)[:len(*number)-1]
		} else if len(*number) == 1 {
			*number = "0"
		}
	} else if key == "." {
		if !strings.Contains(*number, ".") {
			*number += "."
		}
	} else if key == "-" {
		if len(*number) > 0 {
			if (*number)[0] == '-' {
				*number = (*number)[1:]

			} else {
				*number = "-" + *number
			}

		} else {
			*number = "-"
		}
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func NumberPicker(customOpts ...wtopts.Opts) (result float64, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var length int = -1
	var number string = fmt.Sprint(roundTo(opts.DefaultNumber, opts.Decimals))

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		strs := Draw(&number, opts)
		length = len(strs)
		for _, str := range strs {
			fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
		}

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return -1, err
		}

		key := string(ascii)
		if arrow != "" {
			key = arrow
		}
		if ascii == getchar.Backspace {
			key = "Backspace"
		}
		if ascii == getchar.CR {
			fmt.Print(ansi.CursorVisible)
			n := stringToFloat(number)
			return roundTo(n, opts.Decimals), nil
		}
		err = HandleKey(&number, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return 0, err
		}
	}
}
