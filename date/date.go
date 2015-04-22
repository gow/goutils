package date

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	date_stamp_format = "2006-01-02"
)

// DATE
type Date struct {
	time.Time
}

// DateCmp compares d1 and d2. It returns 0 if both dates are equal, -1 if d1>d2 and 1 if d2>d1
func DateCmp(d1 Date, d2 Date) int {
	d1Str := d1.String()
	d2Str := d2.String()
	if d1Str == d2Str {
		return 0
	}
	if d2Str > d1Str {
		return 1
	}
	return -1
}

// DateMin returns the earlier date from the input dates
// TODO: make it work on ...
func DateMin(d1 Date, d2 Date) Date {
	if d1.String() > d2.String() {
		return d2
	}
	return d1
}

// DateMax returns the later date from the input dates.
// TODO: make it work on ...
func DateMax(d1 Date, d2 Date) Date {
	if d1.String() > d2.String() {
		return d1
	}
	return d2
}

// DateEqual returns true if d1 & d2 represent the same date in UTC.
func DateEqual(d1 Date, d2 Date) bool {
	return d1.String() == d2.String()
}

func GetDateFromString(dateStr string) (d Date, err error) {
	t, err := time.Parse(date_stamp_format, dateStr)
	if err != nil {
		return
	}
	d = Date{t}
	return
}

func (d Date) String() string {
	return d.Format(date_stamp_format)
}

func (d Date) GetBSON() (val interface{}, err error) {
	val = d.String()
	return
}

func (d *Date) SetBSON(raw bson.Raw) (err error) {
	var str string
	err = raw.Unmarshal(&str)
	if err != nil {
		return
	}
	t, err := getTimeFromDateString(str)
	if err != nil {
		return
	}
	*d = Date{t}
	return
}

func (d Date) T() time.Time {
	return d.Time
}

func (d Date) AddDays(i int) Date {
	return Date{d.AddDate(0, 0, i)}
}

func getTimeFromDateString(date string) (time.Time, error) {
	return time.Parse(date_stamp_format, date)
}

func GetDateFromUnixTimestamp(ts int64) Date {
	ds := ts - (ts % 86400)
	return Date{time.Unix(ds, 0).In(time.UTC)}
}
