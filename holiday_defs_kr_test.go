package cal

import (
	"testing"
	"time"
)

func TestAddKRHolidays(t *testing.T) {
	// The following is all of the national holidays in Korea for the year 2018
	type date struct {
		day   int
		month time.Month
	}
	holidays := map[string]date{
		"solar_new_year": {
			day:   1,
			month: time.January,
		},
		"lunar_new_year1": {
			day:   15,
			month: time.February,
		},
		"lunar_new_year": {
			day:   16,
			month: time.February,
		},
		"lunar_new_year2": {
			day:   17,
			month: time.February,
		},
		"march_first": {
			day:   1,
			month: time.March,
		},
		"childrens": {
			day:   5,
			month: time.May,
		},
		"memorial": {
			day:   6,
			month: time.June,
		},
		"independence": {
			day:   15,
			month: time.August,
		},
		"midautumn_festival1": {
			day:   23,
			month: time.September,
		},
		"midautumn_festival": {
			day:   24,
			month: time.September,
		},
		"midautumn_festival2": {
			day:   25,
			month: time.September,
		},
		"midautumn_festival3": { // 대체공휴일
			day:   26,
			month: time.September,
		},
		"national_foundation": {
			day:   3,
			month: time.October,
		},
		"hangul": {
			day:   9,
			month: time.October,
		},
		"christmas": {
			day:   25,
			month: time.December,
		},
	}

	for name, holiday := range holidays {
		t.Run(name, func(t *testing.T) {
			c := NewCalendar()
			AddKoreaHolidays(c)
			i := time.Date(2018, holiday.month, holiday.day, 0, 0, 0, 0, time.Local)

			if !c.IsHoliday(i) {
				t.Errorf("Expected %q to be a holiday but wasn't", i)
			}
			if c.IsWorkday(i) {
				t.Errorf("Did not expect %q to be a holiday", i)
			}
		})
	}
}
