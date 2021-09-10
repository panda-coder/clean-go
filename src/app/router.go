package app

import (
	"github.com/gorilla/mux"
)

// função principal
func setup_router() *mux.Router {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/sum", adaptRoute(makeSumController()))
	v1.HandleFunc("/subtract", adaptRoute(makeSubtractController()))
	v1.HandleFunc("/multiply", adaptRoute(makeMultiplyController()))
	v1.HandleFunc("/divide", adaptRoute(makeDivideController()))

	return router
}
