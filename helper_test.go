package humantouch

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomGender(t *testing.T) {
	rand.Seed(42)

	tts := []struct {
		name string
		have string
		want string
	}{
		{
			name: "female",
			have: GenderFemale,
			want: GenderFemale,
		},
		{
			name: "male",
			have: GenderMale,
			want: GenderMale,
		},
		{
			name: "empty",
			have: "",
			want: GenderMale,
		},
	}

	for _, tt := range tts {
		got := randomGender(tt.have)

		if got != tt.want {
			t.Errorf("Name: %q, got: %s want: %s", tt.name, got, tt.want)
		}
	}
}

func TestMonth(t *testing.T) {
	for monthN := 1; monthN <= 12; monthN++ {
		got := month(time.Month(monthN))
		if got != monthN {
			t.Errorf("got: %d want: %d", got, monthN)
		}
	}
}
