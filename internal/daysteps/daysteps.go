package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	errParse          = errors.New("failed to parse the string according to the required format")
	errWrongArguments = errors.New("incorrect input arguments (steps and duration have to be positive)")
)

type DaySteps struct {
	Steps                 int           // Количество шагов
	Duration              time.Duration // Длительность прогулки
	personaldata.Personal               //
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Парсинг строки и проверка корректности результата
	dataParsed := strings.Split(datastring, ",")
	if len(dataParsed) != 2 {
		return errParse
	}

	// Конвертация аргументов к требуемым типам
	steps, err := strconv.Atoi(dataParsed[0])
	if err != nil {
		return err
	}
	duration, err := time.ParseDuration(dataParsed[1])
	if err != nil {
		return err
	}

	// Проверка неотрицательности числа шагов и длительности тренировки
	if (steps <= 0) || (duration <= 0) {
		return errWrongArguments
	}

	// Сохранение полученных данных в полях структуры
	ds.Steps = steps
	ds.Duration = duration

	// Возврат отсутствия ошибки
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисление дистанции и числа калорий
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	// Возврат строки в требуемом формате
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
