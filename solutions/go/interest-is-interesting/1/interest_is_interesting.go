package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
    case balance < 0:
        return 3.213
    case balance < 1000:
        return 0.5
    case balance < 5000:
        return 1.621
    default:
        return 2.475
    }
}

// Interest calculates the interest amount for the provided balance.
func Interest(balance float64) float64 {
    // We convert the float32 rate to float64 to keep the precision during math.
    return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate adds the interest to the current balance.
func AnnualBalanceUpdate(balance float64) float64 {
    return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum years to hit a target.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0
    
    // Using a for loop as a "while" loop. 
    // Go While loop.
    for balance < targetBalance {
        balance = AnnualBalanceUpdate(balance)
        years++
    }
    
    return years
}