package dara

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	date time.Time
}

func NewDate(dateInput string) (*Date, error) {
	var t time.Time
	var err error
	// 解析输入时间，如果输入格式不对，返回错误
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05Z",
	}

	for _, format := range formats {
		t, err = time.Parse(format, dateInput)
		if err == nil {
			return &Date{date: t}, nil
		}
	}

	return nil, fmt.Errorf("unable to parse date: %v", dateInput)
}

func (t *Date) Format(layout string) string {

	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	// layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "hh", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "a", "PM", 1)
	layout = strings.Replace(layout, "EEEE", "Monday", 1)
	layout = strings.Replace(layout, "E", "Mon", 1)
	return t.date.Format(layout)
}

func (t *Date) Unix() int64 {
	return t.date.Unix()
}

func (t *Date) UTC() string {
	return t.date.UTC().Format("2006-01-02 15:04:05.000000000 -0700 MST")
}

func (t *Date) Sub(amount int, unit string) *Date {
	var duration time.Duration
	switch unit {
	case "second", "seconds":
		duration = time.Duration(-amount) * time.Second
	case "minute", "minutes":
		duration = time.Duration(-amount) * time.Minute
	case "hour", "hours":
		duration = time.Duration(-amount) * time.Hour
	case "day", "days":
		duration = time.Duration(-amount) * 24 * time.Hour
	case "week", "weeks":
		duration = time.Duration(-amount) * 7 * 24 * time.Hour
	case "month", "months":
		return &Date{date: t.date.AddDate(0, -amount, 0)}
	case "year", "years":
		return &Date{date: t.date.AddDate(-amount, 0, 0)}
	default:
		return nil
	}
	newDate := t.date.Add(duration)
	return &Date{date: newDate}
}

func (t *Date) Add(amount int, unit string) *Date {
	var duration time.Duration
	switch unit {
	case "second", "seconds":
		duration = time.Duration(amount) * time.Second
	case "minute", "minutes":
		duration = time.Duration(amount) * time.Minute
	case "hour", "hours":
		duration = time.Duration(amount) * time.Hour
	case "day", "days":
		duration = time.Duration(amount) * 24 * time.Hour
	case "week", "weeks":
		duration = time.Duration(amount) * 7 * 24 * time.Hour
	case "month", "months":
		return &Date{date: t.date.AddDate(0, amount, 0)}
	case "year", "years":
		return &Date{date: t.date.AddDate(amount, 0, 0)}
	default:
		return nil
	}

	newDate := t.date.Add(duration)
	return &Date{date: newDate}
}

func (t *Date) Diff(amount string, diffDate *Date) int64 {
	switch amount {
	case "second", "seconds":
		return int64(t.date.Sub(diffDate.date).Seconds())
	case "minute", "minutes":
		return int64(t.date.Sub(diffDate.date).Minutes())
	case "hour", "hours":
		return int64(t.date.Sub(diffDate.date).Hours())
	case "day", "days":
		return int64(t.date.Sub(diffDate.date).Hours() / 24)
	case "week", "weeks":
		return int64(t.date.Sub(diffDate.date).Hours() / (24 * 7))
	case "month", "months":
		return int64(diffDate.date.Year()*12 + int(diffDate.date.Month()) - (t.date.Year()*12 + int(t.date.Month())))
	case "year", "years":
		return int64(t.date.Year() - diffDate.date.Year())
	default:
		return 0
	}
}

func (t *Date) Hour() int {
	return t.date.Hour()
}

func (t *Date) Minute() int {
	return t.date.Minute()
}

func (t *Date) Second() int {
	return t.date.Second()
}

func (t *Date) Month() int {
	return int(t.date.Month())
}

func (t *Date) Year() int {
	return t.date.Year()
}

func (t *Date) DayOfMonth() int {
	return t.date.Day()
}

func (t *Date) DayOfWeek() int {
	weekday := int(t.date.Weekday())
	if weekday == 0 {
		return 7 // Sunday
	}
	return weekday
}

func (t *Date) WeekOfYear() int {
	_, week := t.date.ISOWeek()
	return week
}
