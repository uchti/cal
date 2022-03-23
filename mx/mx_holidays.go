// (c) Denis Chumachenko. Licensed under the BSD license (see LICENSE).

// Package ie provides holiday definitions for the Mexico.
package mx

import (
	"github.com/uchti/cal"
	"time"
)

var (
	NewYearDay = &cal.Holiday{
		Name:  "Año Nuevo [New Year's Day]",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   1,
		Observed: []cal.AltDay{
			{Day: time.Sunday, Offset: +1},
		},
		Func: cal.CalcDayOfMonth,
	}

	ConstitutionDay = &cal.Holiday{
		Name:      "Día de la Constitución [Constitution Day]",
		Type:      cal.ObservancePublic,
		Month:     time.February,
		Day:       1,
		StartYear: 2007,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: +1},
		},
		Func: cal.CalcDayOfMonth,
	}
	ConstitutionDayBefore = &cal.Holiday{
		Name:      "Día de la Constitución [Constitution Day]",
		Type:      cal.ObservancePublic,
		Month:     time.February,
		Day:       5,
		StartYear: 1917,
		EndYear:   2007,
		Func:      cal.CalcDayOfMonth,
	}

	BenitoJuarezBirthday = &cal.Holiday{
		Name:      "2007",
		Type:      cal.ObservancePublic,
		Month:     time.March,
		Day:       1,
		StartYear: 2007,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: +3},
		},
		Func: cal.CalcDayOfMonth,
	}

	BenitoJuarezBirthdayBefore = &cal.Holiday{
		Name:      "Natalicio de Benito Juárez [Benito Juárez's birthday]",
		Type:      cal.ObservancePublic,
		Month:     time.March,
		Day:       21,
		StartYear: 1917,
		EndYear:   2007,
		Func:      cal.CalcDayOfMonth,
	}

	LaborDay = &cal.Holiday{
		Name:      "Día del Trabajo [Labour Day]",
		Type:      cal.ObservancePublic,
		Month:     time.May,
		Day:       1,
		StartYear: 1923,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: -3},
			{Day: time.Sunday, Offset: +1},
		},
		Func: cal.CalcDayOfMonth,
	}

	IndependenceDay = &cal.Holiday{
		Name:  "Día de la Independencia [Independence Day]",
		Type:  cal.ObservancePublic,
		Month: time.September,
		Day:   16,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: +1},
		},
		Func: cal.CalcDayOfMonth,
	}

	RevolutionDay = &cal.Holiday{
		Name:      "Día de la Revolución [Revolution Day]",
		Type:      cal.ObservancePublic,
		Month:     time.November,
		Day:       1,
		StartYear: 2007,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: +3},
		},
		Func: cal.CalcDayOfMonth,
	}

	RevolutionDayBefore = &cal.Holiday{
		Name:      "Día de la Revolución [Revolution Day]",
		Type:      cal.ObservancePublic,
		Month:     time.November,
		Day:       20,
		StartYear: 1917,
		EndYear:   2007,
		Func:      cal.CalcDayOfMonth,
	}

	ChangeOfFederalGovernment = &cal.Holiday{
		Name:      "Transmisión del Poder Ejecutivo Federal [Change of Federal Government]",
		Type:      cal.ObservancePublic,
		Month:     time.December,
		Day:       1,
		StartYear: 1970,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: +1},
		},
		Func: cal.CalcDayOfMonth,
	}

	Christmas = &cal.Holiday{
		Name:  "Navidad [Christmas]",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   25,
		Observed: []cal.AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: +1},
		},
		Func: cal.CalcDayOfMonth,
	}

	Holidays = []*cal.Holiday{
		NewYearDay,
		ConstitutionDay,
		ConstitutionDayBefore,
		BenitoJuarezBirthday,
		BenitoJuarezBirthdayBefore,
		LaborDay,
		IndependenceDay,
		RevolutionDay,
		RevolutionDayBefore,
		ChangeOfFederalGovernment,
		Christmas,
	}
)
