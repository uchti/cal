// (c) Denis Chumachenko. Licensed under the BSD license (see LICENSE).

package in

import (
	"testing"
	"time"

	"github.com/uchti/cal"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func TestHolidays(t *testing.T) {
	tests := []struct {
		h    *cal.Holiday
		y    int
		date time.Time
	}{
		{NewYearDay, 2020, d(2020, 1, 1)},
		{NewYearDay, 2021, d(2021, 1, 1)},
		{NewYearDay, 2022, d(2022, 1, 1)},
		//
		{IndependenceDay, 2020, d(2020, 7, 20)},
		{IndependenceDay, 2021, d(2021, 7, 20)},
		{IndependenceDay, 2022, d(2022, 7, 20)},
		//
		{LaborDay, 2020, d(2020, 5, 1)},
		{LaborDay, 2021, d(2021, 5, 1)},
		{LaborDay, 2022, d(2022, 5, 1)},
		//
		{Christmas, 2020, d(2020, 12, 25)},
		{Christmas, 2021, d(2021, 12, 25)},
		{Christmas, 2022, d(2022, 12, 25)},
	}

	for _, test := range tests {
		gotAct, _ := test.h.Calc(test.y)
		if !gotAct.Equal(test.date) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.date.String())
		}
	}
}
