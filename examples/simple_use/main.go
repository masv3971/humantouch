package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/masv3971/humantouch"
)

func main() {
	rand.Seed(time.Now().Unix())
	person, err := humantouch.New(&humantouch.Config{
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

	f0, err := person.Distribution.Females(50)
	if err != nil {
		panic(err)
	}
	fmt.Println("************** DIST FEMALE")
	for _, f := range f0 {
		fmt.Println(f.Age)
		fmt.Println(f.Firstname, f.Lastname)
		fmt.Println(f.SocialSecurityNumber.Swedish12.Complete)
		fmt.Println(f.SocialSecurityNumber.Swedish10.Complete)
		fmt.Println(f.Gender.General)
	}

	f1, err := person.Distribution.Males(50)
	if err != nil {
		panic(err)
	}
	fmt.Println("************** DIST MALE")
	for _, f := range f1 {
		fmt.Println(f.Age)
		fmt.Println(f.Firstname, f.Lastname)
		fmt.Println(f.SocialSecurityNumber.Swedish12.Complete)
		fmt.Println(f.SocialSecurityNumber.Swedish10.Complete)
		fmt.Println(f.Gender.General)
	}

	f2, err := person.Distribution.RandomHumans(50)
	if err != nil {
		panic(err)
	}
	fmt.Println("************** DIST RANDOM HUMANS")
	for _, f := range f2 {
		fmt.Println(f.Age)
		fmt.Println(f.Firstname, f.Lastname)
		fmt.Println(f.SocialSecurityNumber.Swedish12.Complete)
		fmt.Println(f.SocialSecurityNumber.Swedish10.Complete)
		fmt.Println(f.Gender.General)
	}

	f3, err := person.Females(50)
	if err != nil {
		panic(err)
	}
	fmt.Println("************** FEMALES")
	for _, f := range f3 {
		fmt.Println(f.Age)
		fmt.Println(f.Firstname, f.Lastname)
		fmt.Println(f.SocialSecurityNumber.Swedish12.Complete)
		fmt.Println(f.SocialSecurityNumber.Swedish10.Complete)
		fmt.Println(f.Gender.General)
	}

	f4, err := person.Males(50)
	if err != nil {
		panic(err)
	}
	fmt.Println("************** MALES")
	for _, f := range f4 {
		fmt.Println(f.Age)
		fmt.Println(f.Firstname, f.Lastname)
		fmt.Println(f.SocialSecurityNumber.Swedish12.Complete)
		fmt.Println(f.SocialSecurityNumber.Swedish10.Complete)
		fmt.Println(f.Gender.General)
	}

	f5, err := person.RandomHumans(50)
	if err != nil {
		panic(err)
	}
	fmt.Println("************** RANDOMHUMANS")
	for _, f := range f5 {
		fmt.Println(f.Age)
		fmt.Println(f.Firstname, f.Lastname)
		fmt.Println(f.SocialSecurityNumber.Swedish12.Complete)
		fmt.Println(f.SocialSecurityNumber.Swedish10.Complete)
		fmt.Println(f.Gender.General)
	}

	f6, err := person.RandomHuman()
	if err != nil {
		panic(err)
	}
	fmt.Println("************** RANDOMHUMAN")
	fmt.Println(f6.Age)
	fmt.Println(f6.Firstname, f6.Lastname)
	fmt.Println(f6.SocialSecurityNumber.Swedish12.Complete)
	fmt.Println(f6.SocialSecurityNumber.Swedish10.Complete)
	fmt.Println(f6.Gender.General)

	f7, err := person.Male()
	if err != nil {
		panic(err)
	}
	fmt.Println("************** MALE")
	fmt.Println(f7.Age)
	fmt.Println(f7.Firstname, f7.Lastname)
	fmt.Println(f7.SocialSecurityNumber.Swedish12.Complete)
	fmt.Println(f7.SocialSecurityNumber.Swedish10.Complete)
	fmt.Println(f7.Gender.General)

	f8, err := person.Female()
	if err != nil {
		panic(err)
	}
	fmt.Println("************** FEMALE")
	fmt.Println(f8.Age)
	fmt.Println(f8.Firstname, f8.Lastname)
	fmt.Println(f8.SocialSecurityNumber.Swedish12.Complete)
	fmt.Println(f8.SocialSecurityNumber.Swedish10.Complete)
	fmt.Println(f8.Gender.General)
}
