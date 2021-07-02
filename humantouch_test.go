package humantouch

import (
	"fmt"
	"math/rand"
)

func ExampleNew() {
	rand.Seed(42)

	human, _ := New(&Config{
		DistrubutionCFG: &DistributionCfg{
			Age0to10: AgeData{
				Weight: 100,
				id:     0,
			},
			Age10to20:   AgeData{},
			Age20to30:   AgeData{},
			Age30to40:   AgeData{},
			Age40to50:   AgeData{},
			Age50to60:   AgeData{},
			Age60to70:   AgeData{},
			Age70to80:   AgeData{},
			Age80to90:   AgeData{},
			Age90to100:  AgeData{},
			Age100to110: AgeData{},
		},
	})

	// Return female human, or error
	female, _ := human.Female()

	//Return a male human, or error
	male, _ := human.Male()

	females, _ := human.Females(50)

	males, _ := human.Males(50)

	fmt.Println(female.Firstname, male.Firstname, females[0].Firstname, males[0].Firstname)

	femaleDist, _ := human.Distribution.Females(50)

	maleDist, _ := human.Distribution.Males(50)

	fmt.Println(femaleDist[0].Firstname, maleDist[0].Firstname)
	// Output:
	// Frida Valentino Birgit Maths
	// Dilan Rony
}

func ExampleNew_random() {
	rand.Seed(42)

	human, _ := New(&Config{
		DistrubutionCFG: &DistributionCfg{
			Age0to10: AgeData{
				Weight: 100,
				id:     0,
			},
		},
	})

	randomDist, _ := human.Distribution.RandomHumans(50)
	fmt.Println("randomDist", randomDist[0].Firstname)

	randoms, _ := human.RandomHumans(50)
	fmt.Println("randoms", randoms[0].Firstname)

	random, _ := human.RandomHuman()
	fmt.Println("randomHuman", random.Firstname)
	// Output:
	// randomDist Adolfina
	// randoms Esad
	// randomHuman Celina
}
