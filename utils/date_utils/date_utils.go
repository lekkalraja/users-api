package date_utils

import "time"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format("2006:01:02T15:04:05")
}
