package humantouch

import (
	"time"

	rand2 "github.com/milosgajdos/go-estimate/rand"
)

// DistrubutionCfg holds configuration regarding the age distrubution
type DistrubutionCfg struct {
	Age0to10    AgeData
	Age10to20   AgeData
	Age20to30   AgeData
	Age30to40   AgeData
	Age40to50   AgeData
	Age50to60   AgeData
	Age60to70   AgeData
	Age70to80   AgeData
	Age80to90   AgeData
	Age90to100  AgeData
	Age100to110 AgeData
}

// Distrubution holds both Person and AgeDistrubution
type Distrubution struct {
	Age *DistrubutionCfg
}

func newDistrubutionClient(cfg *Config) (*Distrubution, error) {
	if cfg == nil {
		return nil, ErrAgeDistrubutionNotConfigured
	}
	c := &Distrubution{
		Age: cfg.DistrubutionCFG,
	}

	return c, nil
}

// Females return females according with the distrubution
func (d *Distrubution) Females(n int) ([]*Person, error) {
	p, err := d.newWithDistrubution(GenderFemale, n)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Males return males according with the distrubution
func (d *Distrubution) Males(n int) ([]*Person, error) {
	p, err := d.newWithDistrubution(GenderMale, n)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// RandomHumans return random humans according with the distrubution
func (d *Distrubution) RandomHumans(n int) ([]*Person, error) {
	p, err := d.newWithDistrubution("", n)
	if err != nil {
		return nil, err
	}
	return p, nil

}

func (d *Distrubution) newWithDistrubution(gender string, n int) ([]*Person, error) {
	var persons = []*Person{}

	years, err := d.createYears(n)
	if err != nil {
		return nil, err
	}

	for _, year := range years {
		p := &Person{}
		p.setYear(year)
		p.setMonth()
		p.setDay()
		p.setAge()
		p.setGender(d.createGender(gender))
		p.setName(gender)

		persons = append(persons, p)
	}
	return persons, nil
}

func (d *Distrubution) createGender(gender string) string {
	if gender == "" {
		switch r := random.Intn(1); {
		case r == 0:
			return GenderFemale
		case r == 1:
			return GenderMale
		}
	}
	return gender
}

func (d *Distrubution) createYears(numberOfTimes int) ([]int, error) {
	var p []float64
	d.Age.Age0to10.id = 1
	d.Age.Age10to20.id = 2
	d.Age.Age20to30.id = 3
	d.Age.Age30to40.id = 4
	d.Age.Age40to50.id = 5
	d.Age.Age50to60.id = 6
	d.Age.Age60to70.id = 7
	d.Age.Age70to80.id = 8
	d.Age.Age80to90.id = 9
	d.Age.Age90to100.id = 10
	d.Age.Age100to110.id = 11

	p = append(p, d.Age.Age0to10.Weight)
	p = append(p, d.Age.Age10to20.Weight)
	p = append(p, d.Age.Age20to30.Weight)
	p = append(p, d.Age.Age30to40.Weight)
	p = append(p, d.Age.Age40to50.Weight)
	p = append(p, d.Age.Age50to60.Weight)
	p = append(p, d.Age.Age60to70.Weight)
	p = append(p, d.Age.Age70to80.Weight)
	p = append(p, d.Age.Age80to90.Weight)
	p = append(p, d.Age.Age90to100.Weight)
	p = append(p, d.Age.Age100to110.Weight)

	draws, err := rand2.RouletteDrawN(p, numberOfTimes)
	if err != nil {
		return nil, err
	}

	years := []int{}
	doYear := func(min, max int) {
		yearNow := time.Now().Year()
		yearMax := yearNow - min
		yearMin := yearNow - max

		years = append(years, random.Intn(yearMax-yearMin+1)+yearMin)
	}

	for _, i := range draws {
		switch i {
		case 0:
			doYear(0, 10)
		case 1:
			doYear(10, 20)
		case 2:
			doYear(20, 30)
		case 3:
			doYear(30, 40)
		case 4:
			doYear(40, 50)
		case 5:
			doYear(50, 60)
		case 6:
			doYear(60, 70)
		case 7:
			doYear(70, 80)
		case 8:
			doYear(80, 90)
		case 9:
			doYear(90, 100)
		case 10:
			doYear(100, 110)

		}
	}

	return years, nil
}
