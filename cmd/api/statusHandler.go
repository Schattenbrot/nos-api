package main

import (
	"encoding/json"
	"net/http"
)

//AppStatus is the type for the applications status.
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// statusHandler is the handler for the appstatus.
func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.Marshal(currentStatus)
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
