package helper

import "time"

func ToTimeFormatter(data string) time.Time {
	format := "15:04:05"
	time, _ := time.Parse(format, data)
	return time
}

func ToDateFormatter(data string) time.Time {
	format := "2006-01-02"
	date, _ := time.Parse(format, data)
	return date
}

func DateTOString(date time.Time) string {
	return date.Format("15:04:05")
}
