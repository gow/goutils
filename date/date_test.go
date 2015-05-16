package date

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestDatesFromString(t *testing.T) {
	testValidDates := map[string]string{
		"2014-02-01": "2014-02-01",
		"2014-02-31": "2014-03-03",
		"2014-03-01": "2014-03-01",
		"2020-03-02": "2020-03-02",
	}
	testInvalidDates := []string{
		"0000-00-00",
		"2014-15-01",
		"2015-00-00",
		"2015-12-35",
	}
	for ip, op := range testValidDates {
		d, err := NewFromString(ip)
		if err != nil {
			t.Fatalf("Expected no error. Recieved: [%s]", err)
			continue
		}
		outputDate := d.String()
		if !reflect.DeepEqual(op, outputDate) {
			t.Fatalf("Expected [%#v]. Received[%#v]", op, outputDate)
		}
	}

	for _, dateStr := range testInvalidDates {
		_, err := NewFromString(dateStr)
		if err == nil {
			t.Fatalf("[%s] is not a valid date. Expected an error. Instead received nil", dateStr)
		}
	}
}

func TestDatesFromUnixTimestamp(t *testing.T) {
	testCases := map[int64]string{
		1429685920: "2015-04-22", // Wed Apr 22 06:58:40 UTC 2015
		972000000:  "2000-10-20", // Fri Oct 20 00:00:00 UTC 2000
		2221257039: "2040-05-21", // Mon May 21 23:50:39 UTC 2040
	}
	for ut, dateStr := range testCases {
		d := NewFromUnixTimestamp(ut)
		outputDate := d.String()
		if dateStr != outputDate {
			t.Fatalf("Expected [%s]. Received[%s]", dateStr, outputDate)
		}
	}
}

func TestDateFromTimeStruct(t *testing.T) {
	// error case
	_, err := NewFromTime(time.Time{})
	if err == nil {
		t.Fatal("Expected an error while getting date from empty time struct. Instead received nil")
	}

	// Success case
	expectedDateStr := "2015-02-28"
	d, err := NewFromTime(time.Unix(1425081600, 0)) // "2015-02-28"
	if err != nil {
		t.Fatalf("Expected no error. Recieved: [%s]", err)
	}
	dateStr := d.String()
	if dateStr != expectedDateStr {
		t.Fatalf("Expected [%s]. Instead received [%s]", expectedDateStr, dateStr)
	}
}

func TestDatesFromDifferentFormats(t *testing.T) {
	testCases := map[int64]string{
		1429685920: "2015-04-22", // Wed Apr 22 06:58:40 UTC 2015
		972000000:  "2000-10-20", // Fri Oct 20 00:00:00 UTC 2000
		2221257039: "2040-05-21", // Mon May 21 23:50:39 UTC 2040
	}
	for ut, dateStr := range testCases {
		dateFromUnix := NewFromUnixTimestamp(ut)
		dateFromString, err := NewFromString(dateStr)
		if err != nil {
			t.Fatalf("Expected no error. Recieved: [%s]", err)
			continue
		}

		if !reflect.DeepEqual(dateFromUnix, dateFromString) {
			t.Fatalf("Expected the following dates to be deeply equal:\n[%#v]\n[%#v]", dateFromUnix, dateFromString)
		}

		dateFromTime, err := NewFromTime(time.Unix(ut, 0))
		if err != nil {
			t.Fatalf("Expected no error. Recieved: [%s]", err)
			continue
		}
		if !reflect.DeepEqual(dateFromTime, dateFromString) {
			t.Fatalf("Expected the following dates to be deeply equal:\n[%#v]\n[%#v]", dateFromTime, dateFromString)
		}
	}
}

func TestDateComparison(t *testing.T) {
	d1, err := NewFromString("2015-02-28")
	if err != nil {
		t.Fatalf("Expected no error. Recieved: [%s]", err)
	}
	if Equal(d1, d1) != true {
		t.Fatalf("Expected true. Instead received false. Date: [%v]", d1)
	}
}

