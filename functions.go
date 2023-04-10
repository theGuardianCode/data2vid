package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func encode_frame(data []byte, file_info fs.FileInfo) {
	canvas := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1920, 1080}})

	combined_data := []byte(file_info.Name() + "nrsep" + fmt.Sprint(file_info.Size()) + "nrsep" + string(data))

	frame := write_to_image(combined_data, canvas)

	file, _ := os.Create("frame.png")
	png.Encode(file, frame)
}

func write_to_image(data []byte, frame *image.RGBA) *image.RGBA {

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

	x, y := 0, 0
	for i := 0; i < len(data); i += 3 {
		if x == 1920 {
			x = 0
			y++
		}

		frame.Set(x, y, color.RGBA{uint8(data[i]), uint8(data[i+1]), uint8(data[i+2]), uint8(255)})

		x++
	}

	return frame
}

func decode_frame(file_name string) {
	source, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}

	frame, _ := png.Decode(source)

	source.Close()

	var data []byte
	for y := 0; y < frame.Bounds().Max.Y; y++ {
		for x := 0; x < frame.Bounds().Max.X; x++ {
			r, g, b, _ := frame.At(x, y).RGBA()
			data = append(data, uint8(r), uint8(g), uint8(b))
		}
	}

	raw_text := string(data)
	elements := strings.Split(raw_text, "nrsep")

	file, err := os.Create(elements[0])
	if err != nil {
		panic(err)
	}

	length, _ := strconv.ParseInt(elements[1], 10, 64)

	elements[2] = string(([]byte(elements[2]))[:length])

	file.Write([]byte(elements[2]))
	file.Close()

	fmt.Printf("Image decoded to file %s", elements[0])
}
