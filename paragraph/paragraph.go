package paragraph

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func parseMarkdown(currentLine string, bold, italic *bool, opts wtopts.Opts) string {
	for {
		if strings.Contains(currentLine, "**") {
			if *bold {
				currentLine = strings.Replace(currentLine, "**", ansi.Reset+ansi.Fg256(opts.TextColor), 1)
			} else {
				currentLine = strings.Replace(currentLine, "**", ansi.Bold+ansi.Fg256(opts.AccentColor), 1)
			}
			*bold = !*bold
		} else if strings.Contains(currentLine, "*") {
			if *italic {
				currentLine = strings.Replace(currentLine, "*", ansi.Reset+ansi.Fg256(opts.TextColor), 1)
			} else {
				currentLine = strings.Replace(currentLine, "*", ansi.Italic+ansi.Fg256(opts.AccentColor), 1)
			}
			*italic = !*italic
		} else {
			break
		}
	}
	return currentLine
}

func Draw(str string, opts wtopts.Opts) []string {
	var result []string

	words := strings.Fields(str)

	var currentLine string
	var bold bool = false
	var italic bool = false
	for _, word := range words {
		if len(utils.StripAnsi(currentLine))+len(word)+1 <= opts.MaxCols {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			currentLine = parseMarkdown(currentLine, &bold, &italic, opts)
			result = append(result, currentLine)
			if italic {
				currentLine = ansi.Fg256(opts.AccentColor) + ansi.Italic
			} else if bold {
				currentLine = ansi.Fg256(opts.AccentColor) + ansi.Bold
			} else {
				currentLine = ansi.Fg256(opts.TextColor)
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		currentLine = parseMarkdown(currentLine, &bold, &italic, opts)
		result = append(result, ansi.Fg256(opts.TextColor)+currentLine+ansi.Reset)
	}

	return result
}

func Paragraph(str string, customOpts ...wtopts.Opts) {
	opts := wtopts.GetOpts(customOpts...)

	strs := Draw(str, opts)
	for _, str := range strs {
		fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
	}
	fmt.Println()
}
