package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int // day in which the expense was made
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var outRecords []Record
	for _, record := range in {
		if predicate(record) {
			outRecords = append(outRecords, record)
		}
	}
	return outRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	predicate := func(record Record) bool {
		if record.Day >= p.From && record.Day <= p.To {
			return true
		}
		return false
	}
	return predicate
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	predicate := func(record Record) bool {
		if record.Category == c {
			return true
		}
		return false
	}
	return predicate
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	totalAmount := 0.0
	for _, record := range in {
		if ByDaysPeriod(p)(record) {
			totalAmount += record.Amount
		}
	}
	return totalAmount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {

	var recordsFilteredByCategory []Record

	// first filter by category
	for _, record := range in {
		if ByCategory(c)(record) {
			recordsFilteredByCategory = append(recordsFilteredByCategory, record)
		}
	}

	// if no records of category c, return early errpr
	if len(recordsFilteredByCategory) == 0 {
		errMessage := fmt.Sprintf("unknown category %s", c)
		return 0.0, errors.New(errMessage)
	}

	// otherwise if records of category c present, calculate expenses
	return TotalByPeriod(recordsFilteredByCategory, p), nil
}
