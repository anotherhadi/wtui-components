package utils

type boxStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Verticaly   string
	Horizontaly string
	ToLeft      string
	ToRight     string
	ToTop       string
	ToBottom    string
	Middle      string
}

func GetBoxStyle(s string) boxStyle {
	styles := map[string]boxStyle{
		"single": {
			TopLeft:     "┌",
			TopRight:    "┐",
			BottomLeft:  "└",
			BottomRight: "┘",
			Verticaly:   "│",
			Horizontaly: "─",
			ToLeft:      "┤",
			ToRight:     "├",
			ToBottom:    "┬",
			ToTop:       "┴",
			Middle:      "┼",
		},
		"rounded": {
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "╰",
			BottomRight: "╯",
			Verticaly:   "│",
			Horizontaly: "─",
			ToLeft:      "┤",
			ToRight:     "├",
			ToBottom:    "┬",
			ToTop:       "┴",
			Middle:      "┼",
		},
		"double": {
			TopLeft:     "╔",
			TopRight:    "╗",
			BottomLeft:  "╚",
			BottomRight: "╝",
			Verticaly:   "║",
			Horizontaly: "═",
			ToLeft:      "╣",
			ToRight:     "╠",
			ToTop:       "╩",
			ToBottom:    "╦",
			Middle:      "╬",
		},
		"none": {
			TopLeft:     " ",
			TopRight:    " ",
			BottomLeft:  " ",
			BottomRight: " ",
			Verticaly:   " ",
			Horizontaly: " ",
			ToTop:       " ",
			ToLeft:      " ",
			ToRight:     " ",
			ToBottom:    " ",
			Middle:      " ",
		},
	}

	if _, exists := styles[s]; exists {
		return styles[s]
	} else {
		// Default box style
		return styles["rounded"]
	}
}
