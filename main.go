package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Operation(x int, y int, b string) (z int) {

	switch b {
	case "+":

		z = x + y

	case "-":

		z = x - y

	case "*":

		z = x * y
	case "/":

		z = x / y

	}
	return z

}

func main() {

	var number1 int
	var number2 int

	var roman = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

	fmt.Println("Ведите условие в формате x + y, где x и y-арабские либо римские цифры, а арифметический оператор(+,-,*,/)")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	text := strings.Split(strings.TrimSpace(str), " ")
	if len(text) < 3 {
		fmt.Println("Строка не является математической операцией")
		return

	}
	if len(text) > 3 {
		fmt.Println("Введите два операнда и один оператор(+,-,*,/)")
		return

	}
	numberAr1, err1 := strconv.Atoi(text[0])
	if err1 != nil {
		switch text[0] {
		case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":

			for i, val := range roman {

				if val == text[0] {
					number1 = i
					break

				} else {
					continue
				}

			}
		default:
			fmt.Println("Введите числа от I до X")
			return
		}

	}
	numberAr2, err2 := strconv.Atoi(text[2])
	if err2 != nil {
		switch text[2] {
		case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":

			for i, val := range roman {

				if val == text[2] {
					number2 = i
					break

				} else {
					continue
				}

			}
		default:
			fmt.Println("Введите числа от I до X")
			return
		}

	}
	if numberAr1 > 0 && numberAr1 <= 10 && numberAr2 > 0 && numberAr2 <= 10 {
		fmt.Println(Operation(numberAr1, numberAr2, text[1]))
		return
	}
	if numberAr1 < 0 || numberAr1 > 10 || numberAr2 < 0 || numberAr2 > 10 {
		fmt.Println("Введите числа от 1 до 10")
		return

	}

	if Operation(number1, number2, text[1]) <= 0 {
		fmt.Println("Результатом работы калькулятора с римскими цифрами могут быть только положительные числа")
		return
	}
	if number1 == 0 || number2 == 0 {
		fmt.Println("В условии должны быть либо только арабские символы, либо только римские")
		return
	} else {
		fmt.Println(roman[Operation(number1, number2, text[1])])
		return
	}

}
