package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	latinCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers         = "1234567890"
)

func main() {
	variables := make(map[string]string)
	for {
		expression := strings.TrimSpace(readLine())
		switch {
		case expression == "/exit":
			fmt.Println("Bye!")
			return
		case expression == "":
			continue
		case isCommand(expression):
			executeCommand(expression)
		case isCreateVariable(expression):
			variables = assignVariable(expression, variables)
		case strings.ContainsAny(substituteVariables(expression, variables), latinCharacters):
			fmt.Println("Unknown variable")
		case !isExpressionValid(substituteVariables(expression, variables)):
			fmt.Println("Invalid expression")
		default:
			fmt.Println(calculateResult(substituteVariables(expression, variables)))
		}
	}
}

func assignVariable(expression string, variables map[string]string) (updatedVariables map[string]string) {
	expression = strings.Replace(expression, " ", "", -1)
	sides := strings.Split(expression, "=")
	identifier, assignment := strings.TrimSpace(sides[0]), strings.TrimSpace(sides[1])
	if !isVariableCreationValid(sides, variables) {
		return variables
	}

	variables[identifier] = substituteVariables(assignment, variables)
	return variables
}

func substituteVariables(expression string, variables map[string]string) (updatedString string) {
	identifiers := getVariableIdentifiers(variables)
	sort.Slice(identifiers, func(i, j int) bool {
		return len(identifiers[i]) > len(identifiers[j])
	})

	for _, identifier := range identifiers {
		expression = strings.Replace(expression, identifier, variables[identifier], -1)
	}

	return expression
}

func getVariableIdentifiers(variables map[string]string) (identifiers []string) {
	for key, _ := range variables {
		identifiers = append(identifiers, key)
	}

	return identifiers
}

func isVariableCreationValid(sides []string, variables map[string]string) bool {
	if len(sides) != 2 {
		fmt.Println("Invalid assignment")
		return false
	}
	if strings.Trim(sides[0], latinCharacters) != "" {
		fmt.Println("Invalid identifier")
		return false
	}

	if (strings.Trim(sides[1], latinCharacters) != "") && (strings.Trim(sides[1], numbers) != "") {
		fmt.Println("Invalid assignment")
		return false
	}
	_, ok := variables[sides[1]]
	if (strings.Trim(sides[1], latinCharacters) == "") && !ok {
		fmt.Println("Unknown variable")
		return false
	}

	return true
}

func isCreateVariable(expression string) bool {
	return strings.Contains(expression, "=")
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
