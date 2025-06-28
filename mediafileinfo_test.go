package mediafileinfo

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestAVRational_String(t *testing.T) {
	tests := []struct {
		r    AVRational
		want string
	}{
		{AVRational{Num: 16, Den: 9}, "16:9"},
		{AVRational{Num: 1, Den: 1}, "1:1"},
		{AVRational{Num: 0, Den: 1}, "0:1"},
		{AVRational{Num: 3, Den: 2}, "3:2"},
	}

	for _, tt := range tests {
		got := tt.r.String()
		if got != tt.want {
			t.Errorf("AVRational.String() = %q, want %q", got, tt.want)
		}
	}
}

func TestAVCodecParameters_OmitZero(t *testing.T) {
	params := AVCodecParameters{
		Profile: 42,
		// All other integer fields are zero and should be omitted
		Width: 1024, // Should be present as omitempty is not set for omitzero
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	jsonStr := string(data)

	// Should contain "profile"
	if !strings.Contains(jsonStr, `"profile":42`) {
		t.Errorf("Expected 'profile' to be present in JSON: %s", jsonStr)
	}
	// Should contain "width"
	if !strings.Contains(jsonStr, `"width":1024`) {
		t.Errorf("Expected 'width' to be present in JSON: %s", jsonStr)
	}
	// Should NOT contain "channels" (zero value)
	if strings.Contains(jsonStr, `"channels"`) {
		t.Errorf("Did not expect 'channels' in JSON: %s", jsonStr)
	}
	// Should NOT contain "frame_size" (zero value)
	if strings.Contains(jsonStr, `"frame_size"`) {
		t.Errorf("Did not expect 'frame_size' in JSON: %s", jsonStr)
	}
}

// TestGetMediaInfo checks that GetMediaInfo returns a non-nil AVFormatContext for a valid file.
// You should provide a valid test media file path for this test to pass.
func TestGetMediaInfo(t *testing.T) {
	// Replace with a valid test media file path accessible during testing
	testFile := "testdata/sample.avi"

	info, err := GetMediaInfo(testFile)
	if err != nil {
		t.Fatalf("GetMediaInfo returned error: %v", err)
	}
	if info == nil {
		t.Fatal("GetMediaInfo returned nil AVFormatContext")
	}
	if len(info.Streams) == 0 {
		t.Error("Expected at least one stream in AVFormatContext")
	}
	if info.FormatLongName != "AVI (Audio Video Interleaved)" {
		t.Errorf("Expected FormatLongName to be 'AVI (Audio Video Interleaved)', got '%s'", info.FormatLongName)
	}
}

func TestPrintAVContextJSON(t *testing.T) {
	// Create a minimal AVFormatContext for testing
	ctx := &AVFormatContext{
		Filename: "test.mp4",
		FileSize: 12345,
		FileExt:  ".mp4",
		Streams:  []AVStream{},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := PrintAVContextJSON(ctx)
	if err != nil {
		t.Errorf("PrintAVContextJSON returned error: %v", err)
	}

	// Restore stdout and read output
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Errorf("Failed to read from pipe: %v", err)
	}
	output := buf.String()

	// Check that output contains expected JSON fields
	if !bytes.Contains([]byte(output), []byte(`"Filename": "test.mp4"`)) {
		t.Errorf("Output does not contain expected Filename field: %s", output)
	}
	if !bytes.Contains([]byte(output), []byte(`"FileSize": 12345`)) {
		t.Errorf("Output does not contain expected FileSize field: %s", output)
	}
}
