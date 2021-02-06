package milestone1

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	return celsiusToFahrenheit(c)
}

func Conversion() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "° K is ", celsius, "° C")
	kelvin2 := 0.0
	celsius2 := kelvinToFahrenheit(kelvin2)

	fmt.Print(kelvin2, "° K is ", celsius2, "° C")
}
