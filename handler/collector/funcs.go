package collector

import (
	_ "embed"
	"time"
)

func formatDate(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	return date.Format("02/01/2006")
}

func formatHour(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	return date.Format("15:04")
}

func formatHourOnly(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	return date.Format("15")
}

func formatTime(date time.Time) string {
	return date.Format(time.RFC3339)
}

func currentHour(timezone string) (string, error) {
    loc, err := time.LoadLocation(timezone)
    if err != nil {
        return "", err
    }
    currentTime := time.Now().In(loc)
    return currentTime.Format("15"), nil
}