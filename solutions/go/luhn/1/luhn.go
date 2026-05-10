package luhn
import "strings"

func Valid(id string) bool {
    check := strings.ReplaceAll(id, " ", "")
    if len(check) <=1{
        return false
    }


    
    sum := 0
    double := false
       
        for i := len(check)-1; i >= 0 ; i--{
            char := check[i]
            if char < '0' || char > '9' {
            return false
        }
            // ASCI for 0 is 48 so
            // eg ASCI 5 is 53, so 53-48 = 5
            number := int(check[i] - '0')
            
            if double{
            	number *= 2
                if number > 9{
                    number -= 9  
                }
            }

             sum += number
            double = !double
    }
    return sum%10 == 0
}
