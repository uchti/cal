// (c) Alexey Shiryavsky. Licensed under the BSD license (see LICENSE).

package ru

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
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{NewYear, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{NewYear, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{NewYear, 2017, d(2017, 1, 1), d(2017, 1, 2)},
		{NewYear, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{NewYear, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{NewYear, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1), d(2021, 12, 31)},
		// {NewYear, 2022, d(2022, 1, 2), d(2022, 1, 1)},

		{NewYearHolidays[0], 2021, d(2021, 1, 2), d(2021, 1, 4)},
		{NewYearHolidays[4], 2021, d(2021, 1, 6), d(2021, 1, 6)},

		{NewYearHolidays[0], 2022, d(2022, 1, 2), d(2022, 1, 3)},
		{NewYearHolidays[4], 2022, d(2022, 1, 6), d(2022, 1, 6)},

		{SpringAndLaborDay, 2021, d(2021, 5, 1), d(2021, 5, 3)},

		{VictoryDay, 2021, d(2021, 5, 9), d(2021, 5, 10)},

		{RussiaDay, 2021, d(2021, 6, 12), d(2021, 6, 14)},

		{UnityDay, 2021, d(2021, 11, 4), d(2021, 11, 4)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}
