package humantouch

import (
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestFirstnamesFemale_uniq(t *testing.T) {
	temp := make(map[string]bool)
	want := []string{}
	got := FirstnamesFemale

	for _, fnf := range FirstnamesFemale {
		if _, ok := temp[fnf]; !ok {
			temp[fnf] = true
			want = append(want, fnf)
		}
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestLastnames_uniq(t *testing.T) {
	temp := make(map[string]bool)
	want := []string{}
	got := Lastnames

	for _, fnf := range Lastnames {
		if _, ok := temp[fnf]; !ok {
			temp[fnf] = true
			want = append(want, fnf)
		}
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestUniqList(t *testing.T) {
	tts := []struct {
		name    string
		tempMap map[string]bool
		list    []string
		want    []string
	}{
		{
			name:    "FirstnameFemale",
			tempMap: make(map[string]bool),
			list:    FirstnamesFemale,
			want:    []string{},
		},
		{
			name:    "FirstnameMale",
			tempMap: make(map[string]bool),
			list:    FirstnamesMale,
			want:    []string{},
		},
		{
			name:    "NinFemale",
			tempMap: make(map[string]bool),
			list:    SkatteverketTestSocialSecurityNumbersFemale,
			want:    []string{},
		},
		{
			name:    "NinMale",
			tempMap: make(map[string]bool),
			list:    SkatteverketTestSocialSecurityNumbersMale,
			want:    []string{},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			for _, l := range tt.list {
				if _, ok := tt.tempMap[l]; !ok {
					tt.want = append(tt.want, l)
				}
			}
			assert.Equal(t, tt.want, tt.list)
		})
	}
}

func TestFirstnamesFemale_order(t *testing.T) {
	want := FirstnamesFemale
	sort.Strings(want)
	got := FirstnamesFemale

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestLastnames_order(t *testing.T) {
	want := Lastnames
	sort.Strings(want)
	got := Lastnames

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFirstnamesMale_order(t *testing.T) {
	want := FirstnamesMale
	sort.Strings(want)
	got := FirstnamesMale

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFirstnamesFemale_whitespaces(t *testing.T) {
	for _, f := range FirstnamesFemale {
		if strings.HasPrefix(f, " ") {
			t.Error("Error, FirstnameFemale has a name that starts with a whitespace", f)
		}
		if strings.HasSuffix(f, " ") {
			t.Error("Error, a name ends with a whitespace", f)
		}
	}
}

func TestFirstnamesMale_whitespaces(t *testing.T) {
	for _, f := range FirstnamesMale {
		if strings.HasPrefix(f, " ") {
			t.Error("Error, a name starts with a whitespace", f)
		}
		if strings.HasSuffix(f, " ") {
			t.Error("Error, a name ends with a whitespace", f)
		}
	}
}

func TestLastname_whitespaces(t *testing.T) {
	for _, f := range Lastnames {
		if strings.HasPrefix(f, " ") {
			t.Error("Error, a name starts with a whitespace", f)
		}
		if strings.HasSuffix(f, " ") {
			t.Error("Error, a name ends with a whitespace", f)
		}
	}
}

func TestFirstnamesFemale_short(t *testing.T) {
	for _, name := range FirstnamesFemale {
		if len(name) < 2 {
			t.Errorf("Name: %q in FirstnamesFemale to short", name)
		}
	}
}

func TestFirstnamesMale_short(t *testing.T) {
	for _, name := range FirstnamesMale {
		if len(name) < 2 {
			t.Errorf("Name: %q in FirstnamesMale to short", name)
		}
	}
}

func TestLastnames_short(t *testing.T) {
	for _, name := range Lastnames {
		if len(name) < 2 {
			t.Errorf("Name: %q in Lastnames to short", name)
		}
	}
}

func TestFirstnamesMale_family(t *testing.T) {
	for _, name := range FirstnamesMale {
		if strings.HasSuffix(name, "sson") || strings.HasSuffix(name, "lund") || strings.HasSuffix(name, "strand") || strings.HasSuffix(name, "kvist") || strings.HasSuffix(name, "qvist") || strings.HasSuffix(name, "sén") {
			t.Errorf("Name: %q maybe a family name", name)
		}
	}
}

func TestCountNames(t *testing.T) {
	t.Log("FirstnamesFemale:", len(FirstnamesFemale), "Possible combination:", len(FirstnamesFemale)*len(Lastnames))
	t.Log("FirstnamesMale", len(FirstnamesMale), "Possible combination:", len(FirstnamesMale)*len(Lastnames))
	t.Log("Lastnames", len(Lastnames))

}

func TestSkatteverketTestSocialSecurityNumber(t *testing.T) {
	tts := []struct {
		name string
		list []string
		want string
	}{
		{
			name: "Female list",
			list: SkatteverketTestSocialSecurityNumbersFemale,
			want: GenderFemale,
		},
		{
			name: "Male list",
			list: SkatteverketTestSocialSecurityNumbersMale,
			want: GenderMale,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			for _, nin := range tt.list {
				gender, err := genderFromNIN(nin)
				if !assert.NoError(t, err) {
					t.FailNow()
				}

				if gender != tt.want {
					t.Errorf("Wrong gender for nin: %q in list: %q", nin, tt.name)
				}
			}
		})
	}
}
