package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func encode_frame(data []byte) {
	frame := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1920, 1080}})

	if len(data)%4 != 0 {
		for i := 1; i < 4; i++ {
			if (len(data)+i)%4 != 0 {
				continue
			} else if (len(data)+i)%4 == 0 {
				for j := 0; j < i; j++ {
					data = append(data, uint8(0))
				}
			}
		}
	}

	x, y := 0, 0
	for i := 0; i < len(data); i += 4 {
		if x == 1920 {
			x = 0
			y++
		}

		frame.Set(x, y, color.RGBA{uint8(data[i]), uint8(data[i+1]), uint8(data[i+2]), uint8(data[i+3])})

		x++
	}

	file, _ := os.Create("frame.png")
	png.Encode(file, frame)
}

func decode_frame(file_name string) []string {
	var a []string
	return a
}
