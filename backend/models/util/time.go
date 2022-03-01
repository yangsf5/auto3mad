package util

import (
	"time"
)

const (
	DATE_FORMAT = "2006-01-02"
)

func GetDateTimestamp(date string) (firstSecond, lastSecond int64) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(DATE_FORMAT, date, loc)
	firstSecond = t.Unix()
	lastSecond = t.AddDate(0, 0, 1).Unix() - 1
	return
}

func GetPeriodDates(startDate, endDate string) (dates []string) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	st, _ := time.ParseInLocation(DATE_FORMAT, startDate, loc)
	et, _ := time.ParseInLocation(DATE_FORMAT, endDate, loc)

	if st.Unix() <= et.Unix() {
		for t := st; t.Unix() <= et.Unix(); {
			dates = append(dates, t.Format(DATE_FORMAT))
			t = t.AddDate(0, 0, 1)
		}
	} else {
		for t := st; t.Unix() >= et.Unix(); {
			dates = append(dates, t.Format(DATE_FORMAT))
			t = t.AddDate(0, 0, -1)
		}
	}
	return
}

func GetWeekTimestamp(date string) (firstSecond, lastSecond int64) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(DATE_FORMAT, date, loc)
	wd := int64(t.Weekday())
	if wd == 0 {
		wd = 7
	}

	firstSecond, lastSecond = GetDateTimestamp(date)
	firstSecond = firstSecond - (wd-1)*24*60*60
	lastSecond = lastSecond + (7-wd)*24*60*60
	return
}
