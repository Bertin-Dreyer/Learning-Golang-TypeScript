package booking

import ("time"
        "fmt")

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout,date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    now := time.Now()
    layout := "January 2, 2006 15:04:05"
    scheduledDate, _ := time.Parse(layout,date) 

    return scheduledDate.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    appointment, _ := time.Parse(layout,date)

    return  appointment.Hour() >= 12 && appointment.Hour() < 18
}
// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    dateDesc, _ := time.Parse(layout,date)
	formatedDesc := dateDesc.Format("Monday, January 2, 2006, at 15:04")

    return fmt.Sprintf("You have an appointment on %s.", formatedDesc)
    
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// time.Date(year, month, day, hour, min, sec, nsec, location)
    currentYear := time.Now().Year()
	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)

    return anniversary
}
