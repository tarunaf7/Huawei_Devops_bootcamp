

package main

import (
	"fmt"
	"math"
)


type Quote struct {
	Dollars uint32
	Cents   uint32
}


func (q Quote) String() string {
	return fmt.Sprintf("$%d.%d", q.Dollars, q.Cents)
}


func CreateQuoteFromCount(count int) Quote {
	return CreateQuoteFromFloat(8.99)
}


func CreateQuoteFromFloat(value float64) Quote {
	units, fraction := math.Modf(value)
	return Quote{
		uint32(units),
		uint32(math.Trunc(fraction * 100)),
	}
}