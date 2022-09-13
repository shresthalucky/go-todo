package util

import "time"

func GetCurrentUTCTime() *time.Time {
	t := time.Now().UTC()
	return &t
}
