package main

import "fmt"

type Kilometers float64
type Meters float64
type Miles float64

// 변환 메서드 (km -> Mile, Meter -> Mile)
func (km Kilometers) ToMiles() Miles {
	return Miles(km / 1.609)
}
func (m Meters) ToMiles() Miles {
	return Miles(m / 1609)
}

func main() {
	kmph := Kilometers(151)
	fmt.Printf("%0.2f km/h equals %0.2f mile/p\n", kmph, kmph.ToMiles())
	meter := Meters(151000)
	fmt.Printf("%0.2f meter equals %0.2f miles\n", meter, meter.ToMiles())
}
