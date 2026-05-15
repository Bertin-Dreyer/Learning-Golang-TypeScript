// Package bafflingbirthdays provides utilities to simulate and calculate 
// the birthday paradox, demonstrating that shared birthdays are more common 
// than intuitively expected.
package bafflingbirthdays

import (
	"math/rand"
	"time"
)

// SharedBirthday checks if a slice of dates contains at least two people 
// who share the same birthday (month and day).
// It uses a map to achieve O(n) time complexity.
func SharedBirthday(dates []time.Time) bool {
	if len(dates) <= 1 {
		return false
	}
	// Pre-allocating map size based on input slice length to optimize memory.
	seen := make(map[string]bool, len(dates))

	for _, date := range dates {
		// We format to "MM-DD" to strip year and time, focusing only on the birthday.
		key := date.Format("01-02")

		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

// RandomBirthdates generates a slice of random time.Time objects.
// It assumes a non-leap year (365 days) and a uniform distribution.
func RandomBirthdates(size int) []time.Time {
	birthdays := make([]time.Time, size)
	// Using a fixed non-leap year (2025) as the base for calendar math.
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < size; i++ {
		// Generate a random offset within 365 days.
		days := rand.Intn(365)
		birthdays[i] = start.AddDate(0, 0, days)
	}
	return birthdays
}

// EstimatedProbability calculates the theoretical probability of at least 
// two people sharing a birthday in a group of a given size.
// Returns the result as a percentage (0-100).
func EstimatedProbability(size int) float64 {
	// Probability that no two people share a birthday.
	// We start at 1.0 (100% chance no match with 0 people).
	probNoMatch := 1.0

	for i := 0; i < size; i++ {
		// Each new person has fewer "available" days to not match the others.
		// Formula: (365-0)/365 * (365-1)/365 * ... * (365-(n-1))/365
		probNoMatch *= float64(365-i) / 365.0
	}

	// Subtract the "no match" chance from 1 to get the "at least one match" chance.
	return (1.0 - probNoMatch) * 100.0
}