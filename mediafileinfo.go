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
	NBStreams      uint32
	Streams        []AVStream
	StartTime      int64
	Duration       int64
	BitRate        int64
	Metadata       map[string]string
	FormatName     string
	FormatLongName string
	ProbeScore     int
	// Füge weitere Felder hinzu, falls benötigt
}

// AVStream repräsentiert einen Stream wie in FFmpeg (vereinfachte Abbildung).
type AVStream struct {
	Index                int
	ID                   int
	CodecParameters      *AVCodecParameters
	TimeBaseNum          int
	TimeBaseDen          int
	StartTime            int64
	Duration             int64
	NbFrames             int64
	Disposition          int
	Discard              int
	SampleAspectRatioNum int
	SampleAspectRatioDen int
	Metadata             map[string]string
	// Füge weitere Felder hinzu, falls benötigt
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
	SampleAspectRatio  AVRational // Video only.
	Framerate          AVRational // Video only.
	FieldOrder         int        // Video only.
	ColorRange         int        // Video only.
	ColorPrimaries     int32      //
	ColorTrc           int32      //
	ColorSpace         int32      //
	ChromaLocation     int32      //
	VideoDelay         int        // Video only.
	SampleRate         int        // Audio only.
	BlockAlign         int        // Audio only.
	FrameSize          int        // Audio only.
	InitialPadding     int        // Audio only.
	TrailingPadding    int        // Audio only.
	SeekPreroll        int        // Audio only.
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

	var streams []*AVStream
	for i := range num {
		s := C.Get_stream_by_index(ctx, i)
		codecParams := &AVCodecParameters{
			CodecType:     s.CodecType,
			CodecID:       s.CodecID,
			BitRate:       s.BitRate,
			Width:         s.Width,
			Height:        s.Height,
			SampleRate:    s.SampleRate,
			Channels:      s.Channels,
			ChannelLayout: s.ChannelLayout,
			Format:        s.Format,
			// ggf. weitere Felder
		}
		stream := &AVStream{
			Index:                s.Index,
			ID:                   s.ID,
			CodecParameters:      codecParams,
			TimeBaseNum:          s.TimeBaseNum,
			TimeBaseDen:          s.TimeBaseDen,
			StartTime:            s.StartTime,
			Duration:             s.Duration,
			NbFrames:             s.NbFrames,
			Disposition:          s.Disposition,
			Discard:              s.Discard,
			SampleAspectRatioNum: s.SampleAspectRatioNum,
			SampleAspectRatioDen: s.SampleAspectRatioDen,
			Metadata:             s.Metadata,
			// ggf. weitere Felder
		}
		streams = append(streams, stream)
	}

	// FormatContext mappen
	formatCtx := &AVFormatContext{
		Filename:       ctx.Filename,
		NBStreams:      ctx.NBStreams,
		Streams:        streams,
		StartTime:      ctx.StartTime,
		Duration:       ctx.Duration,
		BitRate:        ctx.BitRate,
		Metadata:       ctx.Metadata,
		FormatName:     ctx.FormatName,
		FormatLongName: ctx.FormatLongName,
		ProbeScore:     ctx.ProbeScore,
		// ggf. weitere Felder
	}

	if info == nil {
		return nil, fmt.Errorf("could not map context to format context for file: %s", filename)
	}

	return info, nil
}
