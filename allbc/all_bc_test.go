// (c) Denis Chumachenko. Licensed under the BSD license (see LICENSE).

package allbc

import (
	"testing"
	"time"
)

func TestIsWorkDayByCountry(t *testing.T) {
	InitAllBussinesCalendar()
	tests := []struct {
		country string
		dt      time.Time
		want    bool
	}{
		{"RU", time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{"RU", time.Date(2021, 3, 8, 12, 0, 0, 0, time.UTC), false},
		{"RU", time.Date(2021, 9, 10, 12, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		got, err := IsWorkDayByCountry(test.country, test.dt)
		if err != nil {
			t.Errorf("got: %t; err: %v; want: %t (%s)", got, err, test.want, test.dt)
		} else {
			if got != test.want {
				t.Errorf("got: %t; want: %t (%s)", got, test.want, test.dt)
			}
		}

	}
}
