package app

import (
	"clean-go/src/presentation"
	"clean-go/src/usecases"
)

func makeSumController() *presentation.SumController {
	sum := &usecases.SumAction{}
	sumController := &presentation.SumController{Sum: sum}
	return sumController
}

func makeSubtractController() *presentation.SubtractController {
	subtract := &usecases.SubtractAction{}
	subtractController := &presentation.SubtractController{Subtract: subtract}
	return subtractController
}

func makeMultiplyController() *presentation.MultiplyController {
	multiply := &usecases.MultiplyAction{}
	multiplyController := &presentation.MultiplyController{Multiply: multiply}
	return multiplyController
}

func makeDivideController() *presentation.DivideController {
	divide := &usecases.DivideAction{}
	divideController := &presentation.DivideController{Divide: divide}
	return divideController
}
