package utility

import "time"

func GetEpochTime() int64 {
	now := time.Now()
	return now.Unix()
}
