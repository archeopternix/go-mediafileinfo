package mediafileinfo

import (
	"testing"
)

func TestAVFieldOrder_String(t *testing.T) {
	tests := []struct {
		val  AVFieldOrder
		want string
	}{
		{AV_FIELD_UNKNOWN, "UNKNOWN"},
		{AV_FIELD_PROGRESSIVE, "PROGRESSIVE"},
		{AV_FIELD_TT, "TT"},
		{AV_FIELD_BB, "BB"},
		{AV_FIELD_TB, "TB"},
		{AV_FIELD_BT, "BT"},
		{AVFieldOrder(99), "UNKNOWN"},
		{AVFieldOrder(-1), "UNKNOWN"},
	}

	for _, tt := range tests {
		got := tt.val.String()
		if got != tt.want {
			t.Errorf("AVFieldOrder(%d).String() = %q, want %q", tt.val, got, tt.want)
		}
	}
}
