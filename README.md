# go-mediafileinfo

A Go library for extracting detailed media file information using FFmpeg via cgo.  
This package provides Go types and functions to access media stream metadata, including codecs, bitrates, durations, and more.

---

## Features

- Extracts file-level and stream-level metadata from audio/video files
- Go types mirroring FFmpeg structures (`AVFormatContext`, `AVStream`, `AVCodecParameters`)
- Outputs media information as formatted JSON
- Uses cgo to wrap FFmpeg libraries (`libavformat`, `libavcodec`, `libavutil`)

---

## Installation

Make sure you have the FFmpeg development libraries installed (`libavformat`, `libavcodec`, `libavutil`).

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

### Types

#### AVFormatContext

Represents the media file context (simplified `AVFormatContext`).

```go
type AVFormatContext struct {
    Filename       string
    Streams        []AVStream
    StartTime      int64
    Duration       int64
    BitRate        int64
    FormatName     string
    FormatLongName string
}
```

#### AVStream

Describes an individual media stream (audio, video, subtitle, etc.).

```go
type AVStream struct {
    Index             int
    ID                int
    CodecParameters   *AVCodecParameters
    TimeBase          AVRational
    Duration          int64
    SampleAspectRatio AVRational
    AverageFrameRate  AVRational
}
```

#### AVCodecParameters

Codec and format-specific stream parameters.

```go
type AVCodecParameters struct {
    CodecType          int
    CodecID            int
    CodecTag           uint32
    ExtradataSize      int
    NbCodedSideData    int
    Format             int
    BitRate            int64
    BitsPerCodedSample int
    BitsPerRawSample   int
    Profile            int
    Level              int
    Width              int
    Height             int
    AspectRatio        AVRational
    FieldOrder         int
    ColorRange         int
    ColorPrimaries     int32
    ColorTrc           int32
    ColorSpace         int32
    ChromaLocation     int32
    ChannelLayout      int64
    Channels           int
    VideoDelay         int
    SampleRate         int
    BlockAlign         int
    FrameSize          int
    InitialPadding     int
    TrailingPadding    int
    SeekPreroll        int
}
```

#### AVRational

Represents a rational number (e.g., time base, aspect ratio).

```go
type AVRational struct {
    Num int
    Den int
}
```

---

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

---

## License

MIT License

---

## Notes

- This package requires cgo and the FFmpeg shared libraries to be available at runtime.
- Only a subset of FFmpeg metadata is currently mapped. Contributions are welcome!

