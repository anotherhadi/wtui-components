package checkbox

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

func drawOption(opts wtopts.Opts, option string, selected, checked bool) string {
	var str strings.Builder
	if selected {
		str.WriteString(ansi.Fg256(opts.AccentColor))
	} else {
		str.WriteString(ansi.Fg256(opts.SecondaryColor))
	}
	str.WriteString("[")

	if checked {
		str.WriteString(ansi.Fg256(opts.AccentColor) + "*")
	} else {
		str.WriteString(" ")
	}
	if selected {
		str.WriteString(ansi.Fg256(opts.AccentColor))
	} else {
		str.WriteString(ansi.Fg256(opts.SecondaryColor))
	}
	str.WriteString("] ")
	if len(option) > opts.MaxCols-4 {
		option = option[:opts.MaxCols-(4+3)] + "..."
	}
	str.WriteString(option)

	str.WriteString(ansi.Reset)
	return str.String()
}

func Draw(selected *int, options *[]string, checked *[]bool, opts wtopts.Opts) []string {
	var strs []string

	start, end, moreBefore, moreAfter := utils.GetRange(len(*options), *selected, opts.MaxRows)

	if moreBefore {
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+asciimoji.Up+" "+asciimoji.Up+ansi.Reset)
	}

	for i, option := range (*options)[start:end] {
		strs = append(strs, drawOption(opts, option, i+start == *selected, (*checked)[i+start]))
	}

	if moreAfter {
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+asciimoji.Down+" "+asciimoji.Down+ansi.Reset)
	}
	return strs
}

func HandleKey(selected *int, checked *[]bool, key string, opts wtopts.Opts) error {
	if key == "Down" || key == "j" {
		*selected++
		if *selected > len(*checked)-1 {
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
				*selected = len(*checked) - 1
			} else {
				*selected++
			}
		}
	} else if key == " " {
		(*checked)[*selected] = !(*checked)[*selected]
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func Checkbox(options []string, customOpts ...wtopts.Opts) (checked []bool, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var selected int = 0
	var length int = -1

	if opts.DefaultCheckbox != nil && len(opts.DefaultCheckbox) == len(options) {
		checked = opts.DefaultCheckbox
	} else {
		checked = make([]bool, len(options))
	}

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		strs := Draw(&selected, &options, &checked, opts)
		length = len(strs)
		for _, str := range strs {
			fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
		}

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return nil, err
		}

		key := string(ascii)
		if arrow != "" {
			key = arrow
		}
		if ascii == getchar.CR {
			fmt.Print(ansi.CursorVisible)
			return checked, nil
		}
		err = HandleKey(&selected, &checked, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return nil, err
		}
	}
}
