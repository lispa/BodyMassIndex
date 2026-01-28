package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func category(indexMass float64) string {
	switch {
	case indexMass < 18.5:
		return "Underweight"
	case indexMass < 25.0:
		return "Normal"
	case indexMass < 30.0:
		return "Overweight"
	case indexMass < 35.0:
		return "Obesity I degree"
	case indexMass < 40.0:
		return "Obesity II degree"
	default:
		return "Obesity III degree"
	}
}

func makeWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("-=Calculator BMI=-")

	var weightKg *widget.Entry = widget.NewEntry()
	weightKg.SetPlaceHolder("Your weight:")

	var heightCm *widget.Entry = widget.NewEntry()
	heightCm.SetPlaceHolder("Your height in cm:")

	var out *widget.Label = widget.NewLabel("Enter the data and click Calculate")

	var btn *widget.Button = widget.NewButton("Calculate now", func() {
		var kg float64
		var cm float64
		var m float64
		var bmi float64
		var cat string
		var e1 error
		var e2 error

		kg, e1 = strconv.ParseFloat(weightKg.Text, 64)
		cm, e2 = strconv.ParseFloat(heightCm.Text, 64)

		if e1 != nil || e2 != nil || kg <= 0 || cm <= 0 {
			out.SetText("Error: Please enter numbers greater than 0.")
			return
		}

		m = cm / 100.0
		bmi = kg / (m * m)
		cat = category(bmi)

		out.SetText(fmt.Sprintf("BMI: %.2f (%s)", bmi, cat))
	})

	window.SetContent(container.NewVBox(weightKg, heightCm, btn, out))
	return window
}

func main() {
	var a fyne.App = app.New()
	var window fyne.Window = makeWindow(a)
	window.ShowAndRun()
}
