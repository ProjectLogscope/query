package timerange

import "time"

var (
	MinTimestampValue time.Time = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	MaxTimestampValue time.Time = time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)
)
