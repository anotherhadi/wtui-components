package list

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/asciimoji"
	"github.com/anotherhadi/wtui-components/getchar"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func drawOption(opts wtopts.Opts, option [2]string, selected bool) [2]string {
	var strs [2]string
	var str strings.Builder
	boxStyle := utils.GetBoxStyle(opts.Style)

	if selected {
		str.WriteString(ansi.Fg256(opts.AccentColor))
	} else {
		str.WriteString(ansi.Fg256(opts.SecondaryColor))
	}
	str.WriteString(boxStyle.Verticaly + " ")
	if len(option[0]) > opts.MaxCols-2 {
		option[0] = option[0][:opts.MaxCols-(2+3)] + "..."
	}
	str.WriteString(option[0])
	strs[0] = str.String()
	str.Reset()
	if selected {
		str.WriteString(ansi.Fg256(opts.AccentColor))
	} else {
		str.WriteString(ansi.Fg256(opts.SecondaryColor))
	}
	str.WriteString(boxStyle.Verticaly + " ")
	if len(option[1]) > opts.MaxCols-2 {
		option[1] = option[1][:opts.MaxCols-(2+3)] + "..."
	}
	str.WriteString(option[1])
	str.WriteString(ansi.Reset)
	strs[1] = str.String()

	return strs
}

func Draw(selected *int, options *[][2]string, opts wtopts.Opts) []string {
	var strs []string

	start, end, moreBefore, moreAfter := utils.GetRange(len(*options), *selected, int(opts.MaxRows/3))

	if moreBefore {
		strs = append(strs, "")
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+"  "+asciimoji.Up+" "+asciimoji.Up+ansi.Reset)
		strs = append(strs, "")
	}

	for i, option := range (*options)[start:end] {
		lines := drawOption(opts, option, i+start == *selected)
		strs = append(strs, lines[0])
		strs = append(strs, lines[1])
		strs = append(strs, "")
	}

	if moreAfter {
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+"  "+asciimoji.Down+" "+asciimoji.Down+ansi.Reset)
		strs = append(strs, "")
		strs = append(strs, "")
	}
	return strs
}

func HandleKey(selected *int, options *[][2]string, key string, opts wtopts.Opts) error {
	if key == "Down" || key == "j" {
		*selected++
		if *selected > len(*options)-1 {
			if opts.Loop {
				*selected = 0
			} else {
				*selected--
			}
		}
	} else if key == "Up" || key == "k" {
		*selected--
		if *selected < 0 {
			if opts.Loop {
				*selected = len(*options) - 1
			} else {
				*selected++
			}
		}
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func List(options [][2]string, customOpts ...wtopts.Opts) (result int, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var selected int = 0
	var length int = -1

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		strs := Draw(&selected, &options, opts)
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
		if ascii == getchar.CR {
			fmt.Print(ansi.CursorVisible)
			return selected, nil
		}
		err = HandleKey(&selected, &options, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return -1, err
		}
	}
}
