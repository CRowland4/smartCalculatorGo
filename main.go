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
	var result int
	for {
		line := readLine()
		switch strings.TrimSpace(line) {
		case "":
			continue
		case "/exit":
			fmt.Println("Bye!")
			return
		case "/help":
			fmt.Println("The program calculates simple math operations using + and -")
			continue
		default:
			result = calculateResult(line)
		}

		fmt.Println(result)
	}
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
