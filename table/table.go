package table

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/utils"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func wrap(input string, maxCol int) []string {
	var result []string
	words := strings.Fields(input)
	var currentLine string

	for _, word := range words {
		if utf8.RuneCountInString(word) > maxCol {
			for len(word) > maxCol {
				result = append(result, word[:maxCol])
				word = word[maxCol:]
			}
			currentLine += word + " "
		} else if utf8.RuneCountInString(currentLine+word) <= maxCol {
			currentLine += word + " "
		} else {
			result = append(result, strings.TrimSpace(currentLine))
			currentLine = word + " "
		}
	}

	if currentLine != "" {
		result = append(result, strings.TrimSpace(currentLine))
	}

	return result
}

func getMaxLengthCol(maxLengthCol []int, defaultColLength int, content *[][]string) (result []int) {
	colMaxLength := make([]int, len(((*content)[0])))
	for _, row := range *content {
		for icol, col := range row {
			colMaxLength[icol] = max(utf8.RuneCountInString(col), colMaxLength[icol])
		}
	}

	result = make([]int, len(((*content)[0])))
	for i, cml := range colMaxLength {
		if len(maxLengthCol) == 0 || len(maxLengthCol) < len((*content)[0]) {
			result[i] = min(cml, defaultColLength)
		} else {
			result[i] = min(cml, maxLengthCol[i])
		}
	}

	return
}

func getMaxHeightRow(content *[][]string, maxLengthCol []int) (result []int) {
	result = make([]int, len(*content))
	for irow, row := range *content {
		for icol, col := range row {
			result[irow] = max(result[irow], len(wrap(col, maxLengthCol[icol]))-1)
		}
	}
	return
}

func Draw(content *[][]string, opts wtopts.Opts) []string {
	var strs []string
	var str string
	columnMaxLength := getMaxLengthCol(opts.ColumnMaxLength, opts.DefaultColLength, content)
	_ = getMaxHeightRow(content, columnMaxLength)

	boxStyle := utils.GetBoxStyle(opts.Style)

	str = ansi.Fg256(opts.SecondaryColor) + boxStyle.TopLeft
	for i, col := range columnMaxLength {
		str += strings.Repeat(boxStyle.Horizontally, col+2)
		if i != len(columnMaxLength)-1 {
			str += boxStyle.ToBottom
		}
	}
	str += boxStyle.TopRight
	str += ansi.Reset
	strs = append(strs, str)

	for irow, row := range *content {
		rowN := 0
		for icol, col := range row {
			rowN = max(rowN, len(wrap(col, columnMaxLength[icol])))
		}

		for i := 0; i < rowN; i++ {
			str := ""
			for icol, col := range row {
				wrapped := wrap(col, columnMaxLength[icol])
				str += ansi.Fg256(opts.SecondaryColor) + boxStyle.Vertically + " "
				if opts.FirstRowAccentColor && irow == 0 {
					str += ansi.Fg256(opts.AccentColor)
				} else {
					str += ansi.Fg256(opts.TextColor)
				}
				c := ""
				if i < len(wrapped) {
					c = wrapped[i]
				}
				str += c
				str += strings.Repeat(" ", columnMaxLength[icol]-utf8.RuneCountInString(c)+1)
			}
			str += ansi.Fg256(opts.SecondaryColor) + boxStyle.Vertically + ansi.Reset
			strs = append(strs, str)
		}

		if opts.Spacer && irow != len(*content)-1 {
			str = ansi.Fg256(opts.SecondaryColor) + boxStyle.ToRight
			for i, col := range columnMaxLength {
				str += strings.Repeat(boxStyle.Horizontally, col+2)
				if i != len(columnMaxLength)-1 {
					str += boxStyle.Middle
				}
			}
			str += boxStyle.ToLeft
			str += ansi.Reset
			strs = append(strs, str)
		}

	}

	str = ansi.Fg256(opts.SecondaryColor) + boxStyle.BottomLeft
	for i, col := range columnMaxLength {
		str += strings.Repeat(boxStyle.Horizontally, col+2)
		if i != len(columnMaxLength)-1 {
			str += boxStyle.ToTop
		}
	}
	str += boxStyle.BottomRight
	str += ansi.Reset
	strs = append(strs, str)

	return strs
}

func Table(content [][]string, customOpts ...wtopts.Opts) {
	opts := wtopts.GetOpts(customOpts...)
	strs := Draw(&content, opts)
	for _, str := range strs {
		fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
	}
}
