package usecases

import "clean-go/src/domain/models"

type Sum interface {
    Calculate(data models.Sum) float32
}