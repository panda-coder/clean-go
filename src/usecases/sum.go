package usecases

import "clean-go/src/domain/models"

type SumAction struct {}

func (action *SumAction) Calculate(data models.Sum) float32 {
	return data.X + data.Y
}