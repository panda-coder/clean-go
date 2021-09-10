package presentation

import (
	"encoding/json"

	"clean-go/src/domain/models"
	"clean-go/src/domain/usecases"
)

type MultiplyController struct {
	Multiply usecases.Multiply
}

func (s *MultiplyController) Handle(req HttpRequest) HttpResponse {
	var calculateData models.Multiply
	err := json.NewDecoder(req.Body).Decode(&calculateData)
	if err != nil {
		return badRequest("Invalid Payload")
	}
	response := s.Multiply.Calculate(calculateData)
	return ok(map[string]float32{"value": response})
}
