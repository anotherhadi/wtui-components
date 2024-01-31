package input

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/getchar"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func Draw(prompt *string, userInput *string, placeholder *string, opts wtopts.Opts) string {
	var str string = ""

	str += ansi.Fg256(opts.TextColor)
	str += *prompt + " "

	if *userInput == "" {
		str += ansi.Fg256(opts.SecondaryColor) + *placeholder
	} else {
		str += ansi.Fg256(opts.AccentColor) + *userInput
	}
	str += ansi.Reset

	return str
}

func HandleKey(userInput *string, key string, opts wtopts.Opts) error {
	if key == "Backspace" || key == "Backspace2" {
		if len(*userInput) != 0 {
			*userInput = (*userInput)[:len(*userInput)-1]
		}
	} else if key != "" {
		*userInput += key
	} else {
		return errors.New("Key not used")
	}
	return nil
}

func Input(prompt, placeholder string, customOpts ...wtopts.Opts) (userInput string, err error) {
	opts := wtopts.GetOpts(customOpts...)

	for {
		str := Draw(&prompt, &userInput, &placeholder, opts)
		fmt.Print(ansi.CursorCol(0), ansi.LineClear+strings.Repeat(" ", opts.LeftPadding)+str)

		ascii, _, err := getchar.GetChar()
		if err != nil {
			return "", err
		}
		if ascii == getchar.CtrlC || ascii == getchar.CtrlD {
			return "", errors.New("SIGINT")
		}

		key := string(ascii)
		if ascii == getchar.CR {
			fmt.Print("\n")
			return userInput, nil
		} else if ascii == getchar.Backspace {
			key = "Backspace"
		} else if ascii == getchar.Escape {
			return "", errors.New("Escape")
		}
		err = HandleKey(&userInput, key, opts)
		if err != nil {
			return "", err
		}
	}
}
