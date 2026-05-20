package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// working with golang calculator
/* Basics functionality of a calculator
1. allow user enters two numbers and store into their variables (firstNumber and secondNumber)
2. suggest for operation type(addition, subtraction, multiplication, division)
3. perform operation based on the type of operator */

// addition
func addition(x, y int) int {
	return x + y
}

// subtraction
func subtraction(x, y float64) float64 {
	return x - y
}

// multiplication
func multiplication(x, y int) int {
	return x * y
}

func division(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	return 1
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter first number: ")
	reader.Scan()
	input, _ := strconv.Atoi(reader.Text())
	// strconv.Atoi(reader.Text()) is same as strconv.ParseInt(reader.Text(), 10, 0)

	reader2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter second number: ")
	reader2.Scan()
	input2, _ := strconv.Atoi(reader2.Text())

	var operator string

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		resultAdd2 := addition(input, input2)
		fmt.Println("results after addition: ", resultAdd2)
	case "-":
		resultSubtraction2 := subtraction(float64(input), float64(input2))
		fmt.Println("results after subtraction: ", resultSubtraction2)
	case "*":
		resultMultiplication2 := multiplication(input, input2)
		fmt.Println("results after multiplication: ", resultMultiplication2)
	case "/":
		resultDivision2 := division(float64(input), float64(input2))
		fmt.Println("results after division: ", resultDivision2)
	default:
		fmt.Println("Incorrect operator type")
	}

}
