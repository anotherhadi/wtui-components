package floatingwindow

import (
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func Draw(width, height int, title string, opts wtopts.Opts) []string {
	if width <= 0 || height <= 0 {
		return []string{}
	}

	boxStyle := utils.GetBoxStyle(opts.Style)
	var strs []string
	strs = append(strs, ansi.Fg256(opts.SecondaryColor)+boxStyle.TopLeft+boxStyle.Horizontally+title+strings.Repeat(boxStyle.Horizontally, width-1-len(title))+boxStyle.TopRight+ansi.Reset)
	for i := 0; i < height; i++ {
		strs = append(strs, ansi.Fg256(opts.SecondaryColor)+boxStyle.Vertically+strings.Repeat(" ", width)+boxStyle.Vertically+ansi.Reset)
	}

	strs = append(strs, ansi.Fg256(opts.SecondaryColor)+boxStyle.BottomLeft+strings.Repeat(boxStyle.Horizontally, width)+boxStyle.BottomRight+ansi.Reset)
	return strs
}

func FloatingWindow(width, height int, title string, customOpts ...wtopts.Opts) (box []string) {
	opts := wtopts.GetOpts(customOpts...)
	box = Draw(width, height, title, opts)
	return
}
