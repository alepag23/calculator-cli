package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}

	return a / b, nil
}

func main() {

	var (
		result float64
		runCli bool = true
	)

	for runCli {
		fmt.Println("| -------------- |")
		fmt.Println("| Calculator cli |")
		fmt.Println("| -------------- |")
		fmt.Print("Enter the first number: ")

		firstNumber := 0.0
		_, err := fmt.Scan(&firstNumber)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Print("Enter the second number: ")

		secondNumber := 0.0
		_, err = fmt.Scan(&secondNumber)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Print("Select the operation:")
		operationAdd := "1) +"
		operationSub := "2) -"
		operationMult := "3) *"
		operationDiv := "4) /"
		operationExit := "5) Exit"
		fmt.Printf("\n %s \n %s \n %s \n %s \n %s \n",
			operationAdd,
			operationSub,
			operationMult,
			operationDiv,
			operationExit,
		)

		fmt.Print("Operation selected: ")
		operationSelected := 0
		fmt.Scan(&operationSelected)

		switch operationSelected {
		case 1:
			result = add(firstNumber, secondNumber)
		case 2:
			result = subtract(firstNumber, secondNumber)
		case 3:
			result = multiply(firstNumber, secondNumber)
		case 4:
			result, err = divide(firstNumber, secondNumber)

			if err != nil {
				fmt.Println("Error:", err)
			}
		case 5:
			runCli = false
			fmt.Println("Exit success")

		default:
			result = 0
			fmt.Println("operation not valid")
		}

		if runCli {
			fmt.Printf("Result: %.2f \n", result)
		}
	}

}
