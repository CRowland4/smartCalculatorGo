package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nums []int
	for {
		input := readLine()
		switch strings.TrimSpace(input) {
		case "":
			continue
		case "/exit":
			fmt.Println("Bye!")
			return
		case "/help":
			fmt.Println("The program calculates the sum of numbers")
			continue
		default:
			nums = parseNums(input)
		}

		fmt.Println(sum(nums))
	}
}

func sum(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}

	return result
}

func parseNums(input string) (nums []int) {
	stringNums := strings.Split(input, " ")
	for _, num := range stringNums {
		integer, _ := strconv.Atoi(num)
		nums = append(nums, integer)
	}

	return nums
}

func readLine() (line string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
