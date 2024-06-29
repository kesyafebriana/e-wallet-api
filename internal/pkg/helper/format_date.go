package helper

import "time"

func Timezone(date time.Time) *time.Time {
	formattedExpired := time.Date(date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		time.Local)

	return &formattedExpired
}
