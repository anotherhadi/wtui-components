package image

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"strings"

	"github.com/anotherhadi/wtui-components/ansi"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func getAvgColor(colors ...color.RGBA) color.RGBA {
	var avgR int
	var avgG int
	var avgB int
	for i := 0; i < len(colors); i++ {
		avgR += int(colors[i].R)
		avgG += int(colors[i].G)
		avgB += int(colors[i].B)
	}
	avgR /= 3
	avgG /= 3
	avgB /= 3

	return color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), 255}
}

func chooseBlocks(topleft, topright, bottomleft, bottomright color.RGBA) (block string, fgColor, bgColor color.RGBA) {
	blocks := []string{" ", "▘", "▝", "▖", "▗", "▄", "▚", "▞", "▐", "▟", "▙", "▜", "▛", "▀", "█"}

	avg := getAvgColor(topright, topleft, bottomright, bottomleft)

	diff1 := math.Abs(float64(topleft.R)-float64(avg.R)) + math.Abs(float64(topleft.G)-float64(avg.G)) + math.Abs(float64(topleft.B)-float64(avg.B))
	diff2 := math.Abs(float64(topright.R)-float64(avg.R)) + math.Abs(float64(topright.G)-float64(avg.G)) + math.Abs(float64(topright.B)-float64(avg.B))
	diff3 := math.Abs(float64(bottomleft.R)-float64(avg.R)) + math.Abs(float64(bottomleft.G)-float64(avg.G)) + math.Abs(float64(bottomleft.B)-float64(avg.B))
	diff4 := math.Abs(float64(bottomright.R)-float64(avg.R)) + math.Abs(float64(bottomright.G)-float64(avg.G)) + math.Abs(float64(bottomright.B)-float64(avg.B))

	minDiff := math.Max(math.Min(diff1, diff2), math.Min(diff3, diff4))

	switch minDiff {
	case diff1:
		block = blocks[1]
		fgColor = topleft
		bgColor = getAvgColor(topright, bottomleft, bottomright)
	case diff2:
		block = blocks[2]
		fgColor = topright
		bgColor = getAvgColor(topleft, bottomleft, bottomright)
	case diff3:
		block = blocks[3]
		fgColor = bottomleft
		bgColor = getAvgColor(topleft, topright, bottomright)
	case diff4:
		block = blocks[4]
		fgColor = bottomright
		bgColor = getAvgColor(topleft, topright, bottomleft)
	}

	return block, fgColor, bgColor
}

func getRGBPixels(imagePath string, width, height int) ([][]color.RGBA, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	img = resizeImage(img, width, height)

	bounds := img.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	pixels := make([][]color.RGBA, imgHeight)
	for y := range pixels {
		pixels[y] = make([]color.RGBA, imgWidth)
	}

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			pixel := img.At(x, y)
			rgba := color.RGBAModel.Convert(pixel).(color.RGBA)
			pixels[y][x] = rgba
		}
	}

	return pixels, nil
}

func resizeImage(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	resizedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			resizedImg.Set(x, y, img.At(x*imgWidth/width, y*imgHeight/height))
		}
	}

	return resizedImg
}

func Draw(imagePath string, width, height int) (img []string, err error) {
	pixels, err := getRGBPixels(imagePath, width, height)
	if err != nil {
		return nil, err
	}

	for y := 0; y < height; y += 2 {
		str := ""
		for x := 0; x < width; x += 2 {
			topleft := pixels[y][x]
			topright := pixels[y][x+1]
			bottomleft := pixels[y+1][x]
			bottomright := pixels[y+1][x+1]

			block, fgColor, bgColor := chooseBlocks(topleft, topright, bottomleft, bottomright)
			str += ansi.FgRgb(fgColor.R, fgColor.G, fgColor.B)
			str += ansi.BgRgb(bgColor.R, bgColor.G, bgColor.B)
			str += block
		}
		str += ansi.Reset
		img = append(img, str)
	}

	return img, nil
}

func Image(imagePath string, width, height int, customOpts ...wtopts.Opts) (err error) {
	opts := wtopts.GetOpts(customOpts...)

	strs, err := Draw(imagePath, width, height)
	if err != nil {
		return err
	}
	for _, str := range strs {
		fmt.Println(ansi.LineClear + strings.Repeat(" ", opts.LeftPadding) + str)
	}
	return nil
}
