// (c) Alexey Shiryavsky. Licensed under the BSD license (see LICENSE).

// Package ru provides holiday definitions for the Russian Federation.
package ru

import (
	"time"

	"github.com/uchti/cal"
	"github.com/uchti/cal/aa"
)

var (
	// Holiday moving rules
	// В соответствии с частью 2 статьи 112 Трудового кодекса Российской Федерации при совпадении выходного и
	//  нерабочего праздничного дней выходной день переносится на следующий после праздничного рабочий день.
	// In accordance with part 2 of article 112 of the Labor Code of the Russian Federation,
	//  if a day off and a non-working holiday coincide, the day off is transferred to the next day after the holiday.
	weekendAlt = []cal.AltDay{
		{Day: time.Saturday, Offset: 2},
		{Day: time.Sunday, Offset: 1},
	}

	// NewYear represents New Year's Day
	NewYear = aa.NewYear.Clone(&cal.Holiday{
		Name: "New Year's Day",
		Type: cal.ObservancePublic,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
	})

	// NewYearHolidays represents typical long russian New Year Holidays
	NewYearHolidays = func() (result []*cal.Holiday) {
		newYearHolidayNumbers := []int{2, 3, 4, 5, 6}
		for _, day := range newYearHolidayNumbers {
			result = append(result, &cal.Holiday{
				Name:     "New Year Holiday",
				Type:     cal.ObservancePublic,
				Month:    time.January,
				Day:      day,
				Observed: weekendAlt,
				Func:     cal.CalcDayOfMonth,
			})
		}
		return
	}()

	// OrthodoxChristmasDay represents Orthodox Christmas Day on 7-jan
	OrthodoxChristmasDay = &cal.Holiday{
		Name:     "Orthodox Christmas Day",
		Type:     cal.ObservanceReligious,
		Month:    time.January,
		Day:      7,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// DefenderOfFatherlandDay represents Defender of the Fatherland Day on the 23-feb
	DefenderOfFatherlandDay = &cal.Holiday{
		Name:     "Defender of the Fatherland Day",
		Type:     cal.ObservancePublic,
		Month:    time.February,
		Day:      23,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// InternationalWomenDay represents International Women's Day on the 8-march
	InternationalWomenDay = &cal.Holiday{
		Name:     "International Women's Day",
		Type:     cal.ObservancePublic,
		Month:    time.March,
		Day:      8,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// SpringAndLaborDay represents Spring and Labor Day on 1-May
	SpringAndLaborDay = aa.WorkersDay.Clone(&cal.Holiday{
		Name:     "Spring and Labor Day",
		Type:     cal.ObservancePublic,
		Observed: weekendAlt,
	})

	// VictoryDay represents Victory Day on the 9-may
	VictoryDay = &cal.Holiday{
		Name:     "Victory Day",
		Type:     cal.ObservancePublic,
		Month:    time.May,
		Day:      9,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// RussiaDay represents Russia Day on the 12-june
	RussiaDay = &cal.Holiday{
		Name:     "Russia Day",
		Type:     cal.ObservancePublic,
		Month:    time.June,
		Day:      12,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// UnityDay represents UnityDay on the 4-nov
	UnityDay = &cal.Holiday{
		Name:     "Unity Day",
		Type:     cal.ObservancePublic,
		Month:    time.November,
		Day:      4,
		Observed: weekendAlt,
		Func:     cal.CalcDayOfMonth,
	}

	// Holidays provides a list of the standard national holidays
	Holidays = append([]*cal.Holiday{
		NewYear,
		OrthodoxChristmasDay,
		DefenderOfFatherlandDay,
		InternationalWomenDay,
		SpringAndLaborDay,
		VictoryDay,
		RussiaDay,
		UnityDay,
	}, NewYearHolidays...)

	//TODO: Re-check winter holidays
)
