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
	//bufio method 

	// no 1
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter first number: ")
	reader.Scan()
	input, _ := strconv.Atoi(reader.Text())

	//no 2
	reader2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter first number: ")
	reader2.Scan()
	input2, _ := strconv.Atoi(reader2.Text())

	//logic

	resultAdd2 := addition(input, input2)
	fmt.Println("resultAdd", resultAdd2)

	resultSubtraction2 := subtraction(float64(input), float64(input2))
	fmt.Println("resultSubtraction", resultSubtraction2)

	resultMultiplication2 := multiplication(input, input2)
	fmt.Println("resultMultiplication", resultMultiplication2)

	resultDivision2 := division(float64(input), float64(input2))
	fmt.Println("resultDivision", resultDivision2)

	/// TRY
	//! using inbuilt fmt package use Scan and store to the int variable defined

	//define variables
	var firstNumber int
	var secondNumber int

	//scan
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	resultAdd := addition(firstNumber, secondNumber)
	fmt.Println("resultAdd", resultAdd)

	resultSubtraction := subtraction(float64(firstNumber), float64(secondNumber))
	fmt.Println("resultSubtraction", resultSubtraction)

	resultMultiplication := multiplication(firstNumber, secondNumber)
	fmt.Println("resultMultiplication", resultMultiplication)

	resultDivision := division(float64(firstNumber), float64(secondNumber))
	fmt.Println("resultDivision", resultDivision)
}
