package usecases

import "clean-go/src/domain/models"

type Divide interface {
	Calculate(data models.Divide) float32
}
