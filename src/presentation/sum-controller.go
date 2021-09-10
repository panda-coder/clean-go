package presentation

import (
	"encoding/json"

	"clean-go/src/domain/models"
	"clean-go/src/domain/usecases"
)

type SumController struct {
	Sum usecases.Sum
}

func (s *SumController) Handle(req HttpRequest) HttpResponse {
	var calculateData models.Sum
	err := json.NewDecoder(req.Body).Decode(&calculateData)
	if err != nil {
		return badRequest("Invalid Payload")
	}
	response := s.Sum.Calculate(calculateData)
	return ok(map[string]float32{"value": response})
}
