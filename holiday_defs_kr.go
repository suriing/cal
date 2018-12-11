package cal

import "time"

// Korea holidays
var (
	// International holidays
	KRSolarNewYear = NewYear
	KRChristmasDay = Christmas
	// Korea fixed holidays
	KRMarchFirstDay         = NewHoliday(time.March, 1)   // 삼일절
	KRChildrensDay          = NewHoliday(time.May, 5)     // 어린이날 *대체공휴일
	KRMemorialDay           = NewHoliday(time.June, 6)    // 현충일
	KRIndependenceDay       = NewHoliday(time.August, 15) // 광복절
	KRNationalFoundationDay = NewHoliday(time.October, 3) // 개천절
	KRHangulDay             = NewHoliday(time.October, 9) // 한글날
	// Korea lunar based holidays
	KRLunarNewYears      = NewHolidayFunc(calculateLunarNewYearsHoliday)      // 설날 (음 1/1)	*대체공휴일
	KRLunarNewYears1     = NewHolidayFunc(calculateLunarNewYearsHoliday1)     // 설 연휴 첫날
	KRLunarNewYears2     = NewHolidayFunc(calculateLunarNewYearsHoliday2)     // 설 연휴 마지막날
	KRBuddhasBirthday    = NewHolidayFunc(calculateBuddhasBirthdayHoliday)    // 부처님 오신 날 (음 4/8)
	KRMidautumnFestival  = NewHolidayFunc(calculateMidautumnFestivalHoliday)  // 추석 (음 8/15)	*대체공휴일
	KRMidautumnFestival1 = NewHolidayFunc(calculateMidautumnFestivalHoliday1) // 추석 연휴 첫날
	KRMidautumnFestival2 = NewHolidayFunc(calculateMidautumnFestivalHoliday2) // 추석 연휴 마지막날
)

// AddKoreaHolidays adds all Korea holidays to the Calender
func AddKoreaHolidays(c *Calendar) {
	c.AddHoliday(
		KRSolarNewYear,
		KRChristmasDay,
		KRMarchFirstDay,
		KRChildrensDay,
		KRMemorialDay,
		KRIndependenceDay,
		KRNationalFoundationDay,
		KRHangulDay,
		KRLunarNewYears,
		KRLunarNewYears1,
		KRLunarNewYears2,
		KRBuddhasBirthday,
		KRMidautumnFestival,
		KRMidautumnFestival1,
		KRMidautumnFestival2,
	)
	// LunarNewYearsDay, Children's day and Midautumn Festival applied to alternative holiday in case the day is Sunday
	if KRChildrensDay.Weekday.String() == "Sunday" {
		c.AddHoliday(NewHoliday(time.May, 6))
	}
	if KRLunarNewYears.Weekday.String() == "Saturday" || KRLunarNewYears.Weekday.String() == "Sunday" || KRLunarNewYears.Weekday.String() == "Monday" {
		c.AddHoliday(NewHolidayFunc(calculateLunarNewYearsHoliday3))
	}
	if KRMidautumnFestival.Weekday.String() == "Saturday" || KRMidautumnFestival.Weekday.String() == "Sunday" || KRMidautumnFestival.Weekday.String() == "Monday" {
		c.AddHoliday(NewHolidayFunc(calculateMidautumnFestivalHoliday3))
	}
}

// Korea lunar based holiday functions
func calculateLunarNewYears(year int, loc *time.Location) time.Time {
	LunarNewYears := NewLunar(year, 1, 1, 0, 0, 0, loc)
	return LunarNewYears.Convert()
}

func calculateLunarNewYearsHoliday(year int, loc *time.Location) (time.Month, int) {
	LunarNewYears := calculateLunarNewYears(year, loc)
	return LunarNewYears.Month(), LunarNewYears.Day()
}

func calculateLunarNewYearsHoliday1(year int, loc *time.Location) (time.Month, int) {
	LunarNewYears := calculateLunarNewYears(year, loc)
	LunarNewYears1 := LunarNewYears.AddDate(0, 0, -1)
	return LunarNewYears1.Month(), LunarNewYears1.Day()
}

func calculateLunarNewYearsHoliday2(year int, loc *time.Location) (time.Month, int) {
	LunarNewYears := calculateLunarNewYears(year, loc)
	LunarNewYears2 := LunarNewYears.AddDate(0, 0, 1)
	return LunarNewYears2.Month(), LunarNewYears2.Day()
}

func calculateLunarNewYearsHoliday3(year int, loc *time.Location) (time.Month, int) { // 대체공휴일
	LunarNewYears := calculateLunarNewYears(year, loc)
	LunarNewYears3 := LunarNewYears.AddDate(0, 0, 2)
	return LunarNewYears3.Month(), LunarNewYears3.Day()
}

func calculateBuddhasBirthdayHoliday(year int, loc *time.Location) (time.Month, int) {
	BuddhasBirthday := NewLunar(year, 4, 8, 0, 0, 0, loc)
	day := BuddhasBirthday.Convert()
	return day.Month(), day.Day()
}

