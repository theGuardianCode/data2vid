package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"os"
	"strconv"
)

func encode_frame(data []byte, file_info fs.FileInfo) {
	canvas := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1920, 1080}})

	frame, x, y := write_to_image([]byte(file_info.Name()), 0, 0, canvas)
	frame, x, y = write_to_image([]byte("|"), x, y, frame)

	frame, x, y = write_to_image([]byte(strconv.FormatInt(file_info.Size(), 10)), x, y, frame)
	frame, x, y = write_to_image([]byte("|"), x, y, frame)

	frame, x, y = write_to_image(data, x, y, frame)
	fmt.Printf("Colour at {%d, %d} is: ", x-1, y)
	fmt.Print(frame.At(x-1, y))

	file, _ := os.Create("frame.png")
	png.Encode(file, frame)
}

func write_to_image(data []byte, x int, y int, frame *image.RGBA) (*image.RGBA, int, int) {

	if len(data)%3 != 0 {
		for i := 1; i < 3; i++ {
			if (len(data)+i)%3 != 0 {
				continue
			} else if (len(data)+i)%3 == 0 {
				for j := 0; j < i; j++ {
					data = append(data, uint8(0))
				}
			}
		}
	}

	for i := 0; i < len(data); i += 3 {
		if x == 1920 {
			x = 0
			y++
		}

		frame.Set(x, y, color.RGBA{uint8(data[i]), uint8(data[i+1]), uint8(data[i+2]), uint8(255)})

		x++
	}

	return frame, x, y
}

func decode_frame(file_name string) string {
	source, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}

	frame, _ := png.Decode(source)

	source.Close()

	var bytes []byte
	for i := 0; i < frame.Bounds().Max.X; i++ {
		r, g, b, _ := frame.At(i, 0).RGBA()
		bytes = append(bytes, uint8(r), uint8(g), uint8(b))
	}

	return string(bytes)
}
