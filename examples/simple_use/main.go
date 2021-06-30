package main

import (
	"fmt"
	"humantouch"
)

func main() {
	human, err := humantouch.New(&humantouch.Config{
		DistrubutionCFG: &humantouch.DistrubutionCfg{
			Age0to10: humantouch.AgeData{
				Weight: 100,
			},
			Age10to20:   humantouch.AgeData{},
			Age20to30:   humantouch.AgeData{},
			Age30to40:   humantouch.AgeData{},
			Age40to50:   humantouch.AgeData{},
			Age50to60:   humantouch.AgeData{},
			Age60to70:   humantouch.AgeData{},
			Age70to80:   humantouch.AgeData{},
			Age80to90:   humantouch.AgeData{},
			Age90to100:  humantouch.AgeData{},
			Age100to110: humantouch.AgeData{},
		},
	})

	// Female ignores DistrubutionCFG, return female human, or error
	female, err := human.Female()
	if err != nil {
		panic(err)
	}

	// Male ingnores DistrubutionCFG, return a male human, or error
	male, err := human.Male()
	if err != nil {
		panic(err)
	}

	// RandomHuman ignores DistrubutionCFG, return a human, or error
	person, err := human.RandomHuman()
	if err != nil {
		panic(err)
	}

	// Females ignores DistubutionCFG, return 50 female humans, or error
	females, err := human.Females(50)
	if err != nil {
		panic(err)
	}

	// Males ignores DistubutionCFG, return 50 male humans
	males, err := human.Males(50)
	if err != nil {
		panic(err)
	}

	// RandomHumans ignores DistubutionCFG, return 50 humans, or error
	persons, err := human.RandomHumans(50)
	if err != nil {
		panic(err)
	}

	// RandomHumasn return 50 humans according to the distrubtion
	personsDist, err := human.Distrubution.RandomHumans(50)
	if err != nil {
		panic(err)
	}

	// Females return 50 females according to the distrubtion
	femalesDist, err := human.Distrubution.Females(50)
	if err != nil {
		panic(err)
	}

	// Males return 50 males according to the distrubution
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
