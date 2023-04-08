package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arabicToRoman(num int) string {
	roman := ""
	for num > 0 {
		switch {
		case num >= 100:
			roman += "C"
			num -= 100
		case num >= 90:
			roman += "XC"
			num -= 90
		case num >= 50:
			roman += "L"
			num -= 50
		case num >= 40:
			roman += "XL"
			num -= 40
		case num >= 10:
			roman += "X"
			num -= 10
		case num >= 9:
			roman += "IX"
			num -= 9
		case num >= 5:
			roman += "V"
			num -= 5
		case num >= 4:
			roman += "IV"
			num -= 4
		default:
			roman += "I"
			num -= 1
		}
	}
	return roman
}

func calculateArabic(a, b int, operator string) (int, error) {
	var result int
	if a > 0 && a <= 10 && b > 0 && b <= 10 {
		switch operator {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		default:
			fmt.Println("Unknown operator")
			os.Exit(1)
			return 0, errors.New("Unknown operator")
		}
	} else {
		fmt.Println("Numbers must be one type, integer, more or equal 1('I') and less or equal 10('X')")
		os.Exit(1)
	}
	return result, nil
}

func calculateRoman(a, b, operator string) (string, error) {
	var x, y, result int
	switch a {
	case "I", "i":
		x = 1
	case "II", "ii":
		x = 2
	case "III", "iii":
		x = 3
	case "IV", "iv":
		x = 4
	case "V", "v":
		x = 5
	case "VI", "vi":
		x = 6
	case "VII", "vii":
		x = 7
	case "VIII", "viii":
		x = 8
	case "IX", "ix":
		x = 9
	case "X", "x":
		x = 10
	default:
		fmt.Println("Numbers must be one type, integer, and more or equal 1('I') and less or equal 10('X')")
		os.Exit(1)
		return "", errors.New("Unknown roman symbol")
	}
	switch b {
	case "I", "i":
		y = 1
	case "II", "ii":
		y = 2
	case "III", "iii":
		y = 3
	case "IV", "iv":
		y = 4
	case "V", "v":
		y = 5
	case "VI", "vi":
		y = 6
	case "VII", "vii":
		y = 7
	case "VIII", "viii":
		y = 8
	case "IX", "ix":
		y = 9
	case "X", "x":
		y = 10
	default:
		fmt.Println("Numbers must be one type, integer, more or equal 1('I') and less or equal 10('X')")
		os.Exit(1)
		return "", errors.New("")
	}
	if x > 0 && x <= 10 && y > 0 && y <= 10 {
		switch operator {
		case "+":
			result = x + y
		case "-":
			result = x - y
		case "*":
			result = x * y
		case "/":
			result = x / y
		default:
			fmt.Println("Unknown operator")
			os.Exit(1)
			return "", errors.New("Unknown operator")
		}
		if result <= 0 {
			fmt.Println("Roman number cannot less than 1")
			os.Exit(1)
		}
		return arabicToRoman(result), nil
	} else {
		return "", errors.New("")
	}
}
func main() {
	var a, b string
	var operator string
	var err error
	var solution int
	var solutionString string

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your task: ")
	if !reader.Scan() {
		fmt.Println("Error", reader.Err())
		return
	}
	input := reader.Text()

	elements := strings.Fields(input)
	if len(elements) != 3 {
		fmt.Println("You must enter only three elements")
		return
	}

	a = elements[0]
	operator = elements[1]
	b = elements[2]

	if aInt, err := strconv.Atoi(a); err != nil {
		solutionString, err = calculateRoman(a, b, operator)
		fmt.Println(solutionString)
		return
	} else if bInt, err := strconv.Atoi(b); err != nil {
		solution, err = calculateArabic(aInt, bInt, operator)
	} else {
		solution, err = calculateArabic(aInt, bInt, operator)
	}
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("Solution is %d\n", solution)

	}
}
