package usecases

import "clean-go/src/domain/models"

type DivideAction struct{}

func (action *DivideAction) Calculate(data models.Divide) float32 {
	return data.X / data.Y
}
