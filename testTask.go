package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите пример:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")
		s := make([]string, 2) //слайс для операндов текстового типа

		if strings.Count(text, "+") > 1 || strings.Count(text, "-") > 1 || strings.Count(text, "*") > 1 || strings.Count(text, "/") > 1 || strings.Count(text, "*") >= 1 && strings.Count(text, "/") >= 1 || strings.Count(text, "*") >= 1 && strings.Count(text, "+") >= 1 || strings.Count(text, "*") >= 1 && strings.Count(text, "-") >= 1 || strings.Count(text, "/") >= 1 && strings.Count(text, "+") >= 1 || strings.Count(text, "+") >= 1 && strings.Count(text, "-") >= 1 || strings.Count(text, "/") >= 1 && strings.Count(text, "-") >= 1 {
			fmt.Println("Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			break
		} else if strings.Count(text, "+") == 1 {
			s = strings.Split(text, "+")
			firstNumber, secondNumber, hasRoman := conversion(s) //функция, преобразующая введеный операнд типа string в int

			answer := firstNumber + secondNumber
			answerRome(answer, hasRoman) //функция, обрабатывающая ответ

		} else if strings.Count(text, "-") == 1 {
			s = strings.Split(text, "-")
			firstNumber, secondNumber, hasRoman := conversion(s)

			if hasRoman == true && firstNumber < secondNumber {
				fmt.Println("ошибкa, так как в римской системе нет отрицательных чисел.")
				os.Exit(0)
			} else if hasRoman == true && firstNumber == secondNumber {
				fmt.Println("ошибкa, так как в римской системе отсутствует ноль.")
				os.Exit(0)
			} else {
				answer := firstNumber - secondNumber
				answerRome(answer, hasRoman)
			}
		} else if strings.Count(text, "*") == 1 {
			s = strings.Split(text, "*")
			firstNumber, secondNumber, hasRoman := conversion(s)

			answer := firstNumber * secondNumber
			answerRome(answer, hasRoman)

		} else if strings.Count(text, "/") == 1 {
			s = strings.Split(text, "/")
			firstNumber, secondNumber, hasRoman := conversion(s)

			answer := firstNumber / secondNumber
			answerRome(answer, hasRoman)

		} else {
			fmt.Println("Ошибка, так как строка не является математической операцией.")
			break
		}
	}
}
func conversion(s []string) (int, int, bool) { //функция, преобразующая введеный операнд из строкового типа в числовой
	hasRoman := true //переменные, по которым отслеживается из какой системы счисления операнд
	hasRoman2 := true

	switch s[0] {
	case "I":
		s[0] = "1"
	case "II":
		s[0] = "2"
	case "III":
		s[0] = "3"
	case "IV":
		s[0] = "4"
	case "V":
		s[0] = "5"
	case "VI":
		s[0] = "6"
	case "VII":
		s[0] = "7"
	case "VIII":
		s[0] = "8"
	case "IX":
		s[0] = "9"
	case "X":
		s[0] = "10"
	default:
		if s[0] == "1" || s[0] == "2" || s[0] == "3" || s[0] == "4" || s[0] == "5" || s[0] == "6" || s[0] == "7" || s[0] == "8" || s[0] == "9" || s[0] == "10" {
			hasRoman = false
		} else {
			fmt.Println("ошибкa, так как калькулятор принимает на вход только целые римские числа от I до X и арабские от 1 до 10 включительно.")
			os.Exit(0)
		}
	}

	switch s[1] {
	case "I":
		s[1] = "1"
	case "II":
		s[1] = "2"
	case "III":
		s[1] = "3"
	case "IV":
		s[1] = "4"
	case "V":
		s[1] = "5"
	case "VI":
		s[1] = "6"
	case "VII":
		s[1] = "7"
	case "VIII":
		s[1] = "8"
	case "IX":
		s[1] = "9"
	case "X":
		s[1] = "10"
	default:
		if s[1] == "1" || s[1] == "2" || s[1] == "3" || s[1] == "4" || s[1] == "5" || s[1] == "6" || s[1] == "7" || s[1] == "8" || s[1] == "9" || s[1] == "10" {
			hasRoman2 = false
		} else {
			fmt.Println("ошибкa, так как калькулятор принимает на вход только целые римские числа от I до X и арабские от 1 до 10 включительно.")
			os.Exit(0)
		}
	}

	for hasRoman != hasRoman2 {
		fmt.Println("ошибкa, так как одновременно используются разные системы счисления.")
		os.Exit(0)
	}

	firstNumber, _ := strconv.Atoi(s[0])
	secondNumber, _ := strconv.Atoi(s[1])
	return firstNumber, secondNumber, hasRoman

}

func answerRome(answer int, hasRoman bool) {
	//функция определяет систему счисления и по необходимости переводит в нужную
	if hasRoman == true {
		tens := answer / 10
		units := answer % 10
		a := make([]string, 2) //слайс, в котором первый элемент - десятки, второй - единицы, записанные римскими цифрами

		switch tens {
		case 10:
			a[0] = "C"
		case 9:
			a[0] = "XC"
		case 8:
			a[0] = "LXXX"
		case 7:
			a[0] = "LXX"
		case 6:
			a[0] = "LX"
		case 5:
			a[0] = "L"
		case 4:
			a[0] = "XL"
		case 3:
			a[0] = "XXX"
		case 2:
			a[0] = "XX"
		case 1:
			a[0] = "X"
		case 0:
			a[0] = ""
		}

		switch units {
		case 9:
			a[1] = "IX"
		case 8:
			a[1] = "VIII"
		case 7:
			a[1] = "VII"
		case 6:
			a[1] = "VI"
		case 5:
			a[1] = "V"
		case 4:
			a[1] = "IV"
		case 3:
			a[1] = "III"
		case 2:
			a[1] = "II"
		case 1:
			a[1] = "I"
		case 0:
			a[1] = ""
			for a[0] == "" {
				fmt.Println("ошибкa, так как в римской системе отсутствует ноль.")
				os.Exit(0)
			}
		}
		fmt.Println(strings.Join(a, ""))
	} else {
		fmt.Println(answer)
	}
}
