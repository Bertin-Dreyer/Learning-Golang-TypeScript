package isbnverifier

func IsValidISBN(isbn string) bool {
    sum := 0
    step := 10

    for i := 0; i < len(isbn); i++ {
        char := isbn[i]

        if char == '-' {
            continue // Path A: Skip it
        }

        if char == 'X' {
            // Path B: Handle X
            if step != 1 {
                return false
            }
            sum += 10 * step // (10 * 1)
            step--
            continue // IMPORTANT: Jump to next iteration now
        }

        if char >= '0' && char <= '9' {
            // Path C: Handle Numbers
            digit := int(char - '0')
            sum += digit * step
            step--
        } else {
            // Path D: Garbage character
            return false
        }
    }

    return step == 0 && sum%11 == 0
}