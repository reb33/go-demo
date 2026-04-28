package main
import (
	"errors"
	"fmt"
	"math"

)

const IMT_POWER = 2

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover", r)
		}
	}()

	println("__ Калькулятор индекса массы тела")
	for {
		userWeight, userHeight := getUserInput()
		imt, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			// fmt.Println(err, "- Некорректные параметры для расчета")
			// continue
			panic("Некорректные параметры для расчета")
		}
		outputResult(imt)
		if !checkRepeatCalculation() {
			break
		}
	}

}

func outputResult(_imt float64) {
	result := fmt.Sprintf("IMT: %.2f \n", _imt)
	print(result)
	switch {
	case _imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case _imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case _imt < 25:
		fmt.Println("У вас нормальная масса тела")
	case _imt < 30:
		fmt.Println("У вас избыточная масса тела")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(weight float64, height float64) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("PARAMS_ERROR")
	}
	imt := float64(weight) / math.Pow(height/100, IMT_POWER)
	return imt, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	print("Введите свой рост в см: ")
	fmt.Scan(&userHeight)
	print("Введите свой вес: ")
	fmt.Scan(&userWeight)

	return userWeight, userHeight
}

func checkRepeatCalculation() bool {
	fmt.Println("\nХотите повторить расчет? (y/n)")
	var input string
	fmt.Scan(&input)
	if input == "y" || input == "Y" {
		return true
	}
	return false
}
