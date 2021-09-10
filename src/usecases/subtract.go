package usecases

import "clean-go/src/domain/models"

type SubtractAction struct{}

func (action *SubtractAction) Calculate(data models.Subtract) float32 {
	return data.X - data.Y
}
