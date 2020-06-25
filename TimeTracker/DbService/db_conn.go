package DbService

import (
	"database/sql"
	"fmt"
	"github.com/MatthewZholud/FinalTask/TimeTracker/Entities"
	"os"
)

type ServeDb interface {
	GetGroupsDb() ([]Entities.Groups, error)
	PostGroup(group *Entities.Groups) (string, error)
	PutGroup(group *Entities.Groups) error
	DeleteGroup(id string) error

	GetTasksDb(GrID string, ch bool) ([]Entities.Tasks, error)
	PostTask(task *Entities.Tasks) (string, error)
	PutTask(task *Entities.Tasks) error
	DeleteTask(id string) error

	GetTimeFramesDb(task_id string) ([]Entities.TimeFrames, error)
	PostTimeFrames(timeframe *Entities.TimeFrames) error
	DeleteTimeframes(id string) error
}

type DbStruct struct {
	db *sql.DB
}


func Db_Conn() {

	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"postgresdb", "5432", "postgres", "mypassword", "time_tracker")

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	InitDbService(&DbStruct{db: db})
}

var Conn ServeDb

func InitDbService(s ServeDb) {
	Conn = s
}
