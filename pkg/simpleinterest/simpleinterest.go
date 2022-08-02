package simpleinterest

import "fmt"

func init() {
	fmt.Println("simpleinterest::init()")
}

func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
