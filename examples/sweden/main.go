package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/masv3971/humantouch"
)

func main() {
	// This will generate the population of Sweden symmetrical according to gender distribution.
	// It will take about 4,5m to run on my M1 macbook Air 2021, and consume about 9,5Gb of memory.
	rand.Seed(time.Now().Unix())

	person, err := humantouch.New(&humantouch.Config{
		DistributionCFG: &humantouch.DistributionCfg{
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

	populationOfSweden := 10099265

	start := time.Now()
	_, err = person.Distribution.RandomHumans(populationOfSweden)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(duration)
	// 4m31.826066042s

}
