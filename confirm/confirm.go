package confirm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/getchar"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func Draw(state *bool, opts wtopts.Opts) string {
	var str strings.Builder

	if *state {
		str.WriteString(ansi.Bg256(opts.AccentColor))
	} else {
		str.WriteString(ansi.Bg256(opts.SecondaryColor))
	}
	str.WriteString(ansi.Fg256(opts.TextColor))
	str.WriteString("  " + opts.Affirmative + "  ")
	str.WriteString(ansi.Reset)
	str.WriteString("   ")
	if !*state {
		str.WriteString(ansi.Bg256(opts.AccentColor))
	} else {
		str.WriteString(ansi.Bg256(opts.SecondaryColor))

	}
	str.WriteString(ansi.Fg256(opts.TextColor))
	str.WriteString("  " + opts.Negative + "  ")
	str.WriteString(ansi.Reset)

	return str.String()
}

func HandleKey(state *bool, key string, opts wtopts.Opts) error {
	if key == "Down" || key == "j" || key == "Left" || key == "h" || key == "y" || key == "Y" {
		if opts.Loop {
			*state = !(*state)
		} else {
			*state = true
		}
	} else if key == "Up" || key == "k" || key == "Right" || key == "l" || key == "n" || key == "N" {
		if opts.Loop {
			*state = !(*state)
		} else {
			*state = false
		}
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func Confirm(customOpts ...wtopts.Opts) (state bool, err error) {
	opts := wtopts.GetOpts(customOpts...)

	var length int = -1
	state = opts.DefaultConfirm

	fmt.Print(ansi.CursorInvisible)

	for {
		fmt.Print(ansi.CursorUpN(length))
		fmt.Println(Draw(&state, opts))
		length = 1

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return false, err
		}

		key := string(ascii)
		if arrow != "" {
			key = arrow
		}
		if ascii == getchar.CR {
			fmt.Print(ansi.CursorVisible)
			return state, nil
		}
		err = HandleKey(&state, key, opts)
		if err != nil {
			fmt.Print(ansi.CursorVisible)
			return false, err
		}
	}
}
