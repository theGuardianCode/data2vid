package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	encode_frame(bytes)
	// fmt.Println(decode_frame("frame.png"))
}
