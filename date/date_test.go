package date

import (
	"testing"
	"time"
)

func TestDatesFromString(t *testing.T) {
	testValidDates := map[string]string{
		"2014-02-01": "2014-02-01",
		"2014-02-31": "2014-03-03",
		"2014-03-01": "2014-03-01",
	}
	testInvalidDates := []string{
		"0000-00-00",
		"2014-15-01",
		"2015-00-00",
		"2015-12-35",
	}
	for ip, op := range testValidDates {
		d, err := GetDateFromString(ip)
		if err != nil {
			t.Fatalf("Expected no error. Recieved: [%s]", err)
			continue
		}
		outputDate := d.String()
		if op != outputDate {
			t.Fatalf("Expected [%s]. Received[%s]", op, outputDate)
		}
	}

	for _, dateStr := range testInvalidDates {
		_, err := GetDateFromString(dateStr)
		if err == nil {
			t.Fatalf("[%s] is not a valid date. Expected an error. Instead received nil", dateStr)
		}
	}
}

func TestDatesFromUnixTimestamp(t *testing.T) {
	testCases := map[int64]string{
		1429685920: "2015-04-22", // Wed Apr 22 06:58:40 UTC 2015
	}
	for ut, dateStr := range testCases {
		d := GetDateFromUnixTimestamp(ut)
		outputDate := d.String()
		if dateStr != outputDate {
			t.Fatalf("Expected [%s]. Received[%s]", dateStr, outputDate)
		}
	}
}

func BenchmarkUnixDate(b *testing.B) {
	ut := time.Now().Unix()

	hrs10 := int64(3600 * 10)
	for i := 0; i < b.N; i++ {
		d1 := GetDateFromUnixTimestamp(ut + hrs10)
		d1.String()
		d2 := GetDateFromUnixTimestamp(ut - hrs10)
		d2.String()
	}
}