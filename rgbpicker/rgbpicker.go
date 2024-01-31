package rgbpicker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/asciimoji"
	"github.com/anotherhadi/wtui-components/getchar"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func Draw(rgb *[3]uint8, selected *int, opts wtopts.Opts) []string {
	boxStyle := utils.GetBoxStyle(opts.Style)
	var strs [5]string

	strs[0] = "  "
	for i := 0; i < 3; i++ {
		if *selected == i {
			strs[0] += ansi.Fg256(opts.AccentColor)
		} else {
			strs[0] += ansi.Fg256(opts.SecondaryColor)
		}
		strs[0] += asciimoji.Up + "   " + ansi.Reset
	}

	strs[1] = ansi.Fg256(opts.SecondaryColor) + boxStyle.TopLeft + strings.Repeat(strings.Repeat(boxStyle.Horizontaly, 3)+boxStyle.ToBottom, 3) + strings.Repeat(boxStyle.Horizontaly, 3) + boxStyle.TopRight + ansi.Reset

	strs[2] = ansi.Fg256(opts.SecondaryColor) + boxStyle.Verticaly
	for i := 0; i < 3; i++ {
		if *selected == i {
			strs[2] += ansi.Fg256(opts.AccentColor)
		} else {
			strs[2] += ansi.Fg256(opts.SecondaryColor)
		}
		strs[2] += fmt.Sprintf("%3d", rgb[i]) + ansi.Fg256(opts.SecondaryColor) + boxStyle.Verticaly
	}
	strs[2] += ansi.BgRgb(rgb[0], rgb[1], rgb[2]) + "   " + ansi.Reset + ansi.Fg256(opts.SecondaryColor) + boxStyle.Verticaly

	strs[3] = ansi.Fg256(opts.SecondaryColor) + boxStyle.BottomLeft + strings.Repeat(strings.Repeat(boxStyle.Horizontaly, 3)+boxStyle.ToTop, 3) + strings.Repeat(boxStyle.Horizontaly, 3) + boxStyle.BottomRight + ansi.Reset

	strs[4] = "  "
	for i := 0; i < 3; i++ {
		if *selected == i {
			strs[4] += ansi.Fg256(opts.AccentColor)
		} else {
			strs[4] += ansi.Fg256(opts.SecondaryColor)
		}
		strs[4] += asciimoji.Down + "   " + ansi.Reset
	}

	return strs[:]
}

func HandleKey(selected *int, rgb *[3]uint8, key string, opts wtopts.Opts) error {
	if key == "Down" || key == "j" {
		if rgb[*selected] == 0 {
			rgb[*selected] = 255
		} else {
			rgb[*selected]--
		}
	} else if key == "Up" || key == "k" {
		if rgb[*selected] == 255 {
			rgb[*selected] = 0
		} else {
			rgb[*selected]++
		}
	} else if key == "Left" || key == "h" {
		*selected--
		if *selected < 0 {
			if opts.Loop {
				*selected = 2
			} else {
				*selected = 0
			}
		}
	} else if key == "Right" || key == "l" {
		*selected++
		if *selected > 2 {
			if opts.Loop {
				*selected = 0
			} else {
				*selected = 2
			}
		}
	} else if key == "0" || key == "1" || key == "2" || key == "3" || key == "4" || key == "5" || key == "6" || key == "7" || key == "8" || key == "9" { // Manually input a number
		n, _ := strconv.ParseInt(key, 10, 64)
		if int(rgb[*selected])*10+int(n) <= 255 {
			rgb[*selected] = uint8(int(rgb[*selected])*10 + int(n))
		}
	} else if key == "Backspace" || key == "Backspace2" {
		if rgb[*selected] >= 10 {
			rgb[*selected] /= 10
		} else {
			rgb[*selected] = 0
		}
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func RGBPicker(customOpts ...wtopts.Opts) (rgb [3]uint8, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var length int = -1
	var selected int = 0
	rgb[0] = opts.DefaultRed
	rgb[1] = opts.DefaultGreen
	rgb[2] = opts.DefaultBlue

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		strs := Draw(&rgb, &selected, opts)
		length = len(strs)
		for _, str := range strs {
			fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
		}

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return [3]uint8{}, err
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
			return rgb, nil
		}
		err = HandleKey(&selected, &rgb, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return [3]uint8{}, err
		}
	}
}
