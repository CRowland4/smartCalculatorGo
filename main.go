package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		input := readLine()
		if input == "/exit" {
			fmt.Println("Bye!")
			return
		}

		if num1, num2, isValid := interpretInput(input); isValid {
			fmt.Println(num1 + num2)
		}
	}

}

func interpretInput(input string) (num1, num2 int, isValid bool) {
	var snum1, snum2 string
	fmt.Sscanf(input, "%s %s", &snum1, &snum2)
	if (strings.TrimSpace(snum1) == "") && (strings.TrimSpace(snum2) == "") {
		return 0, 0, false
	}

	if strings.TrimSpace(snum1) == "" {
		num1 = 0
		num2, _ = strconv.Atoi(snum2)
	} else if strings.TrimSpace(snum2) == "" {
		num2 = 0
		num1, _ = strconv.Atoi(snum1)
	} else {
		num1, _ = strconv.Atoi(snum1)
		num2, _ = strconv.Atoi(snum2)
	}

	return num1, num2, true
}

func readLine() (line string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
