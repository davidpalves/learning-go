package iteration

import (
	"strings"
	"testing"
)

const text = "repeat"
const times = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat(text, times)
	expected := strings.Repeat(text, times)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(text, times)
	}
}
