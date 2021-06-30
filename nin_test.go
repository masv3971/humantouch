package humantouch

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
	fmt.Println("lun", got.LuhnNumber.S, want.LuhnNumber.S)

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
