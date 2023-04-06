package main

import (
	"fmt"
)

func main() {
	// bytes, err := os.ReadFile("test.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// file_info, _ := file.Stat()

	// encode_frame(bytes, file_info)
	fmt.Println(decode_frame("frame.png"))
}
