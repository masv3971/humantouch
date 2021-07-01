package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/masv3971/humantouch"
)

func main() {
	rand.Seed(time.Now().Unix())

	human, err := humantouch.New(&humantouch.Config{
		DistrubutionCFG: &humantouch.DistributionCfg{
			Age0to10:    humantouch.AgeData{Weight: 65},
			Age10to20:   humantouch.AgeData{Weight: 60},
			Age20to30:   humantouch.AgeData{Weight: 63},
			Age30to40:   humantouch.AgeData{Weight: 69},
			Age40to50:   humantouch.AgeData{Weight: 67},
			Age50to60:   humantouch.AgeData{Weight: 66},
			Age60to70:   humantouch.AgeData{Weight: 54},
			Age70to80:   humantouch.AgeData{Weight: 48},
			Age80to90:   humantouch.AgeData{Weight: 18},
			Age90to100:  humantouch.AgeData{Weight: 3},
			Age100to110: humantouch.AgeData{Weight: 0},
		},
	})
	if err != nil {
		panic(err)
	}

	// Female ignores DistributionCfg, return female human, or error
	female, err := human.Female()
	if err != nil {
		panic(err)
	}

	// Male ingnores DistributionCfg, return a male human, or error
	male, err := human.Male()
	if err != nil {
		panic(err)
	}

	// RandomHuman ignores DistributionCfg, return a human, or error
	person, err := human.RandomHuman()
	if err != nil {
		panic(err)
	}

	// Females ignores DistributionCfg, return 50 female humans, or error
	females, err := human.Females(50)
	if err != nil {
		panic(err)
	}

	// Males ignores DistributionCfg, return 50 male humans
	males, err := human.Males(50)
	if err != nil {
		panic(err)
	}

	// RandomHumans ignores DistributionCfg, return 50 humans, or error
	persons, err := human.RandomHumans(50)
	if err != nil {
		panic(err)
	}

	// RandomHumasn return 50 humans according to the distribution
	personsDist, err := human.Distrubution.RandomHumans(50)
	if err != nil {
		panic(err)
	}

	// Females return 50 females according to the distribution
	femalesDist, err := human.Distrubution.Females(50)
	if err != nil {
		panic(err)
	}

	// Males return 50 males according to the distribution
	malesDist, err := human.Distrubution.Males(50)
	if err != nil {
		panic(err)
	}

	fmt.Println(female.Firstname,
		male.Firstname,
		female.Firstname,
		person.Firstname,

		females[0].Firstname,
		males[0].Firstname,
		persons[0].Firstname,

		personsDist[0].Firstname,
		femalesDist[0].Firstname,
		malesDist[0].Firstname,
	)
}
