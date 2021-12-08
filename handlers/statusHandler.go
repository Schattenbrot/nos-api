package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Schattenbrot/nos-api/config"
)

//AppStatus is the type for the applications status.
type appStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// statusHandler is the handler for the appstatus.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := appStatus{
		Status:      "Available",
		Environment: config.Cfg.Env,
		Version:     config.App.Version,
	}

	js, err := json.Marshal(currentStatus)
	if err != nil {
		config.App.Logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
