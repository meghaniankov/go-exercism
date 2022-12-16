package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, dayCount := range birdsPerDay {
		total += dayCount
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekEnd := week * 7
	weekStart := weekEnd - 7

	return TotalBirdCount(birdsPerDay[weekStart : weekEnd])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < (len(birdsPerDay)); i +=2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
