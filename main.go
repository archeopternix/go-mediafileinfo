package main

import (
	"fmt"
)

const fname = "/home/archeopternix/Videos/P1000577_deshaked.mp4"

func main() {

	info, err := GetMediaInfo(fname)
	if err != nil {
		fmt.Printf("Failed to get media info: %v\n", err)
		return
	}

	PrintAVContextJSON(info)

}
