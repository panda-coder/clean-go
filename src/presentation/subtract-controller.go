package presentation

import (
	"encoding/json"

	"clean-go/src/domain/models"
	"clean-go/src/domain/usecases"
)

type SubtractController struct {
	Subtract usecases.Subtract
}

func (s *SubtractController) Handle(req HttpRequest) HttpResponse {
	var calculateData models.Subtract
	err := json.NewDecoder(req.Body).Decode(&calculateData)
	if err != nil {
		return badRequest("Invalid Payload")
	}
	response := s.Subtract.Calculate(calculateData)
	return ok(map[string]float32{"value": response})
}
