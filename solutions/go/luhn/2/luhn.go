package luhn

func Valid(id string) bool {
    sum := 0
    double := false
    count := 0 

    for i := len(id) - 1; i >= 0; i-- {
        char := id[i]

        // Ignore spaces—don't waste memory on ReplaceAll
        if char == ' ' {
            continue
        }

        // Hard exit if it's not a number
        if char < '0' || char > '9' {
            return false
        }

        // ASCII math: char - 48
        digit := int(char - '0')

        // Every second digit gets doubled
        if double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }

        sum += digit
        double = !double
        count++ 
    }

    // Must be more than one digit & divisible by 10
    return count > 1 && sum%10 == 0
}