package date

import (
	"fmt"
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

/*
type Dates []Date

func (d Dates) Len() int {
	return len(d)
}
func (d Dates) Swap(i, j int) {
	d[i], d[j] = d[i], d[j]
}
func (d Dates) Equal(i, j int) bool {
}
*/

// Cmp compares d1 and d2. It returns 0 if both dates are equal, -1 if d1>d2 and 1 if d2>d1
func Cmp(d1 Date, d2 Date) int {
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

// Min returns the earlier date from the input dates
func Min(d1 Date, dates ...Date) Date {
	min := d1
	for _, d := range dates {
		if min.After(d.Time) {
			min = d
		}
	}
	return min
}

// Max returns the later date from the input dates.
func Max(d1 Date, dates ...Date) Date {
	max := d1
	for _, d := range dates {
		if max.Before(d.Time) {
			max = d
		}
	}
	return max
}

// Equal returns true if d1 & d2 represent the same date in UTC.
func Equal(d1 Date, d2 Date) bool {
	return d1 == d2
}

func (d Date) String() string {
	yy, mm, dd := d.Date()
	return fmt.Sprintf("%04d-%02d-%02d", yy, mm, dd)
}

func (d Date) GetBSON() (interface{}, error) {
	return d.String(), nil
}

func (d *Date) SetBSON(raw bson.Raw) error {
	var str string
	err := raw.Unmarshal(&str)
	if err != nil {
		return err
	}
	t, err := getTimeFromDateString(str)
	if err != nil {
		return err
	}
	*d = Date{t}
	return nil
}

/*
func (d Date) T() time.Time {
	return d.Time
}
*/

func (d Date) AddDays(i int) Date {
	return Date{d.AddDate(0, 0, i)}
}

func getTimeFromDateString(date string) (time.Time, error) {
	t, err := time.Parse(date_stamp_format, date)
	if err != nil {
		return t, err
	}
	return t.In(time.UTC), err
}

func NewFromUnixTimestamp(ts int64) Date {
	ds := ts - (ts % 86400)
	return Date{time.Unix(ds, 0).In(time.UTC)}
}

func NewFromString(dateStr string) (d Date, err error) {
	t, err := getTimeFromDateString(dateStr)
	if err != nil {
		return
	}
	d = Date{t}
	return
}

func NewFromTime(t time.Time) (Date, error) {
	if t.IsZero() {
		return Date{}, fmt.Errorf("Invalid time: %#v", t)
	}
	yy, mm, dd := t.Date()
	newTime := time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC)
	return Date{newTime}, nil
}
