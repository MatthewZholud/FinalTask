package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatthewZholud/FinalTask/TimeTracker/DbService"
	"github.com/MatthewZholud/FinalTask/TimeTracker/Entities"
	"github.com/gorilla/mux"
)

func PostTimeFrames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	timeframes := Entities.TimeFrames{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	timeframes.TaskID = r.Form.Get("task_id")
	timeframes.From = r.Form.Get("time_from")
	timeframes.To = r.Form.Get("time_to")

	err = DbService.Conn.PostTimeFrames(&timeframes)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(timeframes)
}

func DeleteTimeframes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := (mux.Vars(r))["id"]
	err := DbService.Conn.DeleteTimeframes(id)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
