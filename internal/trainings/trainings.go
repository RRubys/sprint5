package trainings

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
	errParse                = errors.New("failed to parse the string according to the required format")
	errWrongArguments       = errors.New("incorrect input arguments (steps and duration have to be positive)")
	errUnrecognizedActivity = errors.New("unrecognized type of the training (neither \"Бег\" nor \"Ходьба\")")
)

type Training struct {
	Steps                 int           // Количество шагов, проделанных за тренировку
	TrainingType          string        // Тип тренировки (бег или ходьба)
	Duration              time.Duration // Длительность тренировки
	personaldata.Personal               // Информация о пользователе
}

func (t *Training) Parse(datastring string) (err error) {
	// Парсинг строки и проверка корректности результата
	dataParsed := strings.Split(datastring, ",")
	if len(dataParsed) != 3 {
		return errParse
	}

	// Конвертация аргументов к требуемым типам
	steps, err := strconv.Atoi(dataParsed[0])
	if err != nil {
		return err
	}
	duration, err := time.ParseDuration(dataParsed[2])
	if err != nil {
		return err
	}

	// Проверка неотрицательности числа шагов и длительности тренировки
	if (steps <= 0) || (duration <= 0) {
		return errWrongArguments
	}

	// Сохранение полученных данных в полях структуры
	t.Steps = steps
	t.TrainingType = dataParsed[1]
	t.Duration = duration

	// Возврат отсутствия ошибки
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Проверка вида тренировки, а также вычисление дистанции, средней скорости и калорий
	var (
		activityCalories float64
		err              error
	)
	switch t.TrainingType {
	case "Ходьба":
		activityCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		activityCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errUnrecognizedActivity
	}
	if err != nil {
		return "", err
	}
	activityDistance := spentenergy.Distance(t.Steps, t.Height)
	activityMeanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// Возврат строки в требуемом формате
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), activityDistance, activityMeanSpeed, activityCalories), nil
}
