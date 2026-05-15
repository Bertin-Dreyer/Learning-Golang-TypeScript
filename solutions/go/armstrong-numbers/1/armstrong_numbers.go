package armstrongnumbers
import "math"

func IsNumber(n int) bool {
    original := n
    
    // Count digits
    digitCount := 0
    temp := n
    for temp > 0 {
        digitCount++
        temp /= 10
    }
    
    // Extract digits, raise to power, sum
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += int(math.Pow(float64(digit), float64(digitCount)))
        n /= 10
    }
    
    // Compare
    return sum == original
}
