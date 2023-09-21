package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	newBattery := car.battery - car.batteryDrain
	if newBattery >= 0 {
		car.battery = newBattery
		car.distance = car.distance + car.speed
	}

}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	totalDistanceCovered := 0
	prevDistance := car.distance

	for {
		car.Drive()
		currentDistance := car.distance
		distanceCovered := currentDistance - prevDistance
		if distanceCovered == 0 {
			break
		}
		totalDistanceCovered += distanceCovered
		prevDistance = currentDistance

	}

	if totalDistanceCovered < trackDistance {
		return false
	}
	return true

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
