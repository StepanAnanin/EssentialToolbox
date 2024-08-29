package essential

import "time"

// Timestamp from this moment + t in UNIX time format
func TimestampSinceNow(t time.Duration) int64 {
	return time.Now().Add(t).UTC().UnixMilli()
}

// Current timestamp in UNIX time format
func TimestampNow() int64 {
	return time.Now().UTC().UnixMilli()
}
