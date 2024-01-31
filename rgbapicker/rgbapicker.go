package rgbapicker

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

func Draw(rgba *[4]uint8, selected *int, opts wtopts.Opts) []string {
	boxStyle := utils.GetBoxStyle(opts.Style)
	var strs [5]string

	strs[0] = "  "
	for i := 0; i < 4; i++ {
		if *selected == i {
			strs[0] += ansi.Fg256(opts.AccentColor)
		} else {
			strs[0] += ansi.Fg256(opts.SecondaryColor)
		}
		strs[0] += asciimoji.Up + "   " + ansi.Reset
	}

	strs[1] = ansi.Fg256(opts.SecondaryColor) + boxStyle.TopLeft + strings.Repeat(strings.Repeat(boxStyle.Horizontaly, 3)+boxStyle.ToBottom, 4) + strings.Repeat(boxStyle.Horizontaly, 3) + boxStyle.TopRight + ansi.Reset

	strs[2] = ansi.Fg256(opts.SecondaryColor) + boxStyle.Verticaly
	for i := 0; i < 4; i++ {
		if *selected == i {
			strs[2] += ansi.Fg256(opts.AccentColor)
		} else {
			strs[2] += ansi.Fg256(opts.SecondaryColor)
		}
		strs[2] += fmt.Sprintf("%3d", rgba[i]) + ansi.Fg256(opts.SecondaryColor) + boxStyle.Verticaly
	}
	strs[2] += ansi.BgRgb(rgba[0], rgba[1], rgba[2]) + "   " + ansi.Reset + ansi.Fg256(opts.SecondaryColor) + boxStyle.Verticaly

	strs[3] = ansi.Fg256(opts.SecondaryColor) + boxStyle.BottomLeft + strings.Repeat(strings.Repeat(boxStyle.Horizontaly, 3)+boxStyle.ToTop, 4) + strings.Repeat(boxStyle.Horizontaly, 3) + boxStyle.BottomRight + ansi.Reset

	strs[4] = "  "
	for i := 0; i < 4; i++ {
		if *selected == i {
			strs[4] += ansi.Fg256(opts.AccentColor)
		} else {
			strs[4] += ansi.Fg256(opts.SecondaryColor)
		}
		strs[4] += asciimoji.Down + "   " + ansi.Reset
	}

	return strs[:]
}

func HandleKey(selected *int, rgba *[4]uint8, key string, opts wtopts.Opts) error {
	if key == "Down" || key == "j" {
		if rgba[*selected] == 0 {
			rgba[*selected] = 255
		} else {
			rgba[*selected]--
		}
	} else if key == "Up" || key == "k" {
		if rgba[*selected] == 255 {
			rgba[*selected] = 0
		} else {
			rgba[*selected]++
		}
	} else if key == "Left" || key == "h" {
		*selected--
		if *selected < 0 {
			if opts.Loop {
				*selected = 3
			} else {
				*selected = 0
			}
		}
	} else if key == "Right" || key == "l" {
		*selected++
		if *selected > 3 {
			if opts.Loop {
				*selected = 0
			} else {
				*selected = 3
			}
		}
	} else if key == "0" || key == "1" || key == "2" || key == "3" || key == "4" || key == "5" || key == "6" || key == "7" || key == "8" || key == "9" { // Manually input a number
		n, _ := strconv.ParseInt(key, 10, 64)
		if int(rgba[*selected])*10+int(n) <= 255 {
			rgba[*selected] = uint8(int(rgba[*selected])*10 + int(n))
		}
	} else if key == "Backspace" || key == "Backspace2" {
		if rgba[*selected] >= 10 {
			rgba[*selected] /= 10
		} else {
			rgba[*selected] = 0
		}
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func RGBAPicker(customOpts ...wtopts.Opts) (rgba [4]uint8, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var length int = -1
	var selected int = 0
	rgba[0] = opts.DefaultRed
	rgba[1] = opts.DefaultGreen
	rgba[2] = opts.DefaultBlue
	rgba[3] = opts.DefaultOpacity

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		strs := Draw(&rgba, &selected, opts)
		length = len(strs)
		for _, str := range strs {
			fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
		}

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return [4]uint8{}, err
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
			return rgba, nil
		}
		err = HandleKey(&selected, &rgba, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return [4]uint8{}, err
		}
	}
}
