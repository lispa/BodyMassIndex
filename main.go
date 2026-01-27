package main

import (
	"fmt"
)

func main() {
	var weightKg float64
	var heightCm float64
	var indexMass float64

	fmt.Println("-=Calculator BMI=-")

	fmt.Print("Your weight:")
	fmt.Scanln(&weightKg)

	fmt.Print("Your height in cm:")
	fmt.Scanln(&heightCm)

	var heightM float64
	heightM = heightCm / 100.0

	indexMass = weightKg / (heightM * heightM)

	var category string
	switch {
	case indexMass < 18.5:
		category = "Underweight"
	case indexMass < 25.0:
		category = "Normal"
	case indexMass < 30.0:
		category = "Overweight"
	case indexMass < 35.0:
		category = "Obesity I degree"
	case indexMass < 40.0:
		category = "Obesity II degree"
	default:
		category = "Obesity III degree"

	}

	fmt.Println("Your BMI:", indexMass)
	fmt.Println(category)

}
