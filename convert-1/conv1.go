package main

import (
	"fmt"
)

func main() {

	const rateUSDtoEur = 0.8723
	const rateUSDtoRub = 81.3
	var rateEurtoRub = rateUSDtoRub / rateUSDtoEur

	fmt.Println("Отношение EUR к RUB = ", rateEurtoRub)

}
