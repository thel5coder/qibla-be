package time

import (
	"time"
)

// Parse ...
func Parse(data, format, location string) (res time.Time, err error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return res, err
	}

	res, err = time.ParseInLocation(format, data, loc)
	if err != nil {
		return res, err
	}

	return res, nil
}

// In ...
func In(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

// InFormat ...
func InFormat(t time.Time, name, format string) (string, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return "", err
	}

	return t.In(loc).Format(format), err
}

// InFormatNoErr ...
func InFormatNoErr(t time.Time, name, format string) string {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return ""
	}

	return t.In(loc).Format(format)
}

// Convert ...
func Convert(t, fromFormat, toFormat string) string {
	timeConvert, err := time.Parse(fromFormat, t)
	if err != nil {
		return ""
	}

	return timeConvert.Format(toFormat)
}

// ConvertLocation ...
func ConvertLocation(t, fromFormat, toFormat, location string) string {
	timeConvert, err := Parse(t, fromFormat, location)
	if err != nil {
		return ""
	}

	return timeConvert.Format(toFormat)
}

// Diff ...
func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

// DiffCustom ...
func DiffCustom(start string, b time.Time) (year, month, day, hour, min, sec int) {
	if start == "" {
		return
	}
	a, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return
	}

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

// AddTimezone ...
func AddTimezone(date, timezone string) string {
	if date == "" {
		return ""
	}

	return date + "T00:00:00" + timezone
}

// AddDays ...
func AddDays(date string, days int) string {
	newDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}

	res := newDate.AddDate(0, 0, days)

	return res.Format("2006-01-02")
}

// CheckDate ...
func CheckDate(data, format string) string {
	_, err := time.Parse(format, data)
	if err != nil {
		return ""
	}

	return data
}

// IsValid ...
func IsValid(data, format string) bool {
	_, err := time.Parse(format, data)
	if err != nil {
		return false
	}

	return true
}

// IsValidAfterToday ...
func IsValidAfterToday(data, format, location string) bool {
	date, err := Parse(data, format, location)
	if err != nil {
		return false
	}

	if date.Before(time.Now()) {
		return false
	}

	return true
}

// IsValidBeforeToday ...
func IsValidBeforeToday(data, format, location string) bool {
	date, err := Parse(data, format, location)
	if err != nil {
		return false
	}

	if date.After(time.Now()) {
		return false
	}

	return true
}

// IsValidRange ...
func IsValidRange(start, end, format string) bool {
	startDate, err := time.Parse(format, start)
	if err != nil {
		return false
	}
	endDate, err := time.Parse(format, end)
	if err != nil {
		return false
	}

	if startDate.After(endDate) {
		return false
	}

	return true
}

// IsValidRangeBeforeToday ...
func IsValidRangeBeforeToday(start, end, format, location string) bool {
	startDate, err := Parse(start, format, location)
	if err != nil {
		return false
	}
	endDate, err := Parse(end, format, location)
	if err != nil {
		return false
	}

	if startDate.After(endDate) {
		return false
	}
	if endDate.After(time.Now()) {
		return false
	}

	return true
}

// DateRangeYearMonth ...
func DateRangeYearMonth(year, month, timezone string) (startDate, endDate time.Time, err error) {
	start := year + "-" + month + "-01T00:00:00" + timezone
	startDate, err = time.Parse(time.RFC3339, start)
	if err != nil {
		return startDate, endDate, err
	}

	endDate = startDate.AddDate(0, 1, 0)

	return startDate, endDate, err
}
