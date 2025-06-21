// Package mediafileinfo provides Go functions and structures for extracting
// media file information linking to the FFmpeg libraries by cgo. It defines types that mirror
// FFmpeg's AVFormatContext, AVStream, and AVCodecParameters, and offers a function
// to retrieve media metadata:
//
//	GetMediaInfo(filename string) (*AVFormatContext, error)
//
// and a function to print all the mediafile metainfo in JSON format.
package mediafileinfo

/*
#cgo LDFLAGS: -L. -lavformat -lavcodec -lavutil
#include "mediainfowrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"
)

// AVFormatContext represents the format context for a media file, mirroring FFmpeg's AVFormatContext.
// See: https://ffmpeg.org/doxygen/trunk/structAVFormatContext.html
type AVFormatContext struct {
	Filename       string     // Name of the media file.
	FileExt        string     // File externsion e.g. mp4
	FileSize       int64      // File size
	Streams        []AVStream // List of all streams in the file.
	StartTime      int64      // Start time of the stream in AV_TIME_BASE units.
	Duration       int64      // Duration of the stream in AV_TIME_BASE units.
	BitRate        int64      // Total bitrate of the file in bits per second.
	FormatName     string     // Short name of the format.
	FormatLongName string     // Long name of the format.
}

// AVStream represents a single stream (audio, video, subtitles, etc.) in a media file, similar to FFmpeg's AVStream.
// See: https://ffmpeg.org/doxygen/trunk/structAVStream.html
type AVStream struct {
	Index             int                // Stream index in AVFormatContext.
	ID                int                // Format-specific stream ID.
	CodecParameters   *AVCodecParameters // Codec parameters for this stream.
	TimeBase          AVRational         // Time base for the stream timestamps.
	Duration          int64              // Duration of the stream in stream time_base units.
	SampleAspectRatio AVRational         // Sample aspect ratio (width/height) for video.
	AverageFrameRate  AVRational         // Average frame rate.
}

// AVRational represents a rational number, as used in FFmpeg for time bases and aspect ratios.
// See: https://ffmpeg.org/doxygen/trunk/structAVRational.html
type AVRational struct {
	Num int // Numerator
	Den int // Denominator
}

// AVCodecParameters describes the properties of a single codec context.
// See: https://ffmpeg.org/doxygen/trunk/structAVCodecParameters.html
type AVCodecParameters struct {
	CodecType          int        // General type of the encoded data (see AVMediaType).
	CodecID            int        // Specific type of the encoded data (the codec used).
	CodecTag           uint32     // Additional information about the codec (corresponds to the AVI FOURCC).
	ExtradataSize      int        // Size of the extradata content in bytes.
	NbCodedSideData    int        // Amount of entries in coded_side_data.
	Format             int        // The pixel or sample format.
	BitRate            int64      // The average bitrate of the encoded data (in bits per second).
	BitsPerCodedSample int        // The number of bits per sample in the codedwords.
	BitsPerRawSample   int        // This is the number of valid bits in each output sample.
	Profile            int        // Codec-specific bitstream restrictions that the stream conforms to.
	Level              int        // Codec-specific level.
	Width              int        // Video only: width of the video frame.
	Height             int        // Video only: height of the video frame.
	AspectRatio        AVRational // Video only: sample aspect ratio.
	FieldOrder         int        // Video only: field order.
	ColorRange         int        // Video only: color range.
	ColorPrimaries     int32      // Video only: color primaries.
	ColorTrc           int32      // Video only: color transfer characteristic.
	ColorSpace         int32      // Video only: YUV colorspace type.
	ChromaLocation     int32      // Video only: location of chroma samples.
	ChannelLayout      int64      // Audio only: channel layout mask.
	Channels           int        // Audio only: number of audio channels.
	VideoDelay         int        // Video only: number of frames the decoder should delay.
	SampleRate         int        // Audio only: sampling rate.
	BlockAlign         int        // Audio only: block alignment.
	FrameSize          int        // Audio only: audio frame size.
	InitialPadding     int        // Audio only: initial padding.
	TrailingPadding    int        // Audio only: trailing padding.
	SeekPreroll        int        // Audio only: seek preroll.
}

// PrintAVContextJSON prints the AVFormatContext struct as formatted JSON to stdout.
// Returns an error if the struct cannot be marshaled to JSON.
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

	// Query file size and extension
	var fileSize int64
	var fileExt string
	if fi, err := os.Stat(filename); err == nil {
		fileSize = fi.Size()
		fileExt = filepath.Ext(filename)
	}

	// Map to FormatContext
	formatCtx := &AVFormatContext{
		Filename:       C.GoString((*C.char)(unsafe.Pointer(&ctx.filename[0]))),
		Streams:        streams,
		Duration:       int64(ctx.duration),
		BitRate:        int64(ctx.bit_rate),
		FormatName:     C.GoString(ctx.iformat.name),
		FormatLongName: C.GoString(ctx.iformat.long_name),
		FileSize:       fileSize,
		FileExt:        fileExt,
	}

	return formatCtx, nil
}
