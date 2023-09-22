package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, ncows int) (float64, error) {
	fodderAmount, err := fc.FodderAmount(ncows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (fodderAmount / float64(ncows)) * fatteningFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, ncows int) (float64, error) {
	if ncows > 0 {
		return DivideFood(fc, ncows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

// create a custom error class
type InvalidCowsError struct {
	ncows   int
	message string
}

func (ic *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ic.ncows, ic.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(ncows int) error {
	if ncows < 0 {
		return &InvalidCowsError{
			ncows:   ncows,
			message: "there are no negative cows",
		}
	} else if ncows == 0 {
		return &InvalidCowsError{
			ncows:   ncows,
			message: "no cows don't need food",
		}
	} else {
		return nil
	}
}
