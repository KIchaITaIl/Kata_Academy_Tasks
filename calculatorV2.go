package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RomeNumbers int

const (
	I RomeNumbers = iota + 1
	II
	III
	IV
	V
	VI
	VII
	VIII
	IX
	X
	XI
	XII
	XIII
	XIV
	XV
	XVI
	XVII
	XVIII
	XIX
	XX
	XXI
	XXII
	XXIII
	XXIV
	XXV
	XXVI
	XXVII
	XXVIII
	XXIX
	XXX
	XXXI
	XXXII
	XXXIII
	XXXIV
	XXXV
	XXXVI
	XXXVII
	XXXVIII
	XXXIX
	XL
	XLI
	XLII
	XLIII
	XLIV
	XLV
	XLVI
	XLVII
	XLVIII
	XLIX
	L
	LI
	LII
	LIII
	LIV
	LV
	LVI
	LVII
	LVIII
	LIX
	LX
	LXI
	LXII
	LXIII
	LXIV
	LXV
	LXVI
	LXVII
	LXVIII
	LXIX
	LXX
	LXXI
	LXXII
	LXXIII
	LXXIV
	LXXV
	LXXVI
	LXXVII
	LXXVIII
	LXXIX
	LXXX
	LXXXI
	LXXXII
	LXXXIII
	LXXXIV
	LXXXV
	LXXXVI
	LXXXVII
	LXXXVIII
	LXXXIX
	XC
	XCI
	XCII
	XCIII
	XCIV
	XCV
	XCVI
	XCVII
	XCVIII
	XCIX
	C
)

func main() {

	fmt.Println("Введите A и B")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	result, err := calc(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func calc(input string) (string, error) {
	calcTask := strings.Fields(input)
	if len(calcTask) != 3 {
		return "", errors.New("Формат математической операции не удовлетворяет заданию - два операнда и один оператор (+, -, /, *)")
	}

	var a, b int
	isNumber1 := isNumber(calcTask[0])
	isNumber2 := isNumber(calcTask[2])

	if isNumber1 == "Number" && isNumber2 == "Number" {
		var err error
		a, err = strconv.Atoi(calcTask[0])
		if err != nil {
			return "", err
		}
		b, err = strconv.Atoi(calcTask[2])
		if err != nil {
			return "", err
		}
	} else if isNumber1 == "Rome Number" && isNumber2 == "Rome Number" {
		var err error
		a, err = romanToInt(calcTask[0])
		if err != nil {
			return "", err
		}
		b, err = romanToInt(calcTask[2])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("Разные системы счисления")
	}

	if a < 1 || a > 10 {
		return "", errors.New("Некорректное число, Введите Число от 1 до 10")
	}
	if b < 1 || b > 10 {
		return "", errors.New("Некорректное число, Введите Число от 1 до 10")
	}

	var result int
	switch calcTask[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", errors.New("Деление на ноль")
		}
		result = a / b
	default:
		return "", errors.New("Неправильный формат операции")
	}

	if result < 1 && isNumber1 == "Rome Number" {
		return "", errors.New("В римской системе нет отрицательных чисел")
	}

	if isNumber1 == "Number" {
		return strconv.Itoa(result), nil
	}
	return intToRoman(result), nil
}

func isNumber(number string) string {
	if _, err := strconv.Atoi(number); err == nil {
		return "Number"
	}
	if _, err := romanToInt(number); err == nil {
		return "Rome Number"
	}
	return "Not Number"
}

func romanToInt(s string) (int, error) {
	romanMap := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100,
	}
	total := 0
	prevValue := 0

	for _, char := range s {
		value, exists := romanMap[char]
		if !exists {
			return 0, errors.New("Некорректное римское число")
		}
		if value > prevValue {
			total += value - 2*prevValue // Subtract twice the previous value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, nil
}

func intToRoman(num int) string {
	vals := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	romans := []string{"C", "XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

	result := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += romans[i]
		}
	}
	return result
}
