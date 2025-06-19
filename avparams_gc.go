package main

/*
#cgo pkg-config: libavcodec
#cgo LDFLAGS: -lavcodec
#cgo CFLAGS:  -I/usr/include/x86_64-linux-gnu

#include "avparams.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type AVCodecParameters struct {
	ptr *C.struct_AVCodecParameters
}

func NewAVCodecParameters(file string) *AVCodecParameters {
	p := C.get_video_codec_parameters(file)
	if p == nil {
		return nil
	}
	w := &AVCodecParameters{ptr: p}
	runtime.SetFinalizer(w, func(w *AVCodecParameters) {
		C.free_codec_parameters(w.ptr)
	})
	return w
}

// Access to fields of c-struct

func (w *AVCodecParameters) CodecTag() uint32 {
	return uint32(w.ptr.codec_tag)
}
func (w *AVCodecParameters) Extradata() unsafe.Pointer {
	return unsafe.Pointer(w.ptr.extradata)
}
func (w *AVCodecParameters) ExtradataSize() int {
	return int(w.ptr.extradata_size)
}
func (w *AVCodecParameters) Format() int {
	return int(w.ptr.format)
}
func (w *AVCodecParameters) BitsPerCodedSample() int {
	return int(w.ptr.bits_per_coded_sample)
}
func (w *AVCodecParameters) BitsPerRawSample() int {
	return int(w.ptr.bits_per_raw_sample)
}
func (w *AVCodecParameters) Profile() int {
	return int(w.ptr.profile)
}
func (w *AVCodecParameters) Level() int {
	return int(w.ptr.level)
}
func (w *AVCodecParameters) SampleAspectRatioNum() int {
	return int(w.ptr.sample_aspect_ratio_num)
}
func (w *AVCodecParameters) SampleAspectRatioDen() int {
	return int(w.ptr.sample_aspect_ratio_den)
}
func (w *AVCodecParameters) FieldOrder() int {
	return int(w.ptr.field_order)
}
func (w *AVCodecParameters) ColorRange() int {
	return int(w.ptr.color_range)
}
func (w *AVCodecParameters) ColorPrimaries() int {
	return int(w.ptr.color_primaries)
}
func (w *AVCodecParameters) ColorTrc() int {
	return int(w.ptr.color_trc)
}
func (w *AVCodecParameters) ColorSpace() int {
	return int(w.ptr.color_space)
}
func (w *AVCodecParameters) ChromaLocation() int {
	return int(w.ptr.chroma_location)
}
func (w *AVCodecParameters) VideoDelay() int {
	return int(w.ptr.video_delay)
}
func (w *AVCodecParameters) ChannelLayout() uint64 {
	return uint64(w.ptr.channel_layout)
}
func (w *AVCodecParameters) Channels() int {
	return int(w.ptr.channels)
}
func (w *AVCodecParameters) SampleRate() int {
	return int(w.ptr.sample_rate)
}
func (w *AVCodecParameters) BlockAlign() int {
	return int(w.ptr.block_align)
}
func (w *AVCodecParameters) FrameSize() int {
	return int(w.ptr.frame_size)
}
func (w *AVCodecParameters) InitialPadding() int {
	return int(w.ptr.initial_padding)
}
func (w *AVCodecParameters) TrailingPadding() int {
	return int(w.ptr.trailing_padding)
}
func (w *AVCodecParameters) SeekPreroll() int {
	return int(w.ptr.seek_preroll)
}

/*
// CodecName gets codec name via C-Helper
func (w *AVCodecParameters) CodecName() string {
	cstr := C. get_codec_name(w.ptr.codec_id)
	return C.GoString(cstr)
}
*/

// Ptr expose raw pointer when needed
func (w *AVCodecParameters) Ptr() *C.struct_AVCodecParameters {
	return w.ptr
}
