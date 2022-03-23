// (c) Denis Chumachenko. Licensed under the BSD license (see LICENSE).

// Package ie provides holiday definitions for the Colombia.
package in

import (
	"github.com/uchti/cal"
	"time"
)

var (
	NewYearDay = &cal.Holiday{
		Name:  "New Year's Day",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	LaborDay = &cal.Holiday{
		Name:  "Día del Trabajo [Labour Day]",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	IndependenceDay = &cal.Holiday{
		Name:  "Día de la Independencia [Independence Day]",
		Type:  cal.ObservancePublic,
		Month: time.July,
		Day:   20,
		Func:  cal.CalcDayOfMonth,
	}

	BattleOfBoyaca = &cal.Holiday{
		Name:  "Batalla de Boyacá [Battle of Boyacá]",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   7,
		Func:  cal.CalcDayOfMonth,
	}

	ImmaculateConception = &cal.Holiday{
		Name:  "La Inmaculada Concepción [Immaculate Conception]",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   8,
		Func:  cal.CalcDayOfMonth,
	}

	Christmas = &cal.Holiday{
		Name:  "Navidad [Christmas]",
		Type:  cal.ObservanceReligious,
		Month: time.December,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	Epiphany = &cal.Holiday{
		Name:  "Día de los Reyes Magos [Epiphany]",
		Type:  cal.ObservanceReligious,
		Month: time.January,
		Day:   6,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	SaintJosephDay = &cal.Holiday{
		Name:  "Día de San José [Saint Joseph's Day]",
		Type:  cal.ObservanceReligious,
		Month: time.March,
		Day:   19,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	SaintPeterAndSaintPaul = &cal.Holiday{
		Name:  "San Pedro y San Pablo [Saint Peter and Saint Paul]",
		Type:  cal.ObservanceReligious,
		Month: time.January,
		Day:   29,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	AssumptionOfMary = &cal.Holiday{
		Name:  "La Asunción [Assumption of Mary]",
		Type:  cal.ObservanceReligious,
		Month: time.August,
		Day:   15,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	ColumbusDay = &cal.Holiday{
		Name:  "Día de la Raza [Columbus Day]",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   12,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	AllSaintsDay = &cal.Holiday{
		Name:  "Día de Todos los Santos [All Saint's Day]",
		Type:  cal.ObservanceReligious,
		Month: time.November,
		Day:   1,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	IndependenceOfCartagena = &cal.Holiday{
		Name:  "Independencia de Cartagena [Independence of Cartagena]",
		Type:  cal.ObservanceReligious,
		Month: time.November,
		Day:   11,
		Observed: []cal.AltDay{
			{Day: time.Monday, Offset: 7},
		},
		Func: cal.CalcDayOfMonth,
	}

	Holidays = []*cal.Holiday{
		NewYearDay,
		LaborDay,
		IndependenceDay,
		BattleOfBoyaca,
		ImmaculateConception,
		Christmas,
		Epiphany,
		SaintJosephDay,
		SaintPeterAndSaintPaul,
		AssumptionOfMary,
		IndependenceDay,
		ColumbusDay,
		AllSaintsDay,
		IndependenceOfCartagena,
	}

	// TODO Support Easter year in Colombia
	//# Holidays based on Easter
	//
	//# Maundy Thursday
	//self[
	//easter(year) + rd(weekday=TH(-1))
	//] = "Jueves Santo [Maundy Thursday]"
	//
	//# Good Friday
	//self[easter(year) + rd(weekday=FR(-1))] = "Viernes Santo [Good Friday]"
	//
	//# Holidays based on Easter but are observed the following monday
	//# (unless they occur on a monday)
	//
	//# Ascension of Jesus
	//name = "Ascensión del señor [Ascension of Jesus]"
	//hdate = easter(year) + rd(days=+39)
	//if hdate.weekday() == MON or not self.observed:
	//self[hdate] = name
	//else:
	//self[hdate + rd(weekday=MO)] = name + "(Observed)"
	//
	//# Corpus Christi
	//name = "Corpus Christi [Corpus Christi]"
	//hdate = easter(year) + rd(days=+60)
	//if hdate.weekday() == MON or not self.observed:
	//self[hdate] = name
	//else:
	//self[hdate + rd(weekday=MO)] = name + "(Observed)"
	//
	//# Sacred Heart
	//name = "Sagrado Corazón [Sacred Heart]"
	//hdate = easter(year) + rd(days=+68)
	//if hdate.weekday() == MON or not self.observed:
	//self[hdate] = name
	//else:
	//self[hdate + rd(weekday=MO)] = name + "(Observed)"

)
