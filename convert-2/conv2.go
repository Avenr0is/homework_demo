package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("===КОНВЕРТЕР ВАЛЮТ===\n")

	for {
		fromCurrency, err := getCurrency("Введите исходную валюту (USD, EUR, RUB): ")
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		amount, err := getAmount()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		toCurrency, err := getCurrency("Введите целевую валюту (USD, EUR, RUB): ")
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		result, err := calculate(fromCurrency, toCurrency, amount)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("Результат: %.2f %s -> %.2f %s\n\n", amount, fromCurrency, result, toCurrency)

		fmt.Print("Хотите выполнить еще одну конвертацию? (да/нет): ")
		var answer string
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		if answer != "да" && answer != "yes" && answer != "y" {
			fmt.Println("До свидания!")
			break
		}
		fmt.Println()
	}
}

func getCurrency(prompt string) (string, error) {
	fmt.Print(prompt)
	var currency string
	fmt.Scan(&currency)
	currency = strings.ToUpper(strings.TrimSpace(currency))

	if currency == "USD" || currency == "EUR" || currency == "RUB" {
		return currency, nil
	}
	return "", errors.New("валюта должна быть USD, EUR или RUB")
}

func getAmount() (float64, error) {
	fmt.Print("Введите сумму для конвертации: ")
	var amount float64
	_, err := fmt.Scan(&amount)

	if err != nil {
		return 0, errors.New("введите корректное число")
	}

	if amount <= 0 {
		return 0, errors.New("сумма должна быть больше нуля")
	}

	return amount, nil
}

func calculate(from, to string, amount float64) (float64, error) {
	const usdToRub = 78.73
	const eurToRub = 91.0

	if from == to {
		return 0, errors.New("исходная и целевая валюта не могут совпадать")
	}

	var rubAmount float64

	switch from {
	case "USD":
		rubAmount = amount * usdToRub
	case "EUR":
		rubAmount = amount * eurToRub
	case "RUB":
		rubAmount = amount
	default:
		return 0, errors.New("неизвестная исходная валюта")
	}

	switch to {
	case "USD":
		return rubAmount / usdToRub, nil
	case "EUR":
		return rubAmount / eurToRub, nil
	case "RUB":
		return rubAmount, nil
	default:
		return 0, errors.New("неизвестная целевая валюта")
	}
}