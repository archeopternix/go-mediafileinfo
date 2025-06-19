package main

/*
#cgo LDFLAGS: -L. -lavformat -lavcodec -lavutil
#include "mediainfowrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type MediaInfo struct {
	ctx *C.AVFormatContext
}

// GetMediaInfo opens a media file and returns a MediaInfo.
func GetMediaInfo(filename string) (*MediaInfo, error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	ctx := C.Get_avformat_context(cfilename)
	if ctx == nil {
		return nil, fmt.Errorf("could not open file: %s", filename)
	}
	return &MediaInfo{ctx: ctx}, nil
}

// Close releases the AVFormatContext.
func (mc *MediaInfo) Close() {
	if mc.ctx != nil {
		C.Free_avformat_context(mc.ctx)
		mc.ctx = nil
	}
}

// NbStreams returns the number of streams.
func (mc *MediaInfo) NbStreams() int {
	if mc.ctx == nil {
		return 0
	}
	return int(C.Get_avstreams(mc.ctx))
}

// Duration returns the duration in AV_TIME_BASE units.
func (mc *MediaInfo) Duration() int64 {
	if mc.ctx == nil {
		return 0
	}
	return int64(C.Get_duration(mc.ctx))
}

// FormatName returns the short name of the format.
func (mc *MediaInfo) FormatName() string {
	if mc.ctx == nil {
		return ""
	}
	return C.GoString(C.Get_format_name(mc.ctx))
}

// Example: Add more helpers as needed.
