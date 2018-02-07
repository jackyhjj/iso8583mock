// Package helpers implements commonly used functions (datetime manipulation) //
package helpers

import (
	"time"
)

const (
	dateTimeLayout        string = "2006-01-02 15:04:05"
	timestampLayout       string = "20060102150405"
	timeLayout            string = "15:04:05"
	dateLayout            string = "2006-01-02"
	iso8583DateTimeLayout string = "0102150405"
	iso8583DateLayout     string = "0102"
	iso8583TimeLayout     string = "150405"
)

func timeLoadLocation() *time.Location {
	time, _ := time.LoadLocation("Asia/Jakarta")
	return time
}

// GetCurrentDateTime - get current date and time
// GetCurrentDateTime, get current date and time //
// output time.time
func GetCurrentDateTime() time.Time {
	currentDateTime := time.Now().In(timeLoadLocation())
	return currentDateTime
}

// GetCurrentDate - get today date only
// GetCurrentDate, get today date only //
// output time.Time
func GetCurrentDate() time.Time {
	currentDate := time.Now().Local()
	return currentDate
}

// GetCurrentTime - get today date only
// GetCurrentTime, get today date only //
// output time.Time
func GetCurrentTime() time.Time {
	currentTime := time.Now().Local()
	return currentTime
}

// TimeToString - convert from time.Time into string
// TimeToString, convert from time.Time into string //
// input  dateInput time.Time
// output string
func TimeToString(dateInput time.Time) string {
	return dateInput.Format(dateTimeLayout)
}

// DateTimeToISO8583String - convert from time.Time into ISO8583 String
// DateTimeToISO8583String, convert from time.Time into ISO8583 String //
// input  dateInput time.Time
// output string
func DateTimeToISO8583String(dateInput time.Time) string {
	return dateInput.Format(iso8583DateTimeLayout)
}

// DateToISO8583String - convert from time.Time into ISO8583 String
// DateToISO8583String, convert from time.Time into ISO8583 String //
// input  dateInput time.Time
// output string
func DateToISO8583String(dateInput time.Time) string {
	return dateInput.Format(iso8583DateLayout)
}

// TimeToISO8583String - convert from time.Time into ISO8583 String
// TimeToISO8583String, convert from time.Time into ISO8583 String //
// input  dateInput time.Time
// output string
func TimeToISO8583String(dateInput time.Time) string {
	return dateInput.Format(iso8583TimeLayout)
}

// StringToTime - convert from date string into time.Time
// StringToTime, convert from date string into time.Time //
// input dateInput string
// output error
// output time.Time
func StringToTime(dateInput string) (time.Time, error) {
	result, err := time.Parse(dateTimeLayout, dateInput)
	return result, err
}

// GetCurrentTimestamp - get current timestamp (string) with layout
// GetCurrentTimestamp, get current timestamp (string) with layout //
// output string
func GetCurrentTimestamp() string {
	t := time.Now()
	return t.Format(timestampLayout)
}

// CountDays - count days from 2 date input
// CountDays, count days from 2 date input //
// input dateStart string
// input dateEnd string
// output bool
// output error
func CountDays(dateStart, dateEnd string) (int, error) {

	// Parse firstInput into datetime format //
	firstDate, err := time.Parse(dateTimeLayout, dateStart)
	if err != nil {
		return 0, err
	}

	// Parse secondInput into datetime format //
	secondDate, err := time.Parse(dateTimeLayout, dateEnd)
	if err != nil {
		return 0, err
	}

	days := secondDate.Sub(firstDate).Hours() / 24

	return int(days), err
}

// IndonesianMount - define default Indonesian month name
var IndonesianMount = map[int]string{
	1:  "Januari",
	2:  "Februari",
	3:  "Maret",
	4:  "April",
	5:  "Mei",
	6:  "Juni",
	7:  "Juli",
	8:  "Agustus",
	9:  "September",
	10: "Oktober",
	11: "November",
	12: "Desember",
}
