package humantouch

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	personnummer "github.com/personnummer/go"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestLuhn(t *testing.T) {
	got := TestSwedishNIN
	want := TestSwedishNIN

	p := &Person{
		BirthYear: BirthYear{
			S: "10",
		},
		BirthMonth: BirthMonth{
			S: "02",
		},
		BirthDay: BirthDay{
			S: "23",
		},
	}

	want.LuhnNumber.S = "2"
	got.luhn(p)

	if diff := cmp.Diff(want.LuhnNumber.S, got.LuhnNumber.S); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBirthNumber_male(t *testing.T) {
	c, _ := newNINClient()

	for i := 1; i <= 1000; i++ {
		bn := c.birthNumber(GenderMale)
		if bn.N3i%2 == 0 {
			t.Error("Even, but should be odd", bn.N3i)
		}
	}
}
func TestBirthNumber_female(t *testing.T) {
	c, _ := newNINClient()

	for i := 1; i <= 1000; i++ {
		bn := c.birthNumber(GenderFemale)
		if bn.N3i%2 != 0 {
			t.Error("Odd, but should be even", bn.N3i)
		}
	}
}

func TestDelimiter(t *testing.T) {
	have := TestSwedishNIN

	have.delimiter(100)
	if have.Delimiter != "+" {
		t.Error("Delimiter should be -", have.Delimiter)
	}
}

func TestSetComplete(t *testing.T) {
	type have struct {
		nin *SwedishNIN
		p   *Person
	}
	tts := []struct {
		name string
		have have
		want string
	}{
		{
			name: "OK",
			have: have{
				nin: &SwedishNIN{
					BirthNumber: BirthNumber{
						N1s:      "1",
						N2s:      "2",
						N3s:      "3",
						N1i:      1,
						N2i:      2,
						N3i:      3,
						Complete: "123",
					},
					LuhnNumber: LuhnNumber{
						S: "4",
						I: 4,
					},
					Complete:  "",
					Delimiter: "-",
				},
				p: &Person{
					BirthYear: BirthYear{
						S:            "85",
						I:            85,
						SLong:        "1985",
						ILong:        1985,
						CenturyLongS: "1900",
						CenturyLongI: 1900,
						CenturyS:     "19",
						CenturyI:     19,
					},
					BirthMonth: BirthMonth{
						S: "01",
						I: 1,
					},
					BirthDay: BirthDay{
						S: "02",
						I: 2,
					},
				},
			},
			want: "850102-1234",
		},
	}

	for _, tt := range tts {
		tt.have.nin.setComplet(tt.have.p)

		if diff := cmp.Diff(tt.want, tt.have.nin.Complete); diff != "" {
			t.Errorf("Name: %q, mismatch (-want +got):\n%s", tt.name, diff)
		}
	}
}

func TestSetComplete_12(t *testing.T) {
	type have struct {
		nin *SwedishNIN
		p   *Person
	}
	tts := []struct {
		name string
		have have
		want string
	}{
		{
			name: "OK",
			have: have{
				nin: &SwedishNIN{
					BirthNumber: BirthNumber{
						N1s:      "1",
						N2s:      "2",
						N3s:      "3",
						N1i:      1,
						N2i:      2,
						N3i:      3,
						Complete: "123",
					},
					LuhnNumber: LuhnNumber{
						S: "4",
						I: 4,
					},
					Complete:  "",
					Delimiter: "-",
				},
				p: &Person{
					BirthYear: BirthYear{
						S:            "85",
						I:            85,
						SLong:        "1985",
						ILong:        1985,
						CenturyLongS: "1900",
						CenturyLongI: 1900,
						CenturyS:     "19",
						CenturyI:     19,
					},
					BirthMonth: BirthMonth{
						S: "01",
						I: 1,
					},
					BirthDay: BirthDay{
						S: "02",
						I: 2,
					},
				},
			},
			want: "19850102-1234",
		},
	}

	for _, tt := range tts {
		tt.have.nin.setComplete12(tt.have.p)

		if diff := cmp.Diff(tt.want, tt.have.nin.Complete); diff != "" {
			t.Errorf("Name: %q, mismatch (-want +got):\n%s", tt.name, diff)
		}
	}
}

func TestValidateNIN(t *testing.T) {
	human, err := New(nil)
	if err != nil {
		t.Error(err)
	}

	persons, err := human.Females(59)
	if err != nil {
		t.Error(err)
	}

	for _, person := range persons {
		if !personnummer.Valid(person.SocialSecurityNumber.Swedish10.Complete) {
			t.Error("Not valid personnummer", person.SocialSecurityNumber.Swedish10.Complete)
		}
	}
}

func TestValidateNINDist(t *testing.T) {
	human, err := New(&Config{
		DistributionCFG: &DistributionCfg{
			Age0to10: AgeData{
				Weight: 100,
				id:     0,
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	persons, err := human.Distribution.Females(50)
	if err != nil {
		t.Error(err)
	}

	for _, person := range persons {
		if !personnummer.Valid(person.SocialSecurityNumber.Swedish10.Complete) {
			t.Error("Not valid personnummer", person.SocialSecurityNumber.Swedish10.Complete)
		}
	}
}

func TestValidateNINDist_10(t *testing.T) {
	human, err := New(&Config{
		DistributionCFG: &DistributionCfg{
			Age100to110: AgeData{
				Weight: 100,
				id:     10,
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	persons, err := human.Distribution.Females(50)
	if err != nil {
		t.Error(err)
	}

	for _, person := range persons {
		if !personnummer.Valid(person.SocialSecurityNumber.Swedish10.Complete) {
			t.Error("Not valid personnummer", person.SocialSecurityNumber.Swedish10.Complete)
		}
	}
}
