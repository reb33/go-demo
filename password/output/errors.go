package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errValue, ok := value.(error)
	if ok {
		color.Red(errValue.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")
	fmt.Println(value)
}
