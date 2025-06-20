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

// AVFormatContext represents the format context for a media file, mirroring FFmpeg's AVFormatContext.
// See: https://ffmpeg.org/doxygen/trunk/structAVFormatContext.html
type AVFormatContext struct {
    Filename       string      // Name of the media file (AVFormatContext->url)
    Streams        []AVStream  // List of streams present in the file (AVFormatContext->streams)
    StartTime      int64       // Start time of the stream in AV_TIME_BASE units (AVFormatContext->start_time)
    Duration       int64       // Duration of the media in AV_TIME_BASE units (AVFormatContext->duration)
    BitRate        int64       // Total bitrate of the media file (AVFormatContext->bit_rate)
    FormatName     string      // Short name of the format (AVInputFormat->name)
    FormatLongName string      // Long name of the format (AVInputFormat->long_name)
}

// AVStream represents a single stream (audio, video, subtitles, etc.) in a media file, similar to FFmpeg's AVStream.
// See: https://ffmpeg.org/doxygen/trunk/structAVStream.html
type AVStream struct {
    Index             int               // Stream index in the media file (AVStream->index)
    ID                int               // Stream ID (AVStream->id)
    CodecParameters   *AVCodecParameters// Codec parameters for this stream (AVStream->codecpar)
    TimeBase          AVRational        // Fundamental unit of time (AVStream->time_base)
    Duration          int64             // Duration of the stream (AVStream->duration)
    SampleAspectRatio AVRational        // Sample aspect ratio (AVStream->sample_aspect_ratio)
    AverageFrameRate  AVRational        // Average frame rate (AVStream->avg_frame_rate)
}

// AVRational represents a rational number, as used in FFmpeg for time bases and aspect ratios.
// See: https://ffmpeg.org/doxygen/trunk/structAVRational.html
type AVRational struct {
    Num int   // Numerator
    Den int   // Denominator
}

// AVCodecParameters stores codec parameters, equivalent to FFmpeg's AVCodecParameters.
// See: https://ffmpeg.org/doxygen/trunk/structAVCodecParameters.html
type AVCodecParameters struct {
    CodecType          int        // General type of the encoded data (AVMEDIA_TYPE_* in FFmpeg)
    CodecID            int        // Specific codec ID (AVCodecID)
    CodecTag           uint32     // Additional information about the codec (codec_tag/fourcc)
    ExtradataSize      int        // Size of extra binary data (extradata_size)
    NbCodedSideData    int        // Number of side data elements (nb_coded_side_data)
    Format             int        // Sample format (audio), pixel format (video)
    BitRate            int64      // Average bitrate (bit_rate)
    BitsPerCodedSample int        // Bits per coded sample (bits_per_coded_sample)
    BitsPerRawSample   int        // Bits per raw sample (bits_per_raw_sample)
    Profile            int        // Codec-specific profile (profile)
    Level              int        // Codec-specific level (level)
    Width              int        // Width of video (width)
    Height             int        // Height of video (height)
    AspectRatio        AVRational // Aspect ratio (sample_aspect_ratio)
    FieldOrder         int        // Field order (field_order)
    ColorRange         int        // Color range (color_range)
    ColorPrimaries     int32      // Color primaries (color_primaries)
    ColorTrc           int32      // Color transfer characteristic (color_trc)
    ColorSpace         int32      // Color space (color_space)
    ChromaLocation     int32      // Chroma location (chroma_location)
    ChannelLayout      int64      // Channel layout mask (channel_layout)
    Channels           int        // Number of audio channels (channels)
    VideoDelay         int        // Video delay (video_delay)
    SampleRate         int        // Audio sample rate (sample_rate)
    BlockAlign         int        // Block alignment (block_align)
    FrameSize          int        // Audio frame size (frame_size)
    InitialPadding     int        // Initial padding (initial_padding)
    TrailingPadding    int        // Trailing padding (trailing_padding)
    SeekPreroll        int        // Seek preroll (seek_preroll)
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
		Format: C.GoString(ctx.iformat.name),
		FormatName: C.GoString(ctx.iformat.long_name),
	}

	if formatCtx == nil {
		return nil, fmt.Errorf("could not map context to format context for file: %s", filename)
	}

	return formatCtx, nil
}
