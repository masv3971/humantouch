package humantouch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCoordinationNumber(t *testing.T) {
	t.SkipNow()
	tts := []struct {
		name   string
		person Person
		want   string
	}{
		{
			name: "ok - male",
			person: Person{
				Firstname: "",
				Lastname:  "",
				BirthYear: BirthYear{
					S:            "",
					I:            14,
					SLong:        "",
					ILong:        2014,
					CenturyLongS: "",
					CenturyLongI: 0,
					CenturyS:     "",
					CenturyI:     0,
				},
				BirthMonth: BirthMonth{
					S: "",
					I: 1,
				},
				BirthDay: BirthDay{
					S: "",
					I: 10,
				},
				SocialSecurityNumber: &SocialSecurityNumber{},
				Gender:               Gender{},
				Age:                  0,
			},
			want: "140170",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got := newCoordinationNumber(tt.person)
			assert.Equal(t, tt.want, got)
		})
	}
}
