package stringtonum

import (
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		encoded string
		want    string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
		{"L", "10"},
		{"R", "01"},
		{"=", "00"},
		{"LR", "101"},
		{"RL", "010"},
		{"LL", "210"},
		{"RR", "012"},
		{"==", "000"},
		{"L==R", "10001"},
		{"R==L", "01110"},
	}

	for _, tt := range tests {
		t.Run(tt.encoded, func(t *testing.T) {
			got := Stringtonum(tt.encoded)
			if got != tt.want {
				t.Errorf("Decode(%q) = %q; want %q", tt.encoded, got, tt.want)
			}
		})
	}
}
