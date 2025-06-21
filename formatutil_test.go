// Package mediafileinfo
package mediafileinfo

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{999, "999 B"},
		{1023, "1023 B"},
		{1024, "1.00 KB"},
		{1536, "1.50 KB"},
		{1048576, "1.00 MB"},
		{1572864, "1.50 MB"},
		{1073741824, "1.00 GB"},
		{1610612736, "1.50 GB"},
		{1099511627776, "1.00 TB"},
		{1649267441664, "1.50 TB"},
	}

	for _, tt := range tests {
		got := FormatBytes(tt.in)
		if got != tt.want {
			t.Errorf("FormatBytes(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFormatDurationMS(t *testing.T) {
	tests := []struct {
		in   uint64
		want string
	}{
		{0, "0.000"},
		{1, "0.001"},
		{999, "0.999"},
		{1000, "1.000"},
		{60000, "1:00.000"},
		{3599999, "59:59.999"},
		{3600000, "1:00:00.000"},
		{3661001, "1:01:01.001"},
		{86399999, "23:59:59.999"},
		{86400000, "24:00:00.000"},
		{12345678, "3:25:45.678"},
	}

	for _, tt := range tests {
		got := FormatDurationMS(tt.in)
		if got != tt.want {
			t.Errorf("FormatDurationMS(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