func TestDateBSONConversions(t *testing.T) {
	dateStr := "2015-02-28"
	d, err := NewFromString(dateStr)
	if err != nil {
		t.Fatalf("Expected no error. Recieved: [%s]", err)
	}

	bsonVal, err := d.GetBSON()
	if err != nil {
		t.Fatalf("Expected no error. Recieved: [%s]", err)
	}

	if bsonVal != dateStr {
		t.Fatalf("Expected [%s]. Received [%s]", dateStr, bsonVal)
	}
}

// TestDateMinMax tests the Min & Max functions
func TestDateMinMax(t *testing.T) {
	cases := []struct {
		Input       []Date
		ExpectedMin string
		ExpectedMax string
	}{
		{
			[]Date{
				Date{time.Unix(932774400, 0)}, // Sat Jul 24 00:00:00 UTC 1999
			},
			"1999-07-24",
			"1999-07-24",
		},
		{
			[]Date{
				Date{time.Unix(1429685920, 0)}, // Wed Apr 22 06:58:40 UTC 2015
				Date{time.Unix(972000000, 0)},  // Fri Oct 20 00:00:00 UTC 2000
				Date{time.Unix(2221257039, 0)}, // Mon May 21 23:50:39 UTC 2040
			},
			"2000-10-20",
			"2040-05-21",
		},
		{
			[]Date{
				Date{time.Unix(-2927145600, 0)}, // Fri Mar 30 00:00:00 UTC 1877
				Date{time.Unix(2927145600, 0)},  // Wed Oct  4 00:00:00 UTC 2062
				Date{time.Unix(932774400, 0)},   // Sat Jul 24 00:00:00 UTC 1999
				Date{time.Unix(-2927145500, 0)}, // Fri Mar 30 00:01:40 UTC 1877
			},
			"1877-03-30",
			"2062-10-04",
		},
	}

	for _, tc := range cases {
		output := Min(tc.Input[0], tc.Input[1:]...)
		if tc.ExpectedMin != output.String() {
			t.Fatalf("Expected [%s]. Instead received [%s]", tc.ExpectedMin, output.String())
		}

		output = Max(tc.Input[0], tc.Input[1:]...)
		if tc.ExpectedMax != output.String() {
			t.Fatalf("Expected [%s]. Instead received [%s]", tc.ExpectedMax, output.String())
		}
	}
}

// TestDatesSorting tests the sorting interface methods of Dates.
func TestDatesSorting(t *testing.T) {
	dates := Dates{
		Date{time.Unix(-2927145600, 0)}, // Fri Mar 30 00:00:00 UTC 1877
		Date{time.Unix(2927145600, 0)},  // Wed Oct  4 00:00:00 UTC 2062
		Date{time.Unix(932774400, 0)},   // Sat Jul 24 00:00:00 UTC 1999
		Date{time.Unix(1429685920, 0)},  // Wed Apr 22 06:58:40 UTC 2015
		Date{time.Unix(972000000, 0)},   // Fri Oct 20 00:00:00 UTC 2000
		Date{time.Unix(2221257039, 0)},  // Mon May 21 23:50:39 UTC 2040
		Date{time.Unix(-2927145500, 0)}, // Fri Mar 30 00:01:40 UTC 1877
	}
	sort.Sort(dates)
	if !sort.IsSorted(dates) {
		t.Fatalf("Failed to sort the dates: %v", dates)
	}
}

// TestDateAddition tests date arithmetic.
func TestDateAddition(t *testing.T) {
	dStr := NewFromUnixTimestamp(1425081600).AddDays(1).String() // "2015-02-28" + 1 day
	if dStr != "2015-03-01" {
		t.Fatalf("Expected [2015-03-01]. Instead received [%s]", dStr)
	}
	dStr = NewFromUnixTimestamp(1431763320).AddDays(100).String() // "2015-05-16" + 100 days
	if dStr != "2015-08-24" {
		t.Fatalf("Expected [2015-08-24]. Instead received [%s]", dStr)
	}
}

func BenchmarkUnixDate(b *testing.B) {
	ut := time.Now().Unix()

	hrs10 := int64(3600 * 10)
	for i := 0; i < b.N; i++ {
		d1 := NewFromUnixTimestamp(ut + hrs10)
		d1.String()
		d2 := NewFromUnixTimestamp(ut - hrs10)
		d2.String()
	}
}
