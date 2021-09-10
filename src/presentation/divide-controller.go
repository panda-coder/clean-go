package presentation

import (
	"encoding/json"

	"clean-go/src/domain/models"
	"clean-go/src/domain/usecases"
)

type DivideController struct {
	Divide usecases.Divide
}

func (s *DivideController) Handle(req HttpRequest) HttpResponse {
	var calculateData models.Divide
	err := json.NewDecoder(req.Body).Decode(&calculateData)
	if err != nil {
		return badRequest("Invalid Payload")
	}
	response := s.Divide.Calculate(calculateData)
	return ok(map[string]float32{"value": response})
}
