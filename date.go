package stdlib

import (
	"time"

	"github.com/araddon/dateparse"
)

func dateYear(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.Year())
}

func dateMonth(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.Month())
}

func dateDay(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.Day())
}


func dateYearDay(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.YearDay())
}

func dateHour(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.Hour())
}

func dateMinute(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.Minute())
}

func dateSecond(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return int64(t.Second())
}

func dateUnix(s string) int64 {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return -1
	}

	return t.Unix()
}

func dateRfc3339(s string) string {
	t, err := dateparse.ParseAny(s)
	if err != nil {
		return ""
	}

	return t.Format(time.RFC3339)
}

var dateFunctions = map[string]any{
	"date_year": stringy1int64(dateYear),
	"date_month": stringy1int64(dateMonth),
	"date_day": stringy1int64(dateDay),
	"date_yearday": stringy1int64(dateYearDay),
	"date_hour": stringy1int64(dateHour),
	"date_minute": stringy1int64(dateMinute),
	"date_second": stringy1int64(dateSecond),
	"date_unix": stringy1int64(dateUnix),
	"date_rfc3339": stringy1string(dateRfc3339),
}
