package mediafileinfo

import (
	"bytes"
	"os"
	"testing"
)

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
