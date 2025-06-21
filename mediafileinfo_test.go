package mediafileinfo

import (
    "testing"
)

// TestGetMediaInfo checks that GetMediaInfo returns a non-nil AVFormatContext for a valid file.
// You should provide a valid test media file path for this test to pass.
func TestGetMediaInfo(t *testing.T) {
    // Replace with a valid test media file path accessible during testing
    testFile := "testdata/sample.mp4"

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