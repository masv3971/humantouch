package humantouch

import (
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/masv3971/humantouch/data"
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
	}{
		{
			name:    "FirstnameFemale",
			tempMap: make(map[string]bool),
			list:    FirstnamesFemale,
		},
		{
			name:    "FirstnameMale",
			tempMap: make(map[string]bool),
			list:    FirstnamesMale,
		},
		{
			name:    "NinFemale",
			tempMap: make(map[string]bool),
			list:    data.F2013,
		},
		{
			name:    "NinMale",
			tempMap: make(map[string]bool),
			list:    data.M2013,
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			want := []string{}
			for _, l := range tt.list {
				if _, ok := tt.tempMap[l]; !ok {
					want = append(want, l)
				}
			}
			assert.Equal(t, want, tt.list)
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

func TestNINs(t *testing.T) {
	totalCount := 0
	tts := []struct {
		name   string
		gender string
		list   [][]string
	}{
		{
			name:   "Female",
			gender: GenderFemale,
			list: [][]string{
				data.F1890,
				data.F1891,
				data.F1892,
				data.F1893,
				data.F1894,
				data.F1895,
				data.F1896,
				data.F1897,
				data.F1898,
				data.F1899,
				data.F1900,
				data.F1901,
				data.F1902,
				data.F1903,
				data.F1904,
				data.F1905,
				data.F1906,
				data.F1907,
				data.F1908,
				data.F1909,
				data.F1910,
				data.F1911,
				data.F1912,
				data.F1920,
				data.F1921,
				data.F1922,
				data.F1923,
				data.F1924,
				data.F1925,
				data.F1926,
				data.F1927,
				data.F1928,
				data.F1929,
				data.F1930,
				data.F1931,
				data.F1932,
				data.F1933,
				data.F1934,
				data.F1935,
				data.F1936,
				data.F1937,
				data.F1938,
				data.F1939,
				data.F1940,
				data.F1941,
				data.F1942,
				data.F1943,
				data.F1944,
				data.F1945,
				data.F1946,
				data.F1947,
				data.F1948,
				data.F1949,
				data.F1950,
				data.F1951,
				data.F1952,
				data.F1953,
				data.F1954,
				data.F1955,
				data.F1956,
				data.F1957,
				data.F1958,
				data.F1959,
				data.F1960,
				data.F1961,
				data.F1962,
				data.F1963,
				data.F1964,
				data.F1965,
				data.F1966,
				data.F1967,
				data.F1968,
				data.F1970,
				data.F1971,

				data.F1975,

				data.F1979,
				data.F1980,
				data.F1981,
				data.F1982,
				data.F1983,
				data.F1984,
				data.F1985,
				data.F1986,
				data.F1987,
				data.F1989,
				data.F1990,
				data.F1991,
				data.F1992,
				data.F1993,
				data.F1994,
				data.F1995,
				data.F1996,
				data.F1997,
				data.F1998,
				data.F1999,
				data.F2000,
				data.F2001,
				data.F2002,
				data.F2003,
				data.F2004,
				data.F2005,
				data.F2006,
				data.F2007,
				data.F2008,
				data.F2009,
				data.F2010,
				data.F2011,
				data.F2012,
				data.F2013,
				data.F2014,
			},
		},
		{
			name:   "male",
			gender: GenderMale,
			list: [][]string{
				data.M1890,
				data.M1891,
				data.M1892,
				data.M1893,
				data.M1894,
				data.M1895,
				data.M1896,
				data.M1897,
				data.M1898,
				data.M1899,
				data.M1900,
				data.M1901,
				data.M1902,
				data.M1903,
				data.M1904,
				data.M1905,
				data.M1906,
				data.M1907,
				data.M1908,
				data.M1909,
				data.M1910,
				data.M1911,
				data.M1912,
				data.M1920,
				data.M1921,
				data.M1922,
				data.M1923,
				data.M1924,
				data.M1925,
				data.M1926,
				data.M1927,
				data.M1928,
				data.M1929,
				data.M1930,
				data.M1931,
				data.M1932,
				data.M1933,
				data.M1934,
				data.M1935,
				data.M1936,
				data.M1937,
				data.M1938,
				data.M1939,
				data.M1940,
				data.M1941,
				data.M1942,
				data.M1943,
				data.M1944,
				data.M1945,
				data.M1946,
				data.M1947,
				data.M1948,
				data.M1949,
				data.M1950,
				data.M1951,
				data.M1952,
				data.M1953,
				data.M1954,
				data.M1955,
				data.M1956,
				data.M1957,
				data.M1958,
				data.M1959,
				data.M1960,
				data.M1961,
				data.M1962,
				data.M1963,
				data.M1964,
				data.M1965,
				data.M1966,
				data.M1967,
				data.M1968,
				data.M1970,
				data.M1971,

				data.M1975,

				data.M1979,
				data.M1980,
				data.M1981,
				data.M1982,
				data.M1983,
				data.M1984,
				data.M1985,
				data.M1986,
				data.M1987,
				data.M1989,
				data.M1990,
				data.M1991,
				data.M1992,
				data.M1993,
				data.M1994,
				data.M1995,
				data.M1996,
				data.M1997,
				data.M1998,
				data.M1999,
				data.M2000,
				data.M2001,
				data.M2002,
				data.M2003,
				data.M2004,
				data.M2005,
				data.M2006,
				data.M2007,
				data.M2008,
				data.M2009,
				data.M2010,
				data.M2011,
				data.M2012,
				data.M2013,
				data.M2014,
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			count := 0

			for _, list := range tt.list {
				count += len(list)
				for _, nin := range list {
					gender, err := genderFromNIN(nin)
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if gender != tt.gender {
						t.Errorf("Wrong gender for nin: %q in list: %q", nin, tt.name)
					}
				}
			}
			totalCount += count

			t.Logf("%s: %d", tt.gender, count)
		})
	}
	t.Logf("Total count: %d", totalCount)
}
