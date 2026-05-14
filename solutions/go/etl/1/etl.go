package etl
import "strings"

func Transform(in map[int][]string) map[string]int {

    // We need to take the value of each letter 
    // in the input and assign it to each letter in the input slice string.

    // Create the result map we will return.
	// result := make(map[string]int)

    // We can also pre-allocate memory for faster execution time.
    // Is this overkill? Yes!
    // Doing it so we can understand how pre-allocation can save memory and lean out execution time
    // We expect roughly 26 keys
     result := make(map[string]int, 26)

    // Access the int map.
    for score, letters := range in {
        //Access each letter in letter slice.
        for _, letter := range letters{
            lowerLetter := strings.ToLower(letter)
             result[lowerLetter] = score
        }
    }
    return result
}
