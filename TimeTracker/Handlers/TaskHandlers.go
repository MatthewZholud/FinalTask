package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MatthewZholud/FinalTask/TimeTracker/DbService"
	"github.com/MatthewZholud/FinalTask/TimeTracker/Entities"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := DbService.Conn.GetTasksDb("0", false)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task := Entities.Tasks{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task.Title = r.Form.Get("task_title")
	task.Group = r.Form.Get("group_id")

	task.ID, err = DbService.Conn.PostTask(&task)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task := Entities.Tasks{}
	task.ID = (mux.Vars(r))["id"]

	task.Title = r.Form.Get("task_title")
	task.Group = r.Form.Get("group_id")
	//task.Title = "Cinema"
	//task.Group = "3"

	err := DbService.Conn.PutTask(&task)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := (mux.Vars(r))["id"]
	err := DbService.Conn.DeleteTask(vars)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
