package hamming

import "errors"

func Distance(a, b string) (int, error) {
    hammingLength := 0
    if len(a) == len(b){
        for i := 0; i < len(a); i++{
            if a[i] != b[i]{
                hammingLength ++
            }
        }
    }else{
        return 0, errors.New("lengths must match")
    }
    return hammingLength, nil
	panic("Implement the Distance function")
}
