package spentenergy

import (
	"errors"
	"time"
)

var errWrongArguments = errors.New("incorrect input arguments (steps, duration, weight and height have to be positive)")

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка аргументов на положительность
	if (steps <= 0) || (duration <= 0) || (weight <= 0) || (height <= 0) {
		return 0, errWrongArguments
	}
	// Вычисление и возврат числа калорий
	return walkingCaloriesCoefficient * (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка аргументов на положительность
	if (steps <= 0) || (duration <= 0) || (weight <= 0) || (height <= 0) {
		return 0, errWrongArguments
	}
	// Вычисление и возврат числа калорий
	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка, что duration > 0
	if duration <= 0 {
		return 0
	}
	// Вычисление и возврат средней скорости
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// Вычисление и возврат дистанции
	return stepLengthCoefficient * height * float64(steps) / mInKm
}
