package convert

import (
	"time"
)

func ParseDateString(layout string, dateStr string) (time.Time, error) {

	// layout := "2006-01-02"

	dateTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}
