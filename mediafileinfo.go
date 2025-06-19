package main

/*
#cgo LDFLAGS: -L.  -lavformat -lavcodec -lavutil
#include "mediainfowrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type MediaInfo struct {
	Filename       string
	Duration       int64
	NbStreams      uint
	FormatName     string
	FormatLongName string
	// Add more fields if needed
}

func GetMediaInfo(filename string) (*MediaInfo, error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	cinfo := C.Get_media_info(cfilename)
	if cinfo == nil {
		return nil, fmt.Errorf("error in Get_Media_Info") // define this error as needed
	}
	defer C.Free_media_info(cinfo)

	goInfo := &MediaInfo{
		Filename:       C.GoString(&cinfo.filename[0]),
		Duration:       int64(cinfo.duration),
		NbStreams:      uint(cinfo.nb_streams),
		FormatName:     C.GoString(&cinfo.format_name[0]),
		FormatLongName: C.GoString(&cinfo.format_long_name[0]),
	}
	// Populate more fields as needed.

	return goInfo, nil
}
