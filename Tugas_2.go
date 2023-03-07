package main

import (
	"fmt"
)

func main() {
	var temp float64
	var scale string

	fmt.Print("Enter temperature: ")
	fmt.Scanln(&temp)

	fmt.Print("Enter scale (C, F, K): ")
	fmt.Scanln(&scale)

	switch scale {
	case "C":
		fmt.Printf("%.2f°C = %.2f°F\n", temp, celsiusToFahrenheit(temp))
		fmt.Printf("%.2f°C = %.2fK\n", temp, celsiusToKelvin(temp))
	case "F":
		fmt.Printf("%.2f°F = %.2f°C\n", temp, fahrenheitToCelsius(temp))
		fmt.Printf("%.2f°F = %.2fK\n", temp, fahrenheitToKelvin(temp))
	case "K":
		fmt.Printf("%.2fK = %.2f°C\n", temp, kelvinToCelsius(temp))
	default:
		fmt.Println("Invalid scale. Please enter C, F, or K.")
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit + 459.67) * 5 / 9
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
