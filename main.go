package main

import "fmt"

func main() {
	num1, num2 := readNums()
	answer := num1 + num2
	fmt.Print(answer)
}

func readNums() (num1, num2 int) {
	fmt.Scan(&num1, &num2)
	return num1, num2
}
