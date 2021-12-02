package humantouch

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestGenderFromNIN(t *testing.T) {
	tts := []struct {
		name string
		nin  string
		want string
	}{
		{
			name: "Female 12",
			nin:  "199602230323",
			want: GenderFemale,
		},
		{
			name: "Female 10",
			nin:  "9602230323",
			want: GenderFemale,
		},
		{
			name: "Male 12",
			nin:  "199602230333",
			want: GenderMale,
		},
		{
			name: "Male 10",
			nin:  "9602230333",
			want: GenderMale,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genderFromNIN(tt.nin)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
