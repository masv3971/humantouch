package humantouch

import (
	"testing"
	"time"
)

func TestCreateYears(t *testing.T) {
	type want struct {
		olders  int
		yungest int
	}
	tts := []struct {
		name string
		have *Distrubution
		want want
	}{
		{
			name: "0-10",
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
			have: &Distrubution{
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
	randHumans, err := c.Distrubution.RandomHumans(50)
	if err != nil {
		t.Error("randHumans error", err)
	}

	if randHumans[0].Firstname == randHumans[1].Firstname {
		t.Error("RandHumans has the same firstname", randHumans[0].Firstname, randHumans[1].Firstname)
	}
}
