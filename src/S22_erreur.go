package main

import (
	"fmt"
	"math"
)

type errorDivBy0 struct {
	Dividend int
}

func (e *errorDivBy0) Error() string {
	return fmt.Sprintf("%d can not be divided by 0", e.Dividend)
}

func div(dividend, divisor int) (float64, error) {
	if divisor == 0 {
		return math.NaN(), &errorDivBy0{Dividend: dividend}
		//return math.NaN(), errors.New(fmt.Sprintf("%d can not be divided by 0", dividend))
	}

	return float64(dividend) / float64(divisor), nil
}

func main() {
	if result, err := div(5, 2); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	if result, err := div(5, 0); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	result, _ := div(5, 2)
	fmt.Println(result)
}
