# go-mediafileinfo
[![Go Reference](https://pkg.go.dev/badge/github.com/archeopternix/go-mediafileinfo.svg)](https://pkg.go.dev/github.com/archeopternix/go-mediafileinfo)

A Go library for extracting detailed media file information by linking FFmpeg libraries via cgo and providing a lean Go API.  
This package provides Go types and functions to access media stream metadata, including codecs, bitrates, durations, and more.

---

## Features

- Extracts file-level and stream-level metadata from audio/video files
- Go types mirroring FFmpeg structures (`AVFormatContext`, `AVStream`, `AVCodecParameters`)
- Outputs media information as formatted JSON
- Uses cgo to wrap FFmpeg libraries (`libavformat`, `libavutil`)

---

## Installation

Make sure you have the ffmpeg development libraries installed (`libavformat`, `libavutil`).
Install a C-compiler (e.g. gcc, mingw64..)

```bash
go get github.com/archeopternix/go-mediafileinfo
```

---

## Usage

```go
package main

import (
    "log"
    "github.com/archeopternix/go-mediafileinfo"
)

func main() {
    info, err := mediafileinfo.GetMediaInfo("example.mp4")
    if err != nil {
        log.Fatalf("Failed to get media info: %v", err)
    }
    err = mediafileinfo.PrintAVContextJSON(info)
    if err != nil {
        log.Fatalf("Failed to print media info: %v", err)
    }
}
```

---

## Public API

### Functions

#### GetMediaInfo

```go
func GetMediaInfo(filename string) (*AVFormatContext, error)
```

Opens a media file and returns a pointer to an *AVFormatContext* containing all extracted metadata.
Returns an error if the file cannot be opened or parsed.

#### PrintAVContextJSON

```go
func PrintAVContextJSON(params *AVFormatContext) error
```

Prints the AVFormatContext as JSON

---

## License

MIT License

---

## Notes

- This package requires cgo and the FFmpeg shared libraries to be available at runtime.
- Only a subset of FFmpeg metadata is currently mapped. Contributions are welcome!

