package app

import (
	"encoding/json"
	"net/http"

	"clean-go/src/presentation"
)

func adaptRoute(controller presentation.Controller) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		request := presentation.HttpRequest{Body: r.Body}
		response := controller.Handle(request)

		//vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)

		responseData, _ := json.Marshal(response.Body)

		w.Write(responseData)

		// an example API handler
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		// fmt.Fprintf(w, "Category: %v\n", vars["category"])
	}
}
