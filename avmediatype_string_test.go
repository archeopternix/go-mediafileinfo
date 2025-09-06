package mediafileinfo

import (
	"testing"
)

func TestAVMediaType_String(t *testing.T) {
	tests := []struct {
		val  AVMediaType
		want string
	}{
		{AVMEDIA_TYPE_UNKNOWN, "UNKNOWN"},
		{AVMEDIA_TYPE_VIDEO, "VIDEO"},
		{AVMEDIA_TYPE_AUDIO, "AUDIO"},
		{AVMEDIA_TYPE_DATA, "DATA"},
		{AVMEDIA_TYPE_SUBTITLE, "SUBTITLE"},
		{AVMEDIA_TYPE_ATTACHMENT, "ATTACHMENT"},
		{AVMEDIA_TYPE_NB, "NB"},
		{AVMediaType(99), "UNKNOWN"},
		{AVMediaType(-2), "UNKNOWN"},
		{AVMediaType(-1), "UNKNOWN"},
		{AVMediaType(0), "VIDEO"},
		{AVMediaType(1), "AUDIO"},
	}

	for _, tt := range tests {
		got := tt.val.String()
		if got != tt.want {
			t.Errorf("AVMediaType(%d).String() = %q, want %q", tt.val, got, tt.want)
		}
	}
}
