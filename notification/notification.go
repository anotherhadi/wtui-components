package notification

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/asciimoji"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func formatMessage(str string, opts wtopts.Opts) []string {
	var result []string

	words := strings.Fields(str)

	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 <= opts.MaxCols {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			result = append(result, currentLine)
			currentLine = word
		}
	}

	if currentLine != "" {
		result = append(result, currentLine)
	}

	return result
}

func Draw(message, status string, opts wtopts.Opts) []string {
	var strs []string
	boxStyle := utils.GetBoxStyle(opts.Style)
	paraOpts := wtopts.DefaultOpts()
	paraOpts.MaxCols = opts.MaxCols - 4
	messageParagraph := formatMessage(message, paraOpts)

	strs = append(strs, ansi.Fg256(opts.SecondaryColor)+boxStyle.TopLeft+strings.Repeat(boxStyle.Horizontally, opts.MaxCols)+boxStyle.TopRight+ansi.Reset)
	for i := 0; i < len(messageParagraph); i++ {
		prefix := "   "
		if i == 0 {
			prefix = " "
			switch status {
			case "info":
				prefix += ansi.BrightBlack + asciimoji.Envelope + " "
			case "success":
				prefix += ansi.BrightGreen + asciimoji.Envelope + " "
			case "error":
				prefix += ansi.BrightRed + asciimoji.Envelope + " "
			case "warning":
				prefix += ansi.BrightYellow + asciimoji.Envelope + " "
			default:
				prefix += asciimoji.Envelope + " "
			}
		}
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+boxStyle.Vertically+prefix+ansi.Fg256(opts.TextColor)+messageParagraph[i]+strings.Repeat(" ", opts.MaxCols-len(messageParagraph[i])-3)+ansi.Fg256(opts.SecondaryColor)+boxStyle.Vertically+ansi.Reset)
	}

	strs = append(strs, ansi.Fg256(opts.SecondaryColor)+boxStyle.BottomLeft+strings.Repeat(boxStyle.Horizontally, opts.MaxCols)+boxStyle.BottomRight+ansi.Reset)

	return strs
}

func Notification(message, status string, customOpts ...wtopts.Opts) (box []string) {
	opts := wtopts.GetOpts(customOpts...)
	strs := Draw(message, status, opts)
	for _, str := range strs {
		fmt.Println(str)
	}
	return
}
