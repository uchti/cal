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
		{NewYear, 2020, d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1)},
		//
		{MakarSankrantiDay, 2020, d(2020, 1, 14)},
		{MakarSankrantiDay, 2021, d(2021, 1, 14)},
		{MakarSankrantiDay, 2022, d(2022, 1, 14)},
		//
		{RepublicDay, 2020, d(2020, 1, 26)},
		{RepublicDay, 2021, d(2021, 1, 26)},
		{RepublicDay, 2022, d(2022, 1, 26)},
		//
		{IndependenceDay, 2020, d(2020, 8, 15)},
		{IndependenceDay, 2021, d(2021, 8, 15)},
		{IndependenceDay, 2022, d(2022, 8, 15)},
		//
		{GandhiJayanti, 2020, d(2020, 10, 2)},
		{GandhiJayanti, 2021, d(2021, 10, 2)},
		{GandhiJayanti, 2022, d(2022, 10, 2)},
		//
		{LabourDay, 2020, d(2020, 5, 1)},
		{LabourDay, 2021, d(2021, 5, 1)},
		{LabourDay, 2022, d(2022, 5, 1)},
		//
		{ChristmasDay, 2020, d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25)},
	}

	for _, test := range tests {
		gotAct, _ := test.h.Calc(test.y)
		if !gotAct.Equal(test.date) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.date.String())
		}
	}
}
