package anagram
import 	("strings"
		"maps")

func keyValueCounter (input string) map[rune]int{
    counts := make(map[rune]int)

    for _, char := range input {
    	counts[char]++
	}
	return counts
}

func Detect(subject string, candidates []string) []string {
	var result []string

   	// Convert subject to lowercase.
    lowerSubject := strings.ToLower(subject)
    subjectMap := keyValueCounter(lowerSubject)

    // Create the candidate map in-memory
    candidateMap := make(map[rune]int)
    
    // More idiomatic way of looping instead of using for i++ loop.
	for _, candidate := range candidates {
    	if len(subject) != len(candidate) {
        	continue
    }
        // Convert candidate to lowercase.
        lowerCandidate := strings.ToLower(candidate)

        // Check if subject is not the same as the subject eg. stop == stop
        if lowerSubject == lowerCandidate{
            continue
        }

        // Clear prev candidateMap from memory.
        clear(candidateMap)

        // Now we map the key-value pairs.
        for _, char := range lowerCandidate {
            candidateMap[char]++
        }

        // Compare counts of letters of subject (runes) to the candidate counts.
        if maps.Equal(subjectMap, candidateMap) {
            result = append(result, candidate)
        }
    }
    return result
}
