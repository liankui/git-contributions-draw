package draw

import (
	"reflect"
	"testing"
)

func Test_getFirstSunday(t *testing.T) {
	tests := []struct {
		name string
		year int
		want string
	}{
		{name: "", year: 2017, want: "2017-01-01"},
		{name: "", year: 2018, want: "2018-01-07"},
		{name: "", year: 2019, want: "2019-01-06"},
		{name: "", year: 2020, want: "2020-01-05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstSunday(tt.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFirstSunday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastSaturday(t *testing.T) {
	tests := []struct {
		name string
		year int
		want string
	}{
		{name: "", year: 2017, want: "2017-12-30"},
		{name: "", year: 2018, want: "2018-12-29"},
		{name: "", year: 2019, want: "2019-12-28"},
		{name: "", year: 2020, want: "2020-12-26"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastSaturday(tt.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLastSaturday() = %v, want %v", got, tt.want)
			}
		})
	}
}
