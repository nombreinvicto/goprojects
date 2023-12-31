package purchase

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection.
// It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {

	var selected string
	if option1 < option2 {
		selected = option1
	} else {
		selected = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", selected)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	var resellPrice float64
	if age < 3 {
		resellPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		resellPrice = originalPrice * 0.7
	} else {
		resellPrice = originalPrice * 0.5
	}
	return resellPrice
}
