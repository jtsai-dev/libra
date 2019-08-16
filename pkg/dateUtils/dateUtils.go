package dateUtils

import (
	"time"
)

const defaultFormater = "2006-01-02 15:04:05"

// DefaultFormat return the time in "yyyy-MM-dd HH:mm:ss"
func DefaultFormat(time time.Time) string {
	return time.Format(defaultFormater)
}

func Format(time time.Time, formatter string) string {
	return time.Format(formatter)
}
