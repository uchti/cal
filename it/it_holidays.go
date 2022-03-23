// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package it provides holiday definitions for Italy.
package it

import (
	"time"

	"github.com/uchti/cal"
	"github.com/uchti/cal/aa"
)

var (
	// Capodanno represents New Year's Day on 1-Jan
	Capodanno = aa.NewYear.Clone(&cal.Holiday{Name: "Capodanno", Type: cal.ObservancePublic})

	// Epifania represents Epipany on 6-Jan
	Epifania = aa.Epiphany.Clone(&cal.Holiday{Name: "Epifania", Type: cal.ObservancePublic})

	// Pasquetta represents Easter Monday on the day after Easter
	Pasquetta = aa.EasterMonday.Clone(&cal.Holiday{Name: "Pasquetta", Type: cal.ObservancePublic})

	// FestaDellaLiberazione represents Liberation Day on 25-Apr
	FestaDellaLiberazione = &cal.Holiday{
		Name:  "Festa della Liberazione",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	// FestaDelLavoro represents Labour Day on 1-May
	FestaDelLavoro = aa.WorkersDay.Clone(&cal.Holiday{Name: "Festa del Lavoro", Type: cal.ObservancePublic})

	// FestaDellaRepubblica represents Republic Day on 2-Jun
	FestaDellaRepubblica = &cal.Holiday{
		Name:  "Festa della Repubblica",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// Assunzione represents Assumption of Mary on 15-Aug
	Assunzione = aa.AssumptionOfMary.Clone(&cal.Holiday{Name: "Assunzione", Type: cal.ObservancePublic})

	// TuttiISanti represents All Saints' Day on 1-Nov
	TuttiISanti = aa.AllSaintsDay.Clone(&cal.Holiday{Name: "Tutti i santi", Type: cal.ObservancePublic})

	// Immacolata represents Immaculate Conception on 8-Dec
	Immacolata = aa.ImmaculateConception.Clone(&cal.Holiday{Name: "Immacolata Concezione", Type: cal.ObservancePublic})

	// Natale represents Christmas Day on 25-Dec
	Natale = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Natale", Type: cal.ObservancePublic})

	// SantoStefano represents Saint Stephen's Day on 26-Dec
	SantoStefano = aa.ChristmasDay2.Clone(&cal.Holiday{Name: "Santo Stefano", Type: cal.ObservancePublic})

	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		Capodanno,
		Epifania,
		Pasquetta,
		FestaDellaLiberazione,
		FestaDelLavoro,
		FestaDellaRepubblica,
		Assunzione,
		TuttiISanti,
		Immacolata,
		Natale,
		SantoStefano,
	}
)
