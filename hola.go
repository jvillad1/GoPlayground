package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// Welcome messages
	fmt.Println("¡Hola, Mundo!")
	fmt.Println(quote.Go())
	fmt.Println()

	// Different ways to declare variables in Go
	demonstrateVariables()
	fmt.Println()

	// Constants and calculations
	demonstrateConstants()
	fmt.Println()

	// Working with user data
	demonstrateUserProfile()
}

// demonstrateVariables shows different variable declaration patterns
func demonstrateVariables() {
	fmt.Println("=== Variable Declarations ===")

	// Method 1: Explicit type declaration
	var name string = "John Doe"
	var age int = 25

	// Method 2: Type inference with var
	var city = "São Paulo"
	var isActive = true

	// Method 3: Short variable declaration (most common)
	country := "Brazil"
	salary := 5000.50

	// Method 4: Multiple variable declaration
	firstName, lastName := "Ana", "Silva"

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("City: %s, Active: %t\n", city, isActive)
	fmt.Printf("Country: %s, Salary: R$ %.2f\n", country, salary)
	fmt.Printf("Full Name: %s %s\n", firstName, lastName)
}

// demonstrateConstants shows constant usage and calculations
func demonstrateConstants() {
	fmt.Println("=== Constants and Calculations ===")

	// Constants - values that don't change
	const (
		pi          = 3.14159
		daysInWeek  = 7
		companyName = "Tech Corp"
		taxRate     = 0.15 // 15% tax
	)

	// Using constants in calculations
	radius := 5.0
	area := pi * radius * radius

	grossSalary := 1000.0
	netSalary := grossSalary * (1 - taxRate)

	fmt.Printf("Company: %s\n", companyName)
	fmt.Printf("Circle area (radius %.1f): %.2f\n", radius, area)
	fmt.Printf("Gross salary: R$ %.2f\n", grossSalary)
	fmt.Printf("Net salary (after %.0f%% tax): R$ %.2f\n", taxRate*100, netSalary)
}

// demonstrateUserProfile shows a more realistic example
func demonstrateUserProfile() {
	fmt.Println("=== User Profile Example ===")

	// Struct-like data using individual variables
	userName := "Maria Santos"
	userAge := 28
	userEmail := "maria.santos@email.com"
	isVerified := true
	accountBalance := 1250.75

	// Display user information
	fmt.Printf("User Profile:\n")
	fmt.Printf("  Name: %s\n", userName)
	fmt.Printf("  Age: %d years old\n", userAge)
	fmt.Printf("  Email: %s\n", userEmail)
	fmt.Printf("  Verified: %t\n", isVerified)
	fmt.Printf("  Balance: R$ %.2f\n", accountBalance)

	// Status based on conditions
	status := "Regular"
	if accountBalance > 1000 {
		status = "Premium"
	}
	fmt.Printf("  Status: %s\n", status)
}
