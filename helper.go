package humantouch

import (
	"math/rand"
)

func randomGender() string {
	switch r := rand.Intn(2); r {
	case 0:
		return GenderFemale
	case 1:
		return GenderMale
	}
	return ""
}
