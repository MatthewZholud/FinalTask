package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MatthewZholud/FinalTask/TimeTracker/DbService"
	"github.com/MatthewZholud/FinalTask/TimeTracker/Handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/groups", Handlers.GetGroups).Methods(http.MethodGet)
	r.HandleFunc("/groups/", Handlers.PostGroup).Methods(http.MethodPost)
	r.HandleFunc("/groups/{id}", Handlers.PutGroup).Methods(http.MethodPut)
	r.HandleFunc("/groups/{id}", Handlers.DeleteGroup).Methods(http.MethodDelete)

	r.HandleFunc("/tasks", Handlers.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/", Handlers.PostTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{id}", Handlers.PutTask).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", Handlers.DeleteTask).Methods(http.MethodDelete)

	r.HandleFunc("/timeframes/", Handlers.PostTimeFrames).Methods(http.MethodPost)
	r.HandleFunc("/timeframes/{id}", Handlers.DeleteTimeframes).Methods(http.MethodDelete)
	return r
}

func main() {
	fmt.Println("Starting server...")
	DbService.Db_Conn()

	r := newRouter()

	PORT := ":8000"
	fmt.Println("Serving on port: ", PORT)

	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
