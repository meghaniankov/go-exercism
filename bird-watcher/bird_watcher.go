package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for i := 0; i < (len(birdsPerDay)); i++ {
		total = total + birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekEnd := week * 7
	weekStart := weekEnd - 7
	total := 0

	for i := weekStart; i < (weekEnd); i++ {
		total = total + birdsPerDay[i]
	}
	return total

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < (len(birdsPerDay)); i++ {
		if (i %2 == 0) {
			birdsPerDay[i]++
		} 
	}
	return birdsPerDay
}
