package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func sum(firstNumber float64, secondNumber float64) float64 {
	return firstNumber + secondNumber
}

func sub(firstNumber float64, secondNumber float64) float64 {
	return firstNumber - secondNumber
}

func mult(firstNumber float64, secondNumber float64) float64 {
	return firstNumber * secondNumber
}

func div(firstNumber float64, secondNumber float64) float64 {
	return firstNumber / secondNumber
}

func pow(firstNumber float64, secondNumber float64) float64 {
	return math.Pow(firstNumber, secondNumber)
}

func mod(firstNumber float64, secondNumber float64) float64 {
	return math.Mod(firstNumber, secondNumber)
}

func in[T comparable](valueToFind T, slice []T) bool {
	for _, value := range slice {
		if value == valueToFind {
			return true
		}
	}
	return false
}

func readAndClearBuffer(reader *bufio.Reader, destVar interface{}) error {
	_, readError := fmt.Fscan(reader, destVar)
	_, clearError := reader.ReadString('\n')
	if readError != nil {
		return readError
	}
	if clearError != nil {
		return clearError
	}
	return nil
}

func main() {
	validOperators := []string{"+", "-", "*", "/", "**", "//"}
	var selectedOperator string
	var firstNumber float64
	var secondNumber float64
	reader := bufio.NewReader(os.Stdin)

	isFirstTime := true
	for {
		if isFirstTime {
			fmt.Print("Choose operator ['+', '-', '*', '/', '**', '//'] or 'q' to quit: ")
		} else {
			fmt.Print("\nChoose operator ['+', '-', '*', '/', '**', '//'] or 'q' to quit: ")
		}
		isFirstTime = false

		if readAndClearBuffer(reader, &selectedOperator) != nil {
			fmt.Print("Invalid input. Enter a valid operator\n")
			continue
		} else if in(selectedOperator, []string{"q", "Q"}) {
			fmt.Print("Bye bye")
			return
		} else if !in(selectedOperator, validOperators) {
			fmt.Print("Invalid input. Enter a valid operator\n")
			continue
		}

		fmt.Print("First number: ")
		if readAndClearBuffer(reader, &firstNumber) != nil {
			fmt.Print("Invalid input. Enter a valid number\n")
			continue
		}

		fmt.Print("Second number: ")
		if readAndClearBuffer(reader, &secondNumber) != nil {
			fmt.Print("Invalid input. Enter a valid number\n")
			continue
		}

		if selectedOperator == "+" {
			fmt.Printf("%g + %g = %g\n", firstNumber, secondNumber, sum(firstNumber, secondNumber))
		} else if selectedOperator == "-" {
			fmt.Printf("%g - %g = %g\n", firstNumber, secondNumber, sub(firstNumber, secondNumber))
		} else if selectedOperator == "*" {
			fmt.Printf("%g * %g = %g\n", firstNumber, secondNumber, mult(firstNumber, secondNumber))
		} else if selectedOperator == "/" {
			if secondNumber == 0 {
				fmt.Print("Operation invalid, division by zero\n")
			} else {
				fmt.Printf("%g / %g = %g\n", firstNumber, secondNumber, div(firstNumber, secondNumber))
			}
		} else if selectedOperator == "**" {
			fmt.Printf("%g ** %g = %g\n", firstNumber, secondNumber, pow(firstNumber, secondNumber))
		} else if selectedOperator == "//" {
			if secondNumber == 0 {
				fmt.Print("Operation invalid, division by zero\n")
			} else {
				fmt.Printf("%g // %g = %g\n", firstNumber, secondNumber, mod(firstNumber, secondNumber))
			}
		}
	}
}
