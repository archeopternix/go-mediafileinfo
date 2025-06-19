package main

/*
#cgo LDFLAGS: -L. -lavformat -lavcodec -lavutil
#include "mediainfowrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// AVFormatContext repräsentiert das Format-Kontext-Objekt wie in FFmpeg (vereinfachte Abbildung).
type AVFormatContext struct {
	Filename       string
	Streams        []AVStream
	StartTime      int64
	Duration       int64
	BitRate        int64
	FormatName     string
	FormatLongName string
}

// AVStream repräsentiert einen Stream wie in FFmpeg (vereinfachte Abbildung).
type AVStream struct {
	Index             int
	ID                int
	CodecParameters   *AVCodecParameters
	TimeBase          AVRational
	Duration          int64
	SampleAspectRatio AVRational
	AverageFrameRate  AVRational
}

type AVRational struct {
	Num int
	Den int
}

type AVCodecParameters struct {
	CodecType          int        // General type of the encoded data.
	CodecID            int        // Specific type of the encoded data (the codec used).
	CodecTag           uint32     // Additional information about the codec (corresponds to the AVI FOURCC).
	ExtradataSize      int        // Size of the extradata content in bytes.
	NbCodedSideData    int        // Amount of entries in coded_side_data.
	Format             int        //
	BitRate            int64      // The average bitrate of the encoded data (in bits per second).
	BitsPerCodedSample int        // The number of bits per sample in the codedwords.
	BitsPerRawSample   int        // This is the number of valid bits in each output sample.
	Profile            int        // Codec-specific bitstream restrictions that the stream conforms to.
	Level              int        //
	Width              int        // Video only.
	Height             int        //
	AspectRatio        AVRational // Video only.
	FieldOrder         int        // Video only.
	ColorRange         int        // Video only.
	ColorPrimaries     int32      //
	ColorTrc           int32      //
	ColorSpace         int32      //
	ChromaLocation     int32      //
	ChannelLayout      int64
	Channels           int
	VideoDelay         int // Video only.
	SampleRate         int // Audio only.
	BlockAlign         int // Audio only.
	FrameSize          int // Audio only.
	InitialPadding     int // Audio only.
	TrailingPadding    int // Audio only.
	SeekPreroll        int // Audio only.
}

func PrintAVContextJSON(params *AVFormatContext) error {
	data, err := json.MarshalIndent(*params, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// GetMediaInfo opens a media file and returns a MediaInfo.
func GetMediaInfo(filename string) (*AVFormatContext, error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	ctx := C.Get_avformat_context(cfilename)

	if ctx == nil {
		return nil, fmt.Errorf("could not open file: %s", filename)
	}

	defer C.Free_avformat_context(ctx)

	num := C.Get_stream_count(ctx)

	var streams []AVStream
	for i := range num {
		s := C.Get_stream_by_index(ctx, i)
		codecParams := &AVCodecParameters{
			CodecType:     int(s.codecpar.codec_type),
			CodecID:       int(s.codecpar.codec_id),
			BitRate:       int64(s.codecpar.bit_rate),
			Width:         int(s.codecpar.width),
			Height:        int(s.codecpar.height),
			SampleRate:    int(s.codecpar.sample_rate),
			ChannelLayout: int64(s.codecpar.channel_layout),
			Channels:      int(s.codecpar.channels),
			Format:        int(s.codecpar.format),
			AspectRatio:   AVRational{Num: int(s.codecpar.sample_aspect_ratio.num), Den: int(s.codecpar.sample_aspect_ratio.den)},
			FieldOrder:    int(s.codecpar.field_order),
			// ggf. weitere Felder
		}
		stream := AVStream{
			Index:             int(s.index),
			ID:                int(s.id),
			CodecParameters:   codecParams,
			TimeBase:          AVRational{Num: int(s.time_base.num), Den: int(s.time_base.den)},
			Duration:          int64(s.duration),
			SampleAspectRatio: AVRational{Num: int(s.sample_aspect_ratio.num), Den: int(s.sample_aspect_ratio.den)},
			AverageFrameRate:  AVRational{Num: int(s.avg_frame_rate.num), Den: int(s.avg_frame_rate.den)},
		}
		streams = append(streams, stream)
	}

	// FormatContext mappen
	formatCtx := &AVFormatContext{
		Filename: C.GoString((*C.char)(unsafe.Pointer(&ctx.filename[0]))),
		Streams:  streams,
		Duration: int64(ctx.duration),
		BitRate:  int64(ctx.bit_rate),
	}

	if formatCtx == nil {
		return nil, fmt.Errorf("could not map context to format context for file: %s", filename)
	}

	return formatCtx, nil
}
