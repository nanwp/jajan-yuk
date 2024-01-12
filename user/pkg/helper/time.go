package helper

import "time"

func GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	return time.Now()
}
