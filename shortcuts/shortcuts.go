package shortcuts

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func Draw(shortcuts [][2]string, opts wtopts.Opts) string {
	if len(shortcuts) == 0 {
		return ""
	}
	var strs []string
	var str string
	for _, shortcut := range shortcuts {
		strs = append(strs, ansi.Bold+"["+shortcut[0]+"] "+ansi.Reset+ansi.Fg256(opts.SecondaryColor)+shortcut[1])
	}

	str = ansi.Fg256(opts.SecondaryColor)
	str += strings.Join(strs[:], " ● ")
	str += ansi.Reset

	return str
}

func Shortcuts(shortcuts [][2]string, customOpts ...wtopts.Opts) {
	opts := wtopts.GetOpts(customOpts...)
	str := Draw(shortcuts, opts)
	fmt.Println(str)
}
