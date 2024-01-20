package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	for {
		expression := strings.ToLower(strings.TrimSpace(readLine()))
		switch {
		case expression == "/exit":
			fmt.Println("Bye!")
			return
		case expression == "":
			continue
		case isCommand(expression):
			executeCommand(expression)
		case !isExpressionValid(expression):
			fmt.Println("Invalid expression")
		default:
			fmt.Println(calculateResult(expression))
		}
	}
}

func executeCommand(command string) {
	switch command {
	case "/help":
		fmt.Println("This program calculates simple expressions containing the + and - operators.")
	default:
		fmt.Println("Unknown command")

	}
}

func isCommand(expression string) bool {
	return strings.HasPrefix(expression, "/")
}

func isExpressionValid(expression string) bool {
	if strings.Trim(expression, "+-1234567890 ") != "" {
		return false
	}

	if !strings.ContainsAny(expression[len(expression)-1:], "1234567890") {
		return false
	}

	if hasSpaceWithNoOperator(expression) {
		return false
	}

	return true
}

func hasSpaceWithNoOperator(expression string) bool {
	fields := strings.Fields(expression)
	for i := 0; i < len(fields)-1; i++ {
		if strings.Trim(fields[i], "1234567890") == "" && strings.Trim(fields[i+1], "1234567890") == "" {
			return true
		}
	}

	return false
}

func calculateResult(input string) (result int) {
	for input != "" {
		var digits []rune
		var nonDigits []rune

		input, nonDigits = consumeNonDigits(input)
		input, digits = consumeDigits(input)
		result = newResult(result, getNumFromRunes(digits), getSignFromRunes(nonDigits))
	}

	return result
}

func consumeNonDigits(input string) (newInput string, nonDigits []rune) {
	newInput = input
	for _, char := range newInput {
		if unicode.IsDigit(char) {
			return newInput, nonDigits
		}

		nonDigits = append(nonDigits, char)
		newInput = strings.Replace(newInput, string(char), "", 1)
	}

	return newInput, nonDigits
}

func consumeDigits(input string) (newInput string, digits []rune) {
	newInput = input
	for _, char := range newInput {
		if !unicode.IsDigit(char) {
			return newInput, digits
		}

		digits = append(digits, char)
		newInput = strings.Replace(newInput, string(char), "", 1)
	}

	return newInput, digits
}

func getSignFromRunes(nonDigitRunes []rune) (sign rune) {
	sign = '+'
	for _, char := range nonDigitRunes {
		switch char {
		case '-':
			sign = flipSign(sign)
		default:
			continue
		}
	}

	return sign
}

func flipSign(sign rune) (flippedSign rune) {
	switch sign {
	case '+':
		return '-'
	default:
		return '+'
	}
}

func getNumFromRunes(digitRunes []rune) (num int) {
	stringNum := ""
	for _, digitRune := range digitRunes {
		stringNum += string(digitRune)
	}

	num, _ = strconv.Atoi(stringNum)
	return num
}

func newResult(result, num int, sign rune) (newResult int) {
	switch sign {
	case '+':
		return result + num
	case '-':
		return result - num
	default:
		return 0
	}
}

func readLine() (line string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
