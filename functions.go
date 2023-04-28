package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"math"
	"os"
)

func encode_frame(data []byte, file_info fs.FileInfo) {
	combined_data := []byte(file_info.Name() + "nrsep" + fmt.Sprint(file_info.Size()) + "nrsep" + string(data))
	num_of_frames := math.Ceil(float64(len(combined_data)/(1920*1080))) - 1

	if num_of_frames <= 0 {
		num_of_frames = 1
	}

	var frames []*image.RGBA

	for i := 0; i < int(num_of_frames); i++ {
		frames = append(frames, image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1920, 1080}}))
	}

	filled_frames := write_to_image(combined_data, frames)

	for i := 0; i < len(filled_frames); i++ {
		file, _ := os.Create(fmt.Sprintf("frame%d.png", i))
		png.Encode(file, filled_frames[i])
	}
}

func write_to_image(data []byte, frames []*image.RGBA) []*image.RGBA {

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

	frame_num := 0
	x, y := 0, 0
	for i := 0; i < len(data); i += 3 {
		if x == frames[frame_num].Bounds().Max.X && y != frames[frame_num].Bounds().Max.Y {
			x = 0
			y++
		} else if x == frames[frame_num].Bounds().Max.X && y == frames[frame_num].Bounds().Max.Y {
			frame_num++
			x, y = 0, 0
		}

		frames[frame_num].Set(x, y, color.RGBA{uint8(data[i]), uint8(data[i+1]), uint8(data[i+2]), uint8(255)})

		x++
	}

	return frames
}

func decode_frame(file_name string) []byte {
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

	return data
}
