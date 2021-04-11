package dates

import "time"

const (
	apiDateFormat = "Mon Jan 2 15:04:05 -0700 MST 2006"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}
