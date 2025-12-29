package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	c := 0
	for _, k := range birdsPerDay {
		c += k
	}

	return c
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	c := 0
	if week == 1 {
		for i := 0; i <= 6; i++ {
			c += birdsPerDay[i]
		}
		return c
	}
	c = 0
	if week != 1 {
		w := (week * 7) - 7
		for i := w; i <= w+6; i++ {
			c += birdsPerDay[i]
		}
		return c
	}

	return c
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for k, v := range birdsPerDay {
		if k%2 == 0 {
			birdsPerDay[k] = v + 1
		}
	}

	return birdsPerDay
}
