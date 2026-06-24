package personaldata

import (
	"fmt"
)

type Personal struct {
	Name   string  // Имя пользователя
	Weight float64 // Вес пользователя
	Height float64 // Высота пользователя
}

func (p Personal) Print() {
	// Вывод информации о пользователе
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
