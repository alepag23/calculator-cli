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
		firstNumber, secondNumber, result float64
		operationSelected                 int
		operationAdd,
		operationSub,
		operationMult,
		operationDiv,
		operationExit string = "1) +", "2) -", "3) *", "4) /", "5) Exit"
		err    error
		runCli bool = true
	)

	for runCli {
		fmt.Println("| -------------- |")
		fmt.Println("| Calculator cli |")
		fmt.Println("| -------------- |")
		fmt.Print("Enter the first number: ")
		fmt.Scan(&firstNumber)
		fmt.Print("Enter the second number: ")
		fmt.Scan(&secondNumber)
		fmt.Print("Select the operation:")
		fmt.Printf("\n %s \n %s \n %s \n %s \n %s \n",
			operationAdd,
			operationSub,
			operationMult,
			operationDiv,
			operationExit,
		)
		fmt.Print("Operation selected: ")
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
			result = 0
			runCli = false
			fmt.Println("Exit success")

		default:
			result = 0
			fmt.Println("operation not valid")
		}

		fmt.Println("Result: ", result)
	}

}
