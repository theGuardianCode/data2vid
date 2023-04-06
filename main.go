package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	filename := arguments[2]

	if arguments[1] == "encode" {
		bytes, err := os.ReadFile(filename)

		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		file_info, _ := file.Stat()

		encode_frame(bytes, file_info)
	} else if arguments[1] == "decode" {
		decode_frame(filename)
	}
}
