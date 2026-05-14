package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    totalBirds := 0
	for _, birds := range  birdsPerDay{
        totalBirds += birds
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    start := (week - 1) *7
    end := start +7
	totalBirds := 0

    for _ , birds := range birdsPerDay[start:end]{
        totalBirds += birds
    }
    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0 ;i < len(birdsPerDay) ; i += 2{
        birdsPerDay[i] +=1
    }
    return birdsPerDay
}
