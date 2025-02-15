package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings" 


func arabicToRoman(num int) string { // преобразование арабских чисел в римские
	if num < 0 {
		panic("римское число должно быть больше 0 ")
	}

	 values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	 symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}
	return roman
}

func romanToArabic(roman string) int { // преобразование римского числа в арабское
	romanToArabicMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}

	total := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanToArabicMap[roman[i]]
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total
}

func isRomanNumber(s string) bool { // является ли строка римским числом ?
	for _, char := range s {
		if !strings.ContainsRune("IVXLCDM", char) {
			return false
		}
	}
	return true
}

func isArabicNumber(s string) bool { // является ли строка арабским числом ?
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение с пробелами:") // Наличие пробелов обязательно
	input, _ := reader.ReadString('\n') // Считываем всю строку
	input = strings.TrimSpace(input)    // Удаляем лишние пробелы в начале и конце

	parts := strings.Fields(input) // Разделяем строку на части Разделение по пробелам
	if len(parts) != 3 {
		panic(fmt.Sprintf("Неверный формат ввода. Ожидается 3части 'число оператор число', получено: %d частей", len(parts))) // вечная слава логированию
	}

	isArabic := isArabicNumber(parts[0]) && isArabicNumber(parts[2]) // Определяем тип чисел
	isRoman := isRomanNumber(parts[0]) && isRomanNumber(parts[2])

	if !isArabic && !isRoman { // если вводим римские и арабские символы вместе
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	var a, b int

	if isArabic {
		// Парсер арабских чисел
		a, _ = strconv.Atoi(parts[0])
		b, _ = strconv.Atoi(parts[2])
	} else if isRoman {
		// Парсер римских чисел
		a = romanToArabic(parts[0])
		b = romanToArabic(parts[2])
	}

	// Проверка диапазона чисел
    if a < 1 {
        panic("Числа должны быть от 1 до 10") // P.S. можно короче if a < 1 || a > 10 || b < 1 || b > 10 { panic("") }
    }
	if a > 10 {
        panic("Числа должны быть от 1 до 10")
    }
	if b < 1 {
        panic("Числа должны быть от 1 до 10")
    }
	if b > 10 {
        panic("Числа должны быть от 1 до 10")
    }

	// Выполнение операции
	var result int
	switch parts[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("не дели на ноль")
		}
		result = a / b
	default:
		panic("Незаданая арифметическая операция.")
	}

	// Выводим результат
	if isArabic {
		fmt.Println("Результат:", result)
	} else if isRoman {
		if result <= 0 {
			panic("результат работы с римскими числами должен быть больше или равен 1")
		}
		fmt.Println("Результат:", arabicToRoman(result))
	}
}
