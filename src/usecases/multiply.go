package usecases

import "clean-go/src/domain/models"

type MultiplyAction struct{}

func (action *MultiplyAction) Calculate(data models.Multiply) float32 {
	return data.X * data.Y
}
