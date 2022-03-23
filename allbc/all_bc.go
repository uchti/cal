// (c) Denis Chumachenko. Licensed under the BSD license (see LICENSE).

package allbc

import (
	"errors"
	"github.com/uchti/cal"
	calAR "github.com/uchti/cal/ar"
	calAT "github.com/uchti/cal/at"
	calAU "github.com/uchti/cal/au"
	calBE "github.com/uchti/cal/be"
	calBR "github.com/uchti/cal/br"
	calCA "github.com/uchti/cal/ca"
	calCH "github.com/uchti/cal/ch"
	calCO "github.com/uchti/cal/co"
	calCZ "github.com/uchti/cal/cz"
	calDE "github.com/uchti/cal/de"
	calDK "github.com/uchti/cal/dk"
	calECB "github.com/uchti/cal/ecb"
	calES "github.com/uchti/cal/es"
	calFR "github.com/uchti/cal/fr"
	calGB "github.com/uchti/cal/gb"
	calIE "github.com/uchti/cal/ie"
	calIN "github.com/uchti/cal/in"
	calIT "github.com/uchti/cal/it"
	calJP "github.com/uchti/cal/jp"
	calLT "github.com/uchti/cal/lt"
	calMX "github.com/uchti/cal/mx"
	calNL "github.com/uchti/cal/nl"
	calNO "github.com/uchti/cal/no"
	calNZ "github.com/uchti/cal/nz"
	calPL "github.com/uchti/cal/pl"
	calRU "github.com/uchti/cal/ru"
	calSE "github.com/uchti/cal/se"
	calSI "github.com/uchti/cal/si"
	calSK "github.com/uchti/cal/sk"
	calUA "github.com/uchti/cal/ua"
	calUS "github.com/uchti/cal/us"
	calZA "github.com/uchti/cal/za"
	"strings"
	"time"
)

//TODO: These countries should be added first
//RO,151620
//EG,130590
//MY,127661
//PH,119973
//PE,106530
//SG,99971
//TW,92149
//CL,88725
//VN,87777
//HU,73688
//DZ,71739
//TR,71173
//EC,69317
//TH,68084
//NG,67507
//GR,63137
//KR,62832
//BD,59440
//ID,54566
//PK,54557
//IQ,52932
//LK,52360
//IL,51981
//PT,51353

var (
	AllBC              = map[string]*cal.BusinessCalendar{}
	ErrAllBDNotInit    = errors.New("All BussinesDays objects was not intit. Call InitAllBussinesCalendar().")
	ErrCountryNotAdded = errors.New("This country is not added to the package.")
)

