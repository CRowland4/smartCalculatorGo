package main

import (
	"bufio"
	"errors"
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

type Stack struct {
	storage []string
}

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
		fmt.Println("This program is a calculator that supports +, -, *, /, parenthesis and variable assignment.")
	default:
		fmt.Println("Unknown command")

	}
}

func isCommand(expression string) bool {
	return strings.HasPrefix(expression, "/")
}

func isExpressionValid(expression string) bool {
	if strings.Trim(expression, "()*/+-1234567890 ") != "" {
		return false
	}

	if strings.Contains(expression, "**") || strings.Contains(expression, "//") {
		return false
	}

	if !strings.ContainsAny(expression[len(expression)-1:], "1234567890)") {
		return false
	}

	if !areParenthesisValid(expression) {
		return false
	}

	if hasSpaceWithNoOperator(expression) {
		return false
	}

	return true
}

func areParenthesisValid(expression string) bool {
	var stack Stack
	for _, char := range expression {
		if char == '(' {
			stack.Push(string(char))
		} else if char == ')' && len(stack.storage) == 0 {
			return false
		} else if char == ')' {
			_, _ = stack.Pop()
		}
	}

	return len(stack.storage) == 0
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
	postfix := convertToPostfix(input)
	return calculatePostfix(postfix)
}

func calculatePostfix(postfix string) (result int) {
	var stack Stack
	postfixString := strings.Fields(postfix)

	for _, element := range postfixString {
		if _, ok := strconv.Atoi(element); ok == nil {
			stack.Push(element)
		} else if isOperator(element) {
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()
			float1, _ := strconv.ParseFloat(num1, 64)
			float2, _ := strconv.ParseFloat(num2, 64)
			stack.Push(performOperator(float1, float2, element))
		}
	}
	topOfStack, _ := stack.Pop()
	result, _ = strconv.Atoi(topOfStack)
	return result
}

func performOperator(num1, num2 float64, operator string) (result string) {
	switch operator {
	case "*":
		result = strconv.FormatFloat(num2*num1, 'f', -1, 64)
	case "/":
		result = strconv.FormatFloat(num2/num1, 'f', -1, 64)
	case "+":
		result = strconv.FormatFloat(num2+num1, 'f', -1, 64)
	default:
		result = strconv.FormatFloat(num2-num1, 'f', -1, 64)
	}

	return result
}

func isOperator(char string) bool {
	if char == "/" || char == "+" || char == "-" || char == "*" {
		return true
	}

	return false
}

func convertToPostfix(prefix string) (postfix string) {
	var stack Stack

	for _, char := range prefix {
		if unicode.IsDigit(char) || string(char) == " " {
			postfix += string(char)
		} else if len(stack.storage) == 0 || stack.Peek() == "(" {
			stack.Push(string(char))
		} else if precedence(string(char)) > precedence(stack.Peek()) {
			stack.Push(string(char))
		} else if char == '(' {
			stack.Push(string(char))
		} else if char == ')' {
			for {
				operator, _ := stack.Pop()
				if operator == "(" {
					break
				}
				postfix += " " + operator
			}
		} else if precedence(string(char)) <= precedence(stack.Peek()) {
			for {
				if precedence(stack.Peek()) < precedence(string(char)) || stack.Peek() == ")" || len(stack.storage) == 0 {
					stack.Push(string(char))
					break
				}
				op, _ := stack.Pop()
				postfix += op
			}
		}
	}

	for len(stack.storage) != 0 {
		oper, _ := stack.Pop()
		postfix += " " + string(oper)
	}

	return postfix
}

func precedence(operator string) (precedence int) {
	operatorMap := map[string]int{"+": 0, "-": 0, "*": 1, "/": 1}
	return operatorMap[operator]
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

func (s *Stack) Push(value string) {
	s.storage = append(s.storage, value)
}

func (s *Stack) Pop() (string, error) {
	last := len(s.storage) - 1
	if last <= -1 {
		return "0", errors.New("stack is empty")
	}

	value := s.storage[last]
	s.storage = s.storage[:last]

	return value, nil
}

func (s *Stack) Peek() string {
	if len(s.storage) == 0 {
		return ""
	} else {
		return s.storage[len(s.storage)-1]
	}
}
