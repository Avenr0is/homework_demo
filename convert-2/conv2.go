package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("===КОНВЕРТЕР ВАЛЮТ===\n")

	banner := help()

	fmt.Println(banner)
    
    var choice int
	fmt.Scan(&choice)

	amount, err := userInput(choice)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	
	result := calculate(choice, amount)
	fmt.Printf("Результат конвертации: %.2f\n", result)
}

func help() (helpList string) {

	helpList = `Выбирите опцию 
	1. USD -> RUB
	2. USD -> EUR
	3. EUR -> USD
	4. EUR -> RUB
	5. RUB -> USD
	6. RUB -> EUR
	
Просто укажите цифру:`


	return helpList
}


func userInput(choice int) (float64, error) {
	
	if choice < 1 || choice > 6 {

		return 0, errors.New("Ошибка ввода: неверный номер опции")
	}

	var amount float64
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)

    
	if(amount == 0){

		return 0, errors.New("Вы ввели не число")
	}
	
	return amount, nil

}

func calculate(choice int, amount float64)float64{

 	const usdToRub = 78.73
	const usdToEur = 0.8649
	const eurToUsd = 1.16
	const eurToRub = 91.0
	const rubToUsd = 0.012702
	const rubToEur = 0.10989

	
	switch choice {
	case 1:
		return amount * usdToRub
	case 2:
		return amount * usdToEur
	case 3:
		return amount * eurToUsd
	case 4:
		return amount * eurToRub
	case 5:
		return amount * rubToUsd
	case 6:
		return amount * rubToEur
	default:
		return 0
	}
  

}
