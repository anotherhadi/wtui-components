package breadcrumb

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func Draw(strs []string, opts wtopts.Opts) string {
	var str string
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return ansi.Fg256(opts.AccentColor) + strs[0]
	}

	str = ansi.Fg256(opts.SecondaryColor) + strings.Join(strs[:len(strs)-1], " > ") + " > " + ansi.Fg256(opts.AccentColor) + strs[len(strs)-1] + ansi.Reset

	return str
}

func Breadcrumb(strs []string, customOpts ...wtopts.Opts) {
	opts := wtopts.GetOpts(customOpts...)
	str := Draw(strs, opts)
	fmt.Println(strings.Repeat(" ", opts.LeftPadding) + str)
}
