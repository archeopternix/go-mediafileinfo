package main

import (
	"encoding/json"
	"fmt"
)

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

func PrintAVCodecParametersJSON(params AVCodecParameters) error {
	data, err := json.MarshalIndent(params, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
