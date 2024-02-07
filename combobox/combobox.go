package combobox

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/asciimoji"
	"github.com/anotherhadi/wtui-components/getchar"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func filterStringsByPrefix(original *[]string, prefix string, caseSensitive bool) []string {
	if prefix == "" {
		return *original
	}

	var filtered []string

	if caseSensitive {
		for _, str := range *original {
			if strings.HasPrefix(str, prefix) {
				filtered = append(filtered, str)
			}
		}
	} else {
		var lowerPrefix string = strings.ToLower(prefix)
		var lowerStr string
		for _, str := range *original {
			lowerStr = strings.ToLower(str)
			if strings.HasPrefix(lowerStr, lowerPrefix) {
				filtered = append(filtered, str)
			}
		}
	}

	return filtered
}

func drawOption(opts wtopts.Opts, option string, selected bool) string {
	var str strings.Builder
	if selected {
		str.WriteString(ansi.Fg256(opts.AccentColor))
		str.WriteString(">")
	} else {
		str.WriteString(ansi.Fg256(opts.SecondaryColor))
		str.WriteString(" ")
	}

	str.WriteString(" ")
	if len(option) > opts.MaxCols-2 {
		option = option[:opts.MaxCols-2+3] + "..."
	}
	str.WriteString(option)
	str.WriteString(ansi.Reset)
	return str.String()
}

func Draw(selected *int, options *[]string, filter *string, opts wtopts.Opts) []string {
	var strs []string

	var filteredOptions []string = filterStringsByPrefix(options, *filter, opts.CaseSensitive)
	if len(filteredOptions) != 0 && *selected > len(filteredOptions)-1 {
		*selected = len(filteredOptions) - 1
	}

	start, end, moreBefore, moreAfter := utils.GetRange(len(filteredOptions), *selected, opts.MaxRows-1)

	strs = append(strs, ansi.Fg256(opts.SecondaryColor)+"[Filter]: "+ansi.Fg256(opts.AccentColor)+*filter+ansi.Reset)

	if moreBefore {
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+"  "+asciimoji.Up+" "+asciimoji.Up+ansi.Reset)
	}

	for i, option := range filteredOptions[start:end] {
		strs = append(strs, drawOption(opts, option, i+start == *selected))
	}

	if moreAfter {
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+"  "+asciimoji.Down+" "+asciimoji.Down+ansi.Reset)
	}
	return strs
}

func HandleKey(selected *int, options *[]string, filter *string, key string, opts wtopts.Opts) error {
	var filteredOptions []string = filterStringsByPrefix(options, *filter, opts.CaseSensitive)
	if key == "Down" {
		*selected++
		if *selected > len(filteredOptions)-1 {
			if opts.Loop {
				*selected = 0
			} else {
				*selected--
			}
		}
	} else if key == "Up" {
		*selected--
		if *selected < 0 {
			if opts.Loop {
				*selected = len(filteredOptions) - 1
			} else {
				*selected++
			}
		}
	} else if key == "" {
		return errors.New("Key not used")
	} else if key == "Backspace" || key == "Backspace2" {
		if len(*filter) > 0 {
			*filter = (*filter)[:len(*filter)-1]
		}
	} else {
		*filter += key
	}
	return nil
}

func Combobox(options []string, customOpts ...wtopts.Opts) (result int, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var selected int = 0
	var length int = -1
	var filter string = ""

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		fmt.Print(ansi.ScreenClearDown)
		strs := Draw(&selected, &options, &filter, opts)
		length = len(strs)
		for _, str := range strs {
			fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
		}

		ascii, arrow, err := getchar.GetChar()
		if err != nil || ascii == 3 || ascii == 4 {
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
			var filteredOptions []string = filterStringsByPrefix(&options, filter, opts.CaseSensitive)
			if len(filteredOptions) == 0 {
				return -1, errors.New("Empty selection")
			}
			result = slices.Index(options, filteredOptions[selected])
			if result == -1 {
				return -1, errors.New("Unable to find in list")
			} else {
				return result, nil
			}
		}
		err = HandleKey(&selected, &options, &filter, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return -1, err
		}
	}
}