func calculateMidautumnFestival(year int, loc *time.Location) time.Time {
	LunarMidautumnFestival := NewLunar(year, 8, 15, 0, 0, 0, loc)
	return LunarMidautumnFestival.Convert()
}

func calculateMidautumnFestivalHoliday(year int, loc *time.Location) (time.Month, int) {
	MidautumnFestival := calculateMidautumnFestival(year, loc)
	return MidautumnFestival.Month(), MidautumnFestival.Day()
}

func calculateMidautumnFestivalHoliday1(year int, loc *time.Location) (time.Month, int) {
	MidautumnFestival := calculateMidautumnFestival(year, loc)
	MidautumnFestival1 := MidautumnFestival.AddDate(0, 0, -1)
	return MidautumnFestival1.Month(), MidautumnFestival1.Day()
}

func calculateMidautumnFestivalHoliday2(year int, loc *time.Location) (time.Month, int) {
	MidautumnFestival := calculateMidautumnFestival(year, loc)
	MidautumnFestival2 := MidautumnFestival.AddDate(0, 0, 1)
	return MidautumnFestival2.Month(), MidautumnFestival2.Day()
}

func calculateMidautumnFestivalHoliday3(year int, loc *time.Location) (time.Month, int) { // 대체공휴일
	MidautumnFestival := calculateMidautumnFestival(year, loc)
	MidautumnFestival3 := MidautumnFestival.AddDate(0, 0, 2)
	return MidautumnFestival3.Month(), MidautumnFestival3.Day()
}

// Handling lunar dates
// came from "github.com/xishvai/lunar" and modified by suriing
// Valid year range : 1900~2050
const (
	LunarMinYear = 1900
	LunarMaxYear = 2050
)

var (
	lunarTable = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260,
		0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2,
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255,
		0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977,
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40,
		0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970,
		0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0,
		0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950,
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4,
		0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557,
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0,
		0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0,
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570,
		0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0,
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4,
		0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6,
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a,
		0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570,
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50,
		0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0,
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552,
		0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5,
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9,
		0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930,
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60,
		0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530,
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0,
		0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45,
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577,
		0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0,
	}
	LunarBase = time.Date(LunarMinYear, 1, 31, 0, 0, 0, 0, time.Local)
)

type Lunar struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
	loc    *time.Location
}

func (l *Lunar) Year() int                { return l.year }
func (l *Lunar) Month() int               { return l.month }
func (l *Lunar) Day() int                 { return l.day }
func (l *Lunar) Hour() int                { return l.hour }
func (l *Lunar) Minute() int              { return l.minute }
func (l *Lunar) Second() int              { return l.second }
func (l *Lunar) Location() *time.Location { return l.loc }

func NewLunar(year, month, day, hour, min, sec int, loc *time.Location) *Lunar {
	if !isYearValid(year) {
		return nil
	}
	return &Lunar{year, month, day, hour, min, sec, loc}
}

func (l *Lunar) Convert() time.Time {
	lyear := l.Year()
	lmonth := l.Month()
	lday := l.Day()
	offset := 0
	leap := IsLeap(lyear)

	// increment year
	for i := LunarMinYear; i < lyear; i++ {
		offset += YearDays(i)
	}

	// increment month
	// add days in all months up to the current month
	var cur int
	for cur = 1; cur < lmonth; cur++ {
		// add extra days for leap month
		if cur == LeapMonth(lyear) {
			offset += LeapDays(lyear)
		}
		offset += MonthDays(lyear, cur)
	}
	// if current month is leap month, add days in normal month
	isLeapMonth := (LeapMonth(lyear) == lmonth)

	if leap && isLeapMonth {
		offset += MonthDays(lyear, cur)
	}
	// increment
	offset += (lday - 1)

	//BUG: maybe overflow
	d := time.Duration(offset*24) * time.Hour
	solar := LunarBase.Add(d)

	year := solar.Year()
	month := int(solar.Month())
	day := solar.Day()
	return time.Date(year, time.Month(month), day, l.hour, l.Minute(), l.Second(), 0, l.Location())
}

func IsLeap(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

//the total days of this year
func YearDays(year int) int {
	sum := 348
	for i := 0x8000; i > 0x8; i >>= 1 {
		if (lunarTable[year-LunarMinYear] & i) != 0 {
			sum += 1
		}
	}
	return sum + LeapDays(year)
}

//which month leaps in this year?
//return 1-12(if there is one) or 0(no leap month).
func LeapMonth(year int) int {
	return int(lunarTable[year-LunarMinYear] & 0xf)
}

//the days of this year's leap month
func LeapDays(year int) int {
	if LeapMonth(year) != 0 {
		if (lunarTable[year-LunarMinYear] & 0x10000) != 0 {
			return 30
		}
		return 29
	}
	return 0
}

//the days of the m-th month of this year
func MonthDays(year, month int) int {
	if (lunarTable[year-LunarMinYear] & (0x10000 >> uint(month))) != 0 {
		return 30
	}
	return 29
}

func isYearValid(year int) bool {
	if year > LunarMaxYear || year < LunarMinYear {
		return false
	}
	return true
}
