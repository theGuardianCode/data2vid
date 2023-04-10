package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	// var arguments = [3]string{"a", "decode", "frame.png"}

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
		filename := arguments[2]

		_, err := os.Open(filename)

		if os.IsNotExist(err) {
			fmt.Printf("%s does not exist", filename)
			return
		}

		decode_frame(filename)
	} else {
		fmt.Printf("%s is not a valid argument. Should be encode or decode.\n", arguments[1])
		return
	}
}
