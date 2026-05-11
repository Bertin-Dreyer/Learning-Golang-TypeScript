package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

// Number cleans and validates a North American Numbering Plan phone number.
func Number(phoneNumber string) (string, error) {
	// 1. Strip all common formatting characters
	replacer := strings.NewReplacer("+", "", "(", "", ")", "", "-", "", ".", "", " ", "")
	clean := replacer.Replace(phoneNumber)

	// 2. Ensure no letters or extra punctuation remain
	for _, r := range clean {
		if !unicode.IsDigit(r) {
			return "", fmt.Errorf("invalid characters in number")
		}
	}

	// 3. Handle the 11-digit country code ('1')
	if len(clean) == 11 {
		if clean[0] != '1' {
			return "", fmt.Errorf("11-digit number must start with 1")
		}
		clean = clean[1:] // Trim the country code
	}

	// 4. Must be exactly 10 digits now
	if len(clean) != 10 {
		return "", fmt.Errorf("must be 10 or 11 digits")
	}

	// 5. Area (index 0) and Exchange (index 3) cannot start with 0 or 1
	if clean[0] < '2' {
		return "", fmt.Errorf("area code cannot start with 0 or 1")
	}
	if clean[3] < '2' {
		return "", fmt.Errorf("exchange code cannot start with 0 or 1")
	}

	return clean, nil
}

// AreaCode returns the first three digits of a valid phone number.
func AreaCode(phoneNumber string) (string, error) {
	clean, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return clean[0:3], nil
}

// Format returns a formatted version of the phone number: (XXX) XXX-XXXX
func Format(phoneNumber string) (string, error) {
	clean, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", clean[0:3], clean[3:6], clean[6:]), nil
}