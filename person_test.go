package humantouch

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(42)
}

func diffError(t *testing.T, name, diff string) {
	t.Errorf("Name:%s mismatch (-want +got):\n%s", name, diff)
}

var TestPerson = &Person{
	Firstname: "TestFirstname",
	Lastname:  "TestLastname",
	BirthYear: BirthYear{
		S:        "85",
		I:        85,
		SLong:    "1985",
		ILong:    1985,
		CenturyS: "19",
		CenturyI: 19,
	},
	BirthMonth: BirthMonth{
		S: "03",
		I: 3,
	},
	BirthDay: BirthDay{
		S: "10",
		I: 10,
	},
	Gender: Gender{
		General: GenderFemale,
		Ladok:   0,
	},
}

var TestSwedishNIN = &SwedishNIN{
	BirthNumber: BirthNumber{
		N1s:      "0",
		N2s:      "2",
		N3s:      "5",
		N1i:      0,
		N2i:      2,
		N3i:      5,
		Complete: "025",
	},
	LuhnNumber: LuhnNumber{},
	Complete:   "",
}

func TestSetDay(t *testing.T) {
	p := &Person{}

	for month := 1; month <= 12; month++ {
		for day := 1; day <= 100; day++ {
			p.BirthMonth.I = month
			p.setDay()

			if p.BirthDay.I > months[month] || p.BirthDay.I < 1 {
				t.Fatal("Error not in interval", p.BirthDay.I)
			}
		}
	}
}

func TestSetMonth(t *testing.T) {
	p := &Person{}

	for month := 1; month <= 100; month++ {
		p.setMonth()

		if p.BirthMonth.I > 12 || p.BirthMonth.I < 1 {
			t.Fatal("Error not in interval", p.BirthMonth.I)
		}
	}

}

func TestSetYear(t *testing.T) {
	c, _ := newPersonClient()
	p := &Person{}

	for year := 1; year <= 100; year++ {
		p.setYear(c.createYear())

		if p.BirthYear.ILong < yearMin || p.BirthYear.ILong > yearMax {
			t.Fatal("Error not in interval", p.BirthYear.ILong)
		}
	}
}

func TestSetGender(t *testing.T) {
	type have struct {
		p      *Person
		gender string
	}
	tts := []struct {
		name string
		have have
		want []string
	}{
		{
			name: "female",
			have: have{
				p:      TestPerson,
				gender: GenderFemale,
			},
			want: []string{GenderFemale, ""},
		},
		{
			name: "male",
			have: have{
				p:      TestPerson,
				gender: GenderMale,
			},
			want: []string{GenderMale, ""},
		},
		{
			name: "",
			have: have{
				p:      TestPerson,
				gender: "",
			},
			want: []string{GenderMale, GenderFemale},
		},
	}

	for _, tt := range tts {
		tt.have.p.setGender(tt.have.gender)
		if tt.have.p.Gender.General != tt.want[0] && tt.have.p.Gender.General != tt.want[1] {
			t.Errorf("Name:%q, want: %s/%s got: %s", tt.name, tt.want[0], tt.want[1], tt.have.p.Gender.General)
		}
	}
}

func TestSetName(t *testing.T) {
	tts := []struct {
		name string
		have *Person
	}{
		{
			name: "female",
			have: &Person{},
		},
	}

	for _, tt := range tts {
		tt.have.setName(GenderFemale)
	}
}
