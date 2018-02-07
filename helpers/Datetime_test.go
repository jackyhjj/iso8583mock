package helpers

import (
	"fmt"
	//"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"github.com/magiconair/properties/assert"
)

func TestGetDateRange(t *testing.T) {
	var (
		dateStart = "2017-08-14 15:04:05"
		dateEnd   = "2017-09-14 15:04:05"
		//expected  int    = 31
	)

	days, err := CountDays(dateStart, dateEnd)
	if err != nil {
		t.Error(err)
	}

	//assert.Equal(t, expected, days, ("Result should be %d, be actual get %d", expected, days))
	fmt.Println(days)
}

func TestGetCurrentDateTime(t *testing.T) {
	date := GetCurrentDateTime()
	fmt.Println(date)
}

func TestGetCurrentDate(t *testing.T) {
	date := GetCurrentDate()
	fmt.Println(date)
}

func TestGetCurrentTime(t *testing.T) {
	time := GetCurrentTime()
	fmt.Println(time)
}

func TestTimeToString(t *testing.T) {
	date := GetCurrentDate()
	dateStr := TimeToString(date)
	fmt.Println(dateStr)
}

func TestStringToTime(t *testing.T) {
	date := GetCurrentDate()
	dateStr := TimeToString(date)
	dateDate, err := StringToTime(dateStr)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(dateDate)
}

func TestGetCurrentTimestamp(t *testing.T) {
	timestamp := GetCurrentTimestamp()
	fmt.Println(timestamp)
}

func TestDateTimeToISO8583String(t *testing.T) {
	datetime := time.Date(2018,time.January,31,16,49,30,100,time.Local)
	fmt.Println(datetime)
	isoDateTime := DateTimeToISO8583String(datetime)
	assert.Equal(t,"0131164930",isoDateTime)
}

func TestDateToISO8583String(t *testing.T) {
	datetime := time.Date(2018,time.January,31,16,49,30,100,time.Local)
	fmt.Println(datetime)
	isoDateTime := DateToISO8583String(datetime)
	assert.Equal(t,"0131",isoDateTime)
}

func TestTimeToISO8583String(t *testing.T) {
	datetime := time.Date(2018,time.January,31,16,49,30,100,time.Local)
	fmt.Println(datetime)
	isoDateTime := TimeToISO8583String(datetime)
	assert.Equal(t,"164930",isoDateTime)
}