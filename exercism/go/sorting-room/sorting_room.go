package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %0.1f",
		float64(nb.Number()))
}

// FancyNumber Class
type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

// any class implementing Value() string satisfies this interface
type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		num, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return 0
		}
		return num
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {

	switch fnb.(type) {
	case FancyNumber:
		return fmt.Sprintf("This is a fancy box containing the number %0.1f",
			float64(ExtractFancyNumber(fnb)))
	default:
		return fmt.Sprintf("This is a fancy box containing the number 0.0")
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {

	// i.(type) only works with interfaces
	switch i.(type) {
	case float64:
		return DescribeNumber(i.(float64))
	case int:
		return DescribeNumber(float64(i.(int)))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		return "Return to sender"
	}
}
