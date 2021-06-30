package humantouch

import (
	"math/rand"
	"time"
)

func randomGender() string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	switch r := random.Intn(2); r {
	case 0:
		return GenderFemale
	case 1:
		return GenderMale
	}
	return ""
}
