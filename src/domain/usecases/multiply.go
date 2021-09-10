package usecases

import "clean-go/src/domain/models"

type Multiply interface {
	Calculate(data models.Multiply) float32
}
