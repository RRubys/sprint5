package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	var (
		err     error
		message string
	)
	// Проход по содержимому слайса с данными в виде строк
	for _, elem := range dataset {

		// Парсинг данных из текущей строки
		err = dp.Parse(elem)
		if err != nil {
			log.Println(err)
			continue
		}

		// Вывод сообщения для текущих данных
		message, err = dp.ActionInfo()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(message)
		}
	}
}
