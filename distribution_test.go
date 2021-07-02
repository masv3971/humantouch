package humantouch

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestCreateYears(t *testing.T) {
	type want struct {
		olders  int
		yungest int
	}
	tts := []struct {
		name string
		have *Distribution
		want want
	}{
		{
			name: "0-10",
			have: &Distribution{
				Age: &DistributionCfg{
					Age0to10: AgeData{
						Weight: 100,
						id:     0,
					},
				},
			},
			want: want{
				olders:  10,
				yungest: 0,
			},
		},
		{
			name: "10-20",
			have: &Distribution{
				Age: &DistributionCfg{
					Age10to20: AgeData{
						Weight: 100,
						id:     1,
					},
				},
			},
			want: want{
				olders:  20,
				yungest: 10,
			},
		},
		{
			name: "20-30",
			have: &Distribution{
				Age: &DistributionCfg{
					Age20to30: AgeData{
						Weight: 100,
						id:     2,
					},
				},
			},
			want: want{
				olders:  30,
				yungest: 20,
			},
		},
		{
			name: "30-40",
			have: &Distribution{
				Age: &DistributionCfg{
					Age30to40: AgeData{
						Weight: 100,
						id:     3,
					},
				},
			},
			want: want{
				olders:  40,
				yungest: 30,
			},
		},
		{
			name: "40-50",
			have: &Distribution{
				Age: &DistributionCfg{
					Age40to50: AgeData{
						Weight: 100,
						id:     4,
					},
				},
			},
			want: want{
				olders:  50,
				yungest: 40,
			},
		},
		{
			name: "50-60",
			have: &Distribution{
				Age: &DistributionCfg{
					Age50to60: AgeData{
						Weight: 100,
						id:     5,
					},
				},
			},
			want: want{
				olders:  60,
				yungest: 50,
			},
		},
		{
			name: "60-70",
			have: &Distribution{
				Age: &DistributionCfg{
					Age60to70: AgeData{
						Weight: 100,
						id:     6,
					},
				},
			},
			want: want{
				olders:  70,
				yungest: 60,
			},
		},
		{
			name: "70-80",
			have: &Distribution{
				Age: &DistributionCfg{
					Age70to80: AgeData{
						Weight: 100,
						id:     7,
					},
				},
			},
			want: want{
				olders:  80,
				yungest: 70,
			},
		},
		{
			name: "80-90",
			have: &Distribution{
				Age: &DistributionCfg{
					Age80to90: AgeData{
						Weight: 100,
						id:     8,
					},
				},
			},
			want: want{
				olders:  90,
				yungest: 80,
			},
		},
		{
			name: "90-100",
			have: &Distribution{
				Age: &DistributionCfg{
					Age90to100: AgeData{
						Weight: 100,
						id:     9,
					},
				},
			},
			want: want{
				olders:  100,
				yungest: 90,
			},
		},
		{
			name: "100-110",
			have: &Distribution{
				Age: &DistributionCfg{
					Age100to110: AgeData{
						Weight: 100,
						id:     10,
					},
				},
			},
			want: want{
				olders:  110,
				yungest: 100,
			},
		},
		{
			name: "0-20",
			have: &Distribution{
				Age: &DistributionCfg{
					Age0to10: AgeData{
						Weight: 100,
						id:     0,
					},
					Age10to20: AgeData{
						Weight: 100,
						id:     2,
					},
				},
			},
			want: want{
				olders:  20,
				yungest: 0,
			},
		},
	}

	for _, tt := range tts {
		got, err := tt.have.createYears(1000)
		if err != nil {
			t.Error(err)
		}

		for _, i := range got {
			age := time.Now().Year() - i
			if age > tt.want.olders || age < tt.want.yungest {
				t.Errorf("Name: %q, Error, %d is not int the interval", tt.name, i)
			}
		}
	}
}

func TestRandomHumans(t *testing.T) {
	c, _ := New(&Config{
		DistrubutionCFG: &DistributionCfg{
			Age0to10: AgeData{
				Weight: 100,
				id:     0,
			},
		},
	})
	randHumans, err := c.Distribution.RandomHumans(50)
	if err != nil {
		t.Error("randHumans error", err)
	}

	if randHumans[0].Firstname == randHumans[1].Firstname {
		t.Error("RandHumans has the same firstname", randHumans[0].Firstname, randHumans[1].Firstname)
	}
	for _, rh := range randHumans {
		if rh.Age < 0 {
			t.Error("Age is negative", rh.Age, rh.SocialSecurityNumber.Swedish10.Complete)
		}
	}
}

func TestGenderDistribution(t *testing.T) {
	c, _ := New(&Config{
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

	genders := map[string]float64{}
	n := 1000
	rh, err := c.Distribution.RandomHumans(n)
	if err != nil {
		t.Error(err)
	}
	for _, r := range rh {
		genders[r.Gender.General]++
	}

	if math.Abs(genders[GenderFemale]-genders[GenderMale]) > float64(n)*float64(0.1) {
		t.Errorf("Error, the gender distribution is too asymmetric, female:%f, male:%f", genders[GenderFemale], genders[GenderMale])
	}
}
