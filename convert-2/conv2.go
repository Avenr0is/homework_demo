package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("===КОНВЕРТЕР ВАЛЮТ===\n")

	for {
		fromCurrency := inputAndValidateCurrency("исходную")
		if fromCurrency == "" {
			continue
		}

		amount := inputAndValidateAmount()
		if amount == 0 {
			continue
		}

		toCurrency := inputAndValidateCurrency("целевую")
		if toCurrency == "" {
			continue
		}

		result, err := calculate(fromCurrency, toCurrency, amount)
		if err != nil {
			fmt.Printf("Ошибка: %v\n\n", err)
			continue
		}

		fmt.Printf("\nРезультат: %.2f %s -> %.2f %s\n\n", amount, fromCurrency, result, toCurrency)

		fmt.Print("Продолжить? (да/нет): ")
		var answer string
		fmt.Scan(&answer)
		if strings.ToLower(answer) != "да" {
			fmt.Println("До свидания!")
			break
		}
		fmt.Println()
	}
}

func inputAndValidateCurrency(currencyType string) string {
	for {
		fmt.Printf("Введите %s валюту (доступны: USD, EUR, RUB): ", currencyType)
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(strings.TrimSpace(currency))

		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
		}
		fmt.Println("Ошибка: валюта должна быть USD, EUR или RUB. Попробуйте снова.")
	}
}

func inputAndValidateAmount() float64 {
	for {
		fmt.Print("Введите сумму для конвертации (число > 0): ")
		var amount float64
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка: введите корректное число. Попробуйте снова.")
			var discard string
			fmt.Scan(&discard)
			continue
		}

		if amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть больше нуля. Попробуйте снова.")
			continue
		}

		return amount
	}
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