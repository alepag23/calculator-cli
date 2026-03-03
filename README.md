# Calculator CLI

A simple CLI application in Go to emulate a simple calculator.

## Operations
- add
- subtract
- multiply
- divide

## Setup
```bash
mkdir calculator-cli
go mod init calculator-cli
touch main.go
```

## Usage
```bash
go run main.go
```

## Example Output
```text
| -------------- |
| Calculator cli |
| -------------- |
Enter the first number: 10
Enter the second number: 2
Select the operation:
 1) +
 2) -
 3) *
 4) /
 5) Exit
Operation selected: 4
Result: 5.00
```
## What I Learned
- How to use a for loop with a boolean condition
- If statements and switch cases
- Error handling with multiple return values
- Reading user input with fmt.Scan