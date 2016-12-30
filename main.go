package main

import (
	"fmt"
	"os"
	"strconv"
)

func toKilometers(unit float64) float64 {
	baseMile := 1.609
	return unit * baseMile
}

func toLitres(unit float64) float64 {
	baseGallon := 3.79
	return unit * baseGallon
}

func toLitrePer100Kilometers(unit float64) float64 {
	kilometers := toKilometers(unit)
	return toLitres(1) / kilometers * 100
}

func main() {
	arguments := os.Args[1:]
	unit, _ := strconv.ParseFloat(arguments[0], 64)
	result := 0.0
	var originalUnit string
	var resultUnit string

	switch arguments[1] {
	case "mi":
		result = toKilometers(unit)
		originalUnit = "miles"
		resultUnit = "kilometres"
	case "g":
		result = toLitres(unit)
		originalUnit = "gallons"
		resultUnit = "litres"
	case "mpg":
		result = toLitrePer100Kilometers(unit)
		originalUnit = "mpg"
		resultUnit = "l/100km"
	}

	fmt.Printf("%v %v equals to %.2f %v\n", unit, originalUnit, result, resultUnit)
}
