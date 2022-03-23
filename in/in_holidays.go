// (c) Denis Chumachenko. Licensed under the BSD license (see LICENSE).

// Package ie provides holiday definitions for the India.
package in

import (
	"github.com/uchti/cal"
	"github.com/uchti/cal/aa"
	"time"
)

var (
	NewYear = aa.NewYear.Clone(&cal.Holiday{})

	MakarSankrantiDay = &cal.Holiday{
		Name:  "Makar Sankranti / Pongal",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   14,
		Func:  cal.CalcDayOfMonth,
	}

	RepublicDay = &cal.Holiday{
		Name:      "Republic Day",
		Type:      cal.ObservancePublic,
		Month:     time.January,
		Day:       26,
		StartYear: 1950,
		Func:      cal.CalcDayOfMonth,
	}

	IndependenceDay = &cal.Holiday{
		Name:      "Independence Day",
		Type:      cal.ObservancePublic,
		Month:     time.August,
		Day:       15,
		StartYear: 1947,
		Func:      cal.CalcDayOfMonth,
	}

	GandhiJayanti = &cal.Holiday{
		Name:  "Gandhi Jayanti",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	LabourDay = &cal.Holiday{
		Name:  "Labour Day",
		Type:  cal.ObservancePublic,
		Month: time.May,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	ChristmasDay = &cal.Holiday{
		Name:  "Christmas",
		Type:  cal.ObservancePublic,
		Month: time.December,
		Day:   25,
		Func:  cal.CalcDayOfMonth,
	}

	//// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		ChristmasDay,
		MakarSankrantiDay,
		RepublicDay,
		IndependenceDay,
		GandhiJayanti,
		LabourDay,
		ChristmasDay,
	}

	//# GJ: Gujarat
	//if self.subdiv == "GJ":
	//self[date(year, JAN, 14)] = "Uttarayan"
	//self[date(year, MAY, 1)] = "Gujarat Day"
	//self[date(year, OCT, 31)] = "Sardar Patel Jayanti"
	//
	//if self.subdiv == "BR":
	//self[date(year, MAR, 22)] = "Bihar Day"
	//
	//if self.subdiv == "RJ":
	//self[date(year, MAR, 30)] = "Rajasthan Day"
	//self[date(year, JUN, 15)] = "Maharana Pratap Jayanti"
	//
	//if self.subdiv == "OR":
	//self[date(year, APR, 1)] = "Odisha Day (Utkala Dibasa)"
	//self[date(year, APR, 15)] = (
	//"Maha Vishuva Sankranti / Pana" " Sankranti"
	//)
	//
	//if self.subdiv in (
	//"OR",
	//"AP",
	//"BR",
	//"WB",
	//"KL",
	//"HR",
	//"MH",
	//"UP",
	//"UK",
	//"TN",
	//):
	//self[date(year, APR, 14)] = "Dr. B. R. Ambedkar's Jayanti"
	//
	//if self.subdiv == "TN":
	//self[date(year, APR, 14)] = "Puthandu (Tamil New Year)"
	//self[date(year, APR, 15)] = "Puthandu (Tamil New Year)"
	//
	//if self.subdiv == "WB":
	//self[date(year, APR, 14)] = "Pohela Boishakh"
	//self[date(year, APR, 15)] = "Pohela Boishakh"
	//self[date(year, MAY, 9)] = "Rabindra Jayanti"
	//
	//if self.subdiv == "AS":
	//self[date(year, APR, 15)] = "Bihu (Assamese New Year)"
	//
	//if self.subdiv == "MH":
	//self[date(year, MAY, 1)] = "Maharashtra Day"
	//self[date(year, OCT, 15)] = "Dussehra"
	//
	//if self.subdiv == "SK":
	//self[date(year, MAY, 16)] = "Annexation Day"
	//
	//if self.subdiv == "KA":
	//self[date(year, NOV, 1)] = "Karnataka Rajyotsava"
	//
	//if self.subdiv == "AP":
	//self[date(year, NOV, 1)] = "Andhra Pradesh Foundation Day"
	//
	//if self.subdiv == "HR":
	//self[date(year, NOV, 1)] = "Haryana Foundation Day"
	//
	//if self.subdiv == "MP":
	//self[date(year, NOV, 1)] = "Madhya Pradesh Foundation Day"
	//
	//if self.subdiv == "KL":
	//self[date(year, NOV, 1)] = "Kerala Foundation Day"
	//
	//if self.subdiv == "CG":
	//self[date(year, NOV, 1)] = "Chhattisgarh Foundation Day"
	//
	//if self.subdiv == "TS":
	//self[date(year, OCT, 6)] = "Bathukamma Festival"
	//self[date(year, APR, 6)] = "Eid al-Fitr"

)
