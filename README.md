# BodyMassIndex

# BMI Calculator (Fyne GUI Version)

A simple desktop (window) BMI calculator written in Go using the **Fyne** UI toolkit.  
The app asks for **weight (kg)** and **height (cm)**, then shows the **BMI value** and **category**.

## Features
- Desktop window interface (Fyne)
- Input fields for weight and height
- BMI calculation + category output
- Basic input validation (numbers > 0)

## BMI Formula
BMI = weight(kg) / (height(m) × height(m))

Height is entered in centimeters and converted to meters inside the app.

## Requirements
- Go installed (`go version`)
- C compiler for Windows builds (Fyne uses OpenGL):
  - Recommended: **MSYS2 UCRT64 + GCC**
- Internet access on first build (to download dependencies)

## Install Dependencies
Inside the project folder (where `main.go` is located):

```powershell
go mod init bmi
go get fyne.io/fyne/v2@latest
go mod tidy
