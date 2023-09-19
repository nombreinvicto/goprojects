package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirdCount int
	for _, count := range birdsPerDay {
		totalBirdCount += count
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalBirdCount int
	//week1 = 0 to 7
	//week2 = 7 to 14
	//week3 = 14 to 21
	for _, count := range birdsPerDay[(week-1)*7 : week*7] {
		totalBirdCount += count
	}
	return totalBirdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}
