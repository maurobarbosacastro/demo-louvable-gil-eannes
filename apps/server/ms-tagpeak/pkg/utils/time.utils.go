package utils

import "time"

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func GetTimeRanges() map[string]TimeRange {
	now := time.Now()
	return map[string]TimeRange{
		"currentMonth": {
			Start: time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
			End:   time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Add(-time.Second),
		},
		"lastMonth": {
			Start: time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location()),
			End:   time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Add(-time.Second),
		},
		"last12Months": {
			Start: time.Date(now.Year(), now.Month()-11, 1, 0, 0, 0, 0, now.Location()),
			End:   now,
		},
	}
}
