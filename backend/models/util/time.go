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

func FixQueryPeriodDates(startDate, endDate string) (fixedStartDate, fixedEndDate string) {
	defaultDateSpan := 30

	loc, _ := time.LoadLocation("Asia/Shanghai")
	et, err := time.ParseInLocation(DATE_FORMAT, endDate, loc)
	if err != nil {
		et = time.Now().In(loc)
	}

	st, err := time.ParseInLocation(DATE_FORMAT, startDate, loc)
	if err != nil {
		st = et.AddDate(0, 0, -defaultDateSpan)
	}

	fixedStartDate = st.Format(DATE_FORMAT)
	fixedEndDate = et.Format(DATE_FORMAT)
	return
}
