package pkg

import "time"

func StringTotime(str string) time.Time {
	loc,_ := time.LoadLocation("Local")
	time,_ := time.ParseInLocation("2006-01-02", str, loc)
	return time
}
