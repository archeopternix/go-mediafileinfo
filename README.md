# go-mediafileinfo
[![Go Reference](https://pkg.go.dev/badge/github.com/archeopternix/go-mediafileinfo.svg)](https://pkg.go.dev/github.com/archeopternix/go-mediafileinfo)
[![codecov](https://codecov.io/gh/archeopternix/go-mediafileinfo/branch/main/graph/badge.svg)](https://codecov.io/gh/archeopternix/go-mediafileinfo)

A Go library for extracting detailed media file information by linking the FFmpeg libraries via cgo and providing a lean Go API.

This package enables you to analyze audio and video files and retrieve rich metadata, including:
- File-level information such as file size, extension, duration, bitrate, and format names
- Stream-level details for each audio, video, or subtitle stream (e.g. codec, resolution, sample rate, channel layout, aspect ratio, frame rate, and more)
- Go types that closely mirror FFmpeg's core structures (`AVFormatContext`, `AVStream`, `AVCodecParameters`)
- Output of all metadata as formatted JSON for easy integration or inspection

It is ideal for building media analysis tools, dashboards, or automated workflows that need to inspect media files programmatically in Go.

The package uses cgo to wrap FFmpeg libraries (`libavformat`, `libavutil`) and works on all platforms supported by [Go](https://go.dev/) and [FFmpeg](https://ffmpeg.org/).

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


## Platform Support

This package runs on all platforms that support [Go](https://go.dev/) and [FFmpeg](https://ffmpeg.org/), including Linux, macOS, and Windows (with appropriate FFmpeg and C compiler setup).

---


## Public API

### Functions

#### GetMediaInfo

```go
func GetMediaInfo(filename string) (*AVFormatContext, error)
```

Opens a media file and returns a pointer to an *AVFormatContext* containing all extracted metadata.
Returns an error if the file cannot be opened or parsed.


---

### AVFormatContext Structure

The `AVFormatContext` struct is the main data structure returned by `GetMediaInfo`.  
It contains detailed information about the media file and its streams, mirroring FFmpeg's AVFormatContext.

**Fields include:**

- `Filename` – Name of the media file.
- `FileExt` – File extension (e.g. "mp4").
- `FileSize` – File size in bytes.
- `FileSizeText` – Human-readable file size (e.g. "12.3 MB").
- `Streams` – List of all streams in the file (audio, video, subtitles, etc.).
- `StartTime` – Start time of the stream in AV_TIME_BASE units.
- `Duration` – Duration of the stream in AV_TIME_BASE units.
- `DurationText` – Human-readable duration (e.g. "00:03:21.45").
- `BitRate` – Total bitrate of the file in bits per second.
- `FormatName` – Short name of the format (e.g. "mov,mp4,m4a,3gp,3g2,mj2").
- `FormatLongName` – Long name of the format (e.g. "QuickTime / MOV").

Each stream in `Streams` is represented by an `AVStream` struct, which contains codec parameters and stream-specific metadata.

---

### AVStream Structure

The `AVStream` struct represents a single stream (audio, video, subtitles, etc.) in a media file, similar to FFmpeg's AVStream.

**Fields include:**

- `Index` – Stream index in AVFormatContext.
- `ID` – Format-specific stream ID.
- `CodecParameters` – Pointer to an `AVCodecParameters` struct describing codec properties.
- `TimeBase` – Time base for the stream timestamps (as an `AVRational`).
- `Duration` – Duration of the stream in stream time_base units.
- `DurationText` – Human-readable duration for the stream.
- `SampleAspectRatio` – Sample aspect ratio (width/height) for video streams.
- `AverageFrameRate` – Average frame rate for the stream.

The `CodecParameters` field contains detailed codec information such as codec type, codec ID, bitrate, resolution, sample rate, and more.

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

## License

MIT License

---

## Notes

- This package requires cgo and the FFmpeg shared libraries to be available at runtime.
- Only a subset of FFmpeg metadata is currently mapped. Contributions are welcome!
