package ledger

import (
	"errors"
	"strconv"
	"strings"
    "sort"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func GetHeader(locale string) (string, error) {
    if locale == "nl-NL" {
        // "Omschrijving" (12 chars) + 13 spaces = 25 chars
        return "Datum      | Omschrijving              | Verandering  \n", nil
    } else if locale == "en-US" {
        // "Description" (11 chars) + 14 spaces = 25 chars
        return "Date       | Description               | Change       \n", nil
    }
    return "", errors.New("invalid locale")
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
    // Get the header first. 
    // If the locale is bad, we fail immediately.
    s, err := GetHeader(locale)
    if err != nil {
        return "", err
    }

    // Check valid currency
	if currency != "USD" && currency != "EUR" {
        return "", errors.New("invalid currency")
    }
    
	var entriesCopy []Entry
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}
	// if len(entries) == 0 {
       // What on earth!
       // if locale is a certain place with a different date format
       // it should automatically format said date to the 
      // correct locale-date format.
	// if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
	// 		return "", err
	// 	}

	// If there are no entries, just return the header we just got.
    if len(entries) == 0 {
    	return s, nil // Return the header so they see something!
	}
    
	//Sorting date format using sort.SliceStable O(n \log n)
    sort.SliceStable(entriesCopy, func(i, j int) bool {
    if entriesCopy[i].Date != entriesCopy[j].Date {
        return entriesCopy[i].Date < entriesCopy[j].Date
    }
    if entriesCopy[i].Description != entriesCopy[j].Description {
        return entriesCopy[i].Description < entriesCopy[j].Description
    }
    return entriesCopy[i].Change < entriesCopy[j].Change
})


	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
            // We can use split here instead
			// d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]

            // We split the date format into 3 seperate parts.
        	parts := strings.Split(entry.Date,"-")
            if len(parts) != 3{
                co <- struct {
					i int
					s string
					e error
                    }{e: errors.New("Partion of date format failed")}
                return
            }
            

			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			var d string
            // Use our parts to format correctly
			if locale == "nl-NL" {
				d = parts[2] + "-" + parts[1] + "-" + parts[0]
			} else if locale == "en-US" {
				d = parts[1] + "/" + parts[2] + "/" + parts[0]
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				if negative {
					a += "-"
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				a += " "
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
	var al int
            for range a {
                al++
            }

            // Calculate padding with safety guards
            padding := 13 - al
            if padding < 0 {
                padding = 0
            }

            datePadding := 10 - len(d)
            if datePadding < 0 {
                datePadding = 0
            }
            
            co <- struct {
                i int
                s string
                e error
            }{
                i: i, 
                // Use the variables here! 
                s: d + strings.Repeat(" ", datePadding) + " | " + de + " | " +
                    strings.Repeat(" ", padding) + a + "\n",
            }
        }(i, et)
    }
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := range len(entriesCopy) {
		s += ss[i]
	}
	return s, nil
}
