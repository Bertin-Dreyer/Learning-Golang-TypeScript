package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
    seen := make(map[int]bool)
    total := 0
    
	for _, diValue := range divisors{
        if diValue == 0{
            continue
        }
        for i := 1 ; ; i++{
            multiple := diValue * i

            if multiple >= limit {
				break
			}

            if !seen[multiple] {
				total += multiple
				seen[multiple] = true
			}
        }
    }
    return total
}