func InitAllBussinesCalendar() {
	AllBC = make(map[string]*cal.BusinessCalendar)

	AllBC["ar"] = &cal.BusinessCalendar{}
	AllBC["ar"].SetWorkday(time.Monday, true)
	AllBC["ar"].SetWorkday(time.Tuesday, true)
	AllBC["ar"].SetWorkday(time.Wednesday, true)
	AllBC["ar"].SetWorkday(time.Thursday, true)
	AllBC["ar"].SetWorkday(time.Friday, true)
	AllBC["ar"].SetWorkdayStart(9 * time.Hour)
	AllBC["ar"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ar"].Holidays = calAR.Holidays
	AllBC["ar"].Name = "Argentina"

	AllBC["at"] = &cal.BusinessCalendar{}
	AllBC["at"].SetWorkday(time.Monday, true)
	AllBC["at"].SetWorkday(time.Tuesday, true)
	AllBC["at"].SetWorkday(time.Wednesday, true)
	AllBC["at"].SetWorkday(time.Thursday, true)
	AllBC["at"].SetWorkday(time.Friday, true)
	AllBC["at"].SetWorkdayStart(9 * time.Hour)
	AllBC["at"].SetWorkdayEnd(17 * time.Hour)
	AllBC["at"].Holidays = calAT.Holidays
	AllBC["at"].Name = "Austria"

	AllBC["au"] = &cal.BusinessCalendar{}
	AllBC["au"].SetWorkday(time.Monday, true)
	AllBC["au"].SetWorkday(time.Tuesday, true)
	AllBC["au"].SetWorkday(time.Wednesday, true)
	AllBC["au"].SetWorkday(time.Thursday, true)
	AllBC["au"].SetWorkday(time.Friday, true)
	AllBC["au"].SetWorkdayStart(9 * time.Hour)
	AllBC["au"].SetWorkdayEnd(17 * time.Hour)
	AllBC["au"].Holidays = calAU.HolidaysACT
	AllBC["au"].Name = "Australia"

	AllBC["be"] = &cal.BusinessCalendar{}
	AllBC["be"].SetWorkday(time.Monday, true)
	AllBC["be"].SetWorkday(time.Tuesday, true)
	AllBC["be"].SetWorkday(time.Wednesday, true)
	AllBC["be"].SetWorkday(time.Thursday, true)
	AllBC["be"].SetWorkday(time.Friday, true)
	AllBC["be"].SetWorkdayStart(9 * time.Hour)
	AllBC["be"].SetWorkdayEnd(17 * time.Hour)
	AllBC["be"].Holidays = calBE.Holidays
	AllBC["be"].Name = "Belgium"

	AllBC["br"] = &cal.BusinessCalendar{}
	AllBC["br"].SetWorkday(time.Monday, true)
	AllBC["br"].SetWorkday(time.Tuesday, true)
	AllBC["br"].SetWorkday(time.Wednesday, true)
	AllBC["br"].SetWorkday(time.Thursday, true)
	AllBC["br"].SetWorkday(time.Friday, true)
	AllBC["br"].SetWorkdayStart(9 * time.Hour)
	AllBC["br"].SetWorkdayEnd(17 * time.Hour)
	AllBC["br"].Holidays = calBR.Holidays
	AllBC["br"].Name = "Brazil"

	AllBC["ca"] = &cal.BusinessCalendar{}
	AllBC["ca"].SetWorkday(time.Monday, true)
	AllBC["ca"].SetWorkday(time.Tuesday, true)
	AllBC["ca"].SetWorkday(time.Wednesday, true)
	AllBC["ca"].SetWorkday(time.Thursday, true)
	AllBC["ca"].SetWorkday(time.Friday, true)
	AllBC["ca"].SetWorkdayStart(9 * time.Hour)
	AllBC["ca"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ca"].Holidays = calCA.Holidays
	AllBC["ca"].Name = "Canada"

	AllBC["co"] = &cal.BusinessCalendar{}
	AllBC["co"].SetWorkday(time.Monday, true)
	AllBC["co"].SetWorkday(time.Tuesday, true)
	AllBC["co"].SetWorkday(time.Wednesday, true)
	AllBC["co"].SetWorkday(time.Thursday, true)
	AllBC["co"].SetWorkday(time.Friday, true)
	AllBC["co"].SetWorkdayStart(9 * time.Hour)
	AllBC["co"].SetWorkdayEnd(17 * time.Hour)
	AllBC["co"].Holidays = calCO.Holidays
	AllBC["co"].Name = "Colombia"

	AllBC["ch"] = &cal.BusinessCalendar{}
	AllBC["ch"].SetWorkday(time.Monday, true)
	AllBC["ch"].SetWorkday(time.Tuesday, true)
	AllBC["ch"].SetWorkday(time.Wednesday, true)
	AllBC["ch"].SetWorkday(time.Thursday, true)
	AllBC["ch"].SetWorkday(time.Friday, true)
	AllBC["ch"].SetWorkdayStart(9 * time.Hour)
	AllBC["ch"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ch"].Holidays = calCH.Holidays
	AllBC["ch"].Name = "Switzerland"

	AllBC["cz"] = &cal.BusinessCalendar{}
	AllBC["cz"].SetWorkday(time.Monday, true)
	AllBC["cz"].SetWorkday(time.Tuesday, true)
	AllBC["cz"].SetWorkday(time.Wednesday, true)
	AllBC["cz"].SetWorkday(time.Thursday, true)
	AllBC["cz"].SetWorkday(time.Friday, true)
	AllBC["cz"].SetWorkdayStart(9 * time.Hour)
	AllBC["cz"].SetWorkdayEnd(17 * time.Hour)
	AllBC["cz"].Holidays = calCZ.Holidays
	AllBC["cz"].Name = "Czech Republic"

	AllBC["de"] = &cal.BusinessCalendar{}
	AllBC["de"].SetWorkday(time.Monday, true)
	AllBC["de"].SetWorkday(time.Tuesday, true)
	AllBC["de"].SetWorkday(time.Wednesday, true)
	AllBC["de"].SetWorkday(time.Thursday, true)
	AllBC["de"].SetWorkday(time.Friday, true)
	AllBC["de"].SetWorkdayStart(9 * time.Hour)
	AllBC["de"].SetWorkdayEnd(17 * time.Hour)
	AllBC["de"].Holidays = calDE.Holidays
	AllBC["de"].Name = "Germany"

	AllBC["dk"] = &cal.BusinessCalendar{}
	AllBC["dk"].SetWorkday(time.Monday, true)
	AllBC["dk"].SetWorkday(time.Tuesday, true)
	AllBC["dk"].SetWorkday(time.Wednesday, true)
	AllBC["dk"].SetWorkday(time.Thursday, true)
	AllBC["dk"].SetWorkday(time.Friday, true)
	AllBC["dk"].SetWorkdayStart(9 * time.Hour)
	AllBC["dk"].SetWorkdayEnd(17 * time.Hour)
	AllBC["dk"].Holidays = calDK.Holidays
	AllBC["dk"].Name = "Denmark"

	AllBC["ecb"] = &cal.BusinessCalendar{}
	AllBC["ecb"].SetWorkday(time.Monday, true)
	AllBC["ecb"].SetWorkday(time.Tuesday, true)
	AllBC["ecb"].SetWorkday(time.Wednesday, true)
	AllBC["ecb"].SetWorkday(time.Thursday, true)
	AllBC["ecb"].SetWorkday(time.Friday, true)
	AllBC["ecb"].SetWorkdayStart(9 * time.Hour)
	AllBC["ecb"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ecb"].Holidays = calECB.Holidays
	AllBC["ecb"].Name = "European Central Bank"

	AllBC["es"] = &cal.BusinessCalendar{}
	AllBC["es"].SetWorkday(time.Monday, true)
	AllBC["es"].SetWorkday(time.Tuesday, true)
	AllBC["es"].SetWorkday(time.Wednesday, true)
	AllBC["es"].SetWorkday(time.Thursday, true)
	AllBC["es"].SetWorkday(time.Friday, true)
	AllBC["es"].SetWorkdayStart(9 * time.Hour)
	AllBC["es"].SetWorkdayEnd(17 * time.Hour)
	AllBC["es"].Holidays = calES.Holidays
	AllBC["es"].Name = "Spain"

	AllBC["fr"] = &cal.BusinessCalendar{}
	AllBC["fr"].SetWorkday(time.Monday, true)
	AllBC["fr"].SetWorkday(time.Tuesday, true)
	AllBC["fr"].SetWorkday(time.Wednesday, true)
	AllBC["fr"].SetWorkday(time.Thursday, true)
	AllBC["fr"].SetWorkday(time.Friday, true)
	AllBC["fr"].SetWorkdayStart(9 * time.Hour)
	AllBC["fr"].SetWorkdayEnd(17 * time.Hour)
	AllBC["fr"].Holidays = calFR.Holidays
	AllBC["fr"].Name = "France"

	AllBC["gb"] = &cal.BusinessCalendar{}
	AllBC["gb"].SetWorkday(time.Monday, true)
	AllBC["gb"].SetWorkday(time.Tuesday, true)
	AllBC["gb"].SetWorkday(time.Wednesday, true)
	AllBC["gb"].SetWorkday(time.Thursday, true)
	AllBC["gb"].SetWorkday(time.Friday, true)
	AllBC["gb"].SetWorkdayStart(9 * time.Hour)
	AllBC["gb"].SetWorkdayEnd(17 * time.Hour)
	AllBC["gb"].Holidays = calGB.Holidays
	AllBC["gb"].Name = "United Kingdom"

	AllBC["uk"] = AllBC["gb"]

	AllBC["ie"] = &cal.BusinessCalendar{}
	AllBC["ie"].SetWorkday(time.Monday, true)
	AllBC["ie"].SetWorkday(time.Tuesday, true)
	AllBC["ie"].SetWorkday(time.Wednesday, true)
	AllBC["ie"].SetWorkday(time.Thursday, true)
	AllBC["ie"].SetWorkday(time.Friday, true)
	AllBC["ie"].SetWorkdayStart(9 * time.Hour)
	AllBC["ie"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ie"].Holidays = calIE.Holidays
	AllBC["ie"].Name = "Republic of Ireland"

	AllBC["in"] = &cal.BusinessCalendar{}
	AllBC["in"].SetWorkday(time.Monday, true)
	AllBC["in"].SetWorkday(time.Tuesday, true)
	AllBC["in"].SetWorkday(time.Wednesday, true)
	AllBC["in"].SetWorkday(time.Thursday, true)
	AllBC["in"].SetWorkday(time.Friday, true)
	AllBC["in"].SetWorkday(time.Saturday, true)
	AllBC["in"].SetWorkdayStart(9 * time.Hour)
	AllBC["in"].SetWorkdayEnd(17 * time.Hour)
	AllBC["in"].Holidays = calIN.Holidays
	AllBC["in"].Name = "India"

	AllBC["it"] = &cal.BusinessCalendar{}
	AllBC["it"].SetWorkday(time.Monday, true)
	AllBC["it"].SetWorkday(time.Tuesday, true)
	AllBC["it"].SetWorkday(time.Wednesday, true)
	AllBC["it"].SetWorkday(time.Thursday, true)
	AllBC["it"].SetWorkday(time.Friday, true)
	AllBC["it"].SetWorkdayStart(9 * time.Hour)
	AllBC["it"].SetWorkdayEnd(17 * time.Hour)
	AllBC["it"].Holidays = calIT.Holidays
	AllBC["it"].Name = "Italy"

	AllBC["jp"] = &cal.BusinessCalendar{}
	AllBC["jp"].SetWorkday(time.Monday, true)
	AllBC["jp"].SetWorkday(time.Tuesday, true)
	AllBC["jp"].SetWorkday(time.Wednesday, true)
	AllBC["jp"].SetWorkday(time.Thursday, true)
	AllBC["jp"].SetWorkday(time.Friday, true)
	AllBC["jp"].SetWorkdayStart(9 * time.Hour)
	AllBC["jp"].SetWorkdayEnd(17 * time.Hour)
	AllBC["jp"].Holidays = calJP.Holidays
	AllBC["jp"].Name = "Japan"

	AllBC["lt"] = &cal.BusinessCalendar{}
	AllBC["lt"].SetWorkday(time.Monday, true)
	AllBC["lt"].SetWorkday(time.Tuesday, true)
	AllBC["lt"].SetWorkday(time.Wednesday, true)
	AllBC["lt"].SetWorkday(time.Thursday, true)
	AllBC["lt"].SetWorkday(time.Friday, true)
	AllBC["lt"].SetWorkdayStart(9 * time.Hour)
	AllBC["lt"].SetWorkdayEnd(17 * time.Hour)
	AllBC["lt"].Holidays = calLT.Holidays
	AllBC["lt"].Name = "Lithuania"

	AllBC["mx"] = &cal.BusinessCalendar{}
	AllBC["mx"].SetWorkday(time.Monday, true)
	AllBC["mx"].SetWorkday(time.Tuesday, true)
	AllBC["mx"].SetWorkday(time.Wednesday, true)
	AllBC["mx"].SetWorkday(time.Thursday, true)
	AllBC["mx"].SetWorkday(time.Friday, true)
	AllBC["mx"].SetWorkdayStart(9 * time.Hour)
	AllBC["mx"].SetWorkdayEnd(17 * time.Hour)
	AllBC["mx"].Holidays = calMX.Holidays
	AllBC["mx"].Name = "Mexico"

	AllBC["nl"] = &cal.BusinessCalendar{}
	AllBC["nl"].SetWorkday(time.Monday, true)
	AllBC["nl"].SetWorkday(time.Tuesday, true)
	AllBC["nl"].SetWorkday(time.Wednesday, true)
	AllBC["nl"].SetWorkday(time.Thursday, true)
	AllBC["nl"].SetWorkday(time.Friday, true)
	AllBC["nl"].SetWorkdayStart(9 * time.Hour)
	AllBC["nl"].SetWorkdayEnd(17 * time.Hour)
	AllBC["nl"].Holidays = calNL.Holidays
	AllBC["nl"].Name = "Netherlands"

	AllBC["no"] = &cal.BusinessCalendar{}
	AllBC["no"].SetWorkday(time.Monday, true)
	AllBC["no"].SetWorkday(time.Tuesday, true)
	AllBC["no"].SetWorkday(time.Wednesday, true)
	AllBC["no"].SetWorkday(time.Thursday, true)
	AllBC["no"].SetWorkday(time.Friday, true)
	AllBC["no"].SetWorkdayStart(9 * time.Hour)
	AllBC["no"].SetWorkdayEnd(17 * time.Hour)
	AllBC["no"].Holidays = calNO.Holidays
	AllBC["no"].Name = "Norway"

	AllBC["nz"] = &cal.BusinessCalendar{}
	AllBC["nz"].SetWorkday(time.Monday, true)
	AllBC["nz"].SetWorkday(time.Tuesday, true)
	AllBC["nz"].SetWorkday(time.Wednesday, true)
	AllBC["nz"].SetWorkday(time.Thursday, true)
	AllBC["nz"].SetWorkday(time.Friday, true)
	AllBC["nz"].SetWorkdayStart(9 * time.Hour)
	AllBC["nz"].SetWorkdayEnd(17 * time.Hour)
	AllBC["nz"].Holidays = calNZ.Holidays
	AllBC["nz"].Name = "New Zealand"

	AllBC["pl"] = &cal.BusinessCalendar{}
	AllBC["pl"].SetWorkday(time.Monday, true)
	AllBC["pl"].SetWorkday(time.Tuesday, true)
	AllBC["pl"].SetWorkday(time.Wednesday, true)
	AllBC["pl"].SetWorkday(time.Thursday, true)
	AllBC["pl"].SetWorkday(time.Friday, true)
	AllBC["pl"].SetWorkdayStart(9 * time.Hour)
	AllBC["pl"].SetWorkdayEnd(17 * time.Hour)
	AllBC["pl"].Holidays = calPL.Holidays
	AllBC["pl"].Name = "Poland"

	AllBC["ru"] = &cal.BusinessCalendar{}
	AllBC["ru"].SetWorkday(time.Monday, true)
	AllBC["ru"].SetWorkday(time.Tuesday, true)
	AllBC["ru"].SetWorkday(time.Wednesday, true)
	AllBC["ru"].SetWorkday(time.Thursday, true)
	AllBC["ru"].SetWorkday(time.Friday, true)
	AllBC["ru"].SetWorkdayStart(9 * time.Hour)
	AllBC["ru"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ru"].Holidays = calRU.Holidays
	AllBC["ru"].Name = "Russian Federation"

	AllBC["se"] = &cal.BusinessCalendar{}
	AllBC["se"].SetWorkday(time.Monday, true)
	AllBC["se"].SetWorkday(time.Tuesday, true)
	AllBC["se"].SetWorkday(time.Wednesday, true)
	AllBC["se"].SetWorkday(time.Thursday, true)
	AllBC["se"].SetWorkday(time.Friday, true)
	AllBC["se"].SetWorkdayStart(9 * time.Hour)
	AllBC["se"].SetWorkdayEnd(17 * time.Hour)
	AllBC["se"].Holidays = calSE.Holidays
	AllBC["se"].Name = "Sweden"

	AllBC["si"] = &cal.BusinessCalendar{}
	AllBC["si"].SetWorkday(time.Monday, true)
	AllBC["si"].SetWorkday(time.Tuesday, true)
	AllBC["si"].SetWorkday(time.Wednesday, true)
	AllBC["si"].SetWorkday(time.Thursday, true)
	AllBC["si"].SetWorkday(time.Friday, true)
	AllBC["si"].SetWorkdayStart(9 * time.Hour)
	AllBC["si"].SetWorkdayEnd(17 * time.Hour)
	AllBC["si"].Holidays = calSI.Holidays
	AllBC["si"].Name = "Slovenia"

	AllBC["sk"] = &cal.BusinessCalendar{}
	AllBC["sk"].SetWorkday(time.Monday, true)
	AllBC["sk"].SetWorkday(time.Tuesday, true)
	AllBC["sk"].SetWorkday(time.Wednesday, true)
	AllBC["sk"].SetWorkday(time.Thursday, true)
	AllBC["sk"].SetWorkday(time.Friday, true)
	AllBC["sk"].SetWorkdayStart(9 * time.Hour)
	AllBC["sk"].SetWorkdayEnd(17 * time.Hour)
	AllBC["sk"].Holidays = calSK.Holidays
	AllBC["sk"].Name = "Slovakia"

	AllBC["ua"] = &cal.BusinessCalendar{}
	AllBC["ua"].SetWorkday(time.Monday, true)
	AllBC["ua"].SetWorkday(time.Tuesday, true)
	AllBC["ua"].SetWorkday(time.Wednesday, true)
	AllBC["ua"].SetWorkday(time.Thursday, true)
	AllBC["ua"].SetWorkday(time.Friday, true)
	AllBC["ua"].SetWorkdayStart(9 * time.Hour)
	AllBC["ua"].SetWorkdayEnd(17 * time.Hour)
	AllBC["ua"].Holidays = calUA.Holidays
	AllBC["ua"].Name = "Ukraine"

	AllBC["us"] = &cal.BusinessCalendar{}
	AllBC["us"].SetWorkday(time.Monday, true)
	AllBC["us"].SetWorkday(time.Tuesday, true)
	AllBC["us"].SetWorkday(time.Wednesday, true)
	AllBC["us"].SetWorkday(time.Thursday, true)
	AllBC["us"].SetWorkday(time.Friday, true)
	AllBC["us"].SetWorkdayStart(9 * time.Hour)
	AllBC["us"].SetWorkdayEnd(17 * time.Hour)
	AllBC["us"].Holidays = calUS.Holidays
	AllBC["us"].Name = "United States of America"

	AllBC["za"] = &cal.BusinessCalendar{}
	AllBC["za"].SetWorkday(time.Monday, true)
	AllBC["za"].SetWorkday(time.Tuesday, true)
	AllBC["za"].SetWorkday(time.Wednesday, true)
	AllBC["za"].SetWorkday(time.Thursday, true)
	AllBC["za"].SetWorkday(time.Friday, true)
	AllBC["za"].SetWorkdayStart(9 * time.Hour)
	AllBC["za"].SetWorkdayEnd(17 * time.Hour)
	AllBC["za"].Holidays = calZA.Holidays
	AllBC["za"].Name = "South Africa"
}

func IsWorkDayByCountry(countryCode string, dt time.Time) (isWork bool, err error) {
	countryCode = strings.ToLower(countryCode)
	if len(AllBC) != 0 {
		if bdays, ok := AllBC[countryCode]; ok {
			isWork = bdays.IsWorkday(dt)
		} else {
			err = ErrCountryNotAdded
		}
	} else {
		err = ErrAllBDNotInit
	}
	return isWork, err
}
