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

	fmt.Printf("File: %s\n", info.Filename)
	fmt.Printf("Duration: %d\n", info.Duration)
	fmt.Printf("Number of streams: %d\n", info.NbStreams)
	fmt.Printf("Format: %s (%s)\n", info.FormatName, info.FormatLongName)
}
