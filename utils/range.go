package utils

func GetRange(length, cursor, maxRows int) (start, end int, moreBefore, moreAfter bool) {
	if length <= maxRows {
		start = 0
		end = length
		moreBefore = false
		moreAfter = false
		return
	}

	cursor++

	onTop := int((maxRows-1)/2) + (maxRows+1)%2
	onBottom := int((maxRows - 1) / 2)

	if cursor-onTop <= 0 {
		start = 0
		end = start + maxRows
	} else if cursor+onBottom >= length {
		end = length
		start = length - maxRows
	} else {
		start = cursor - onTop - 1
		end = cursor + onBottom
	}

	if start != 0 {
		start++
		moreBefore = true
	}
	if end != length {
		end--
		moreAfter = true
	}

	return
}
