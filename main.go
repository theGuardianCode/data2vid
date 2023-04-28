package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	// var arguments = [3]string{"a", "encode", "archive\\powerpoint.pptx"}

	if arguments[1] == "encode" {
		filename := arguments[2]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.Open(filename)
		if os.IsNotExist(err) {
			fmt.Printf("%s does not exist", filename)
			return
		}

		file_info, _ := file.Stat()

		encode_frame(data, file_info)
	} else if arguments[1] == "decode" {
		directory := arguments[2]

		frames, _ := os.ReadDir(directory)

		var data [][]byte
		for i := 0; i < len(frames); i++ {
			path := ".\\" + directory + "\\" + frames[i].Name()
			_, err := os.Open(path)

			if os.IsNotExist(err) {
				fmt.Printf("%s does not exist", path)
				return
			}
			data = append(data, decode_frame(path))
		}

		var raw_text string
		if len(frames) > 1 {
			raw_text = string(data[0]) + string(data[1])
		} else {
			raw_text = string(data[0])
		}

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
	} else {
		fmt.Printf("%s is not a valid argument. Should be encode or decode.\n", arguments[1])
		return
	}
}
