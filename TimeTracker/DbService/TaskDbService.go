package DbService

import "github.com/MatthewZholud/FinalTask/TimeTracker/Entities"

func (conn *DbStruct) GetTasksDb(GrID string, ch bool) ([]Entities.Tasks, error) {
	if ch { //if from Groups
		rows, err := conn.db.Query("SELECT * from tasks where group_id = $1", GrID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		tasks := []Entities.Tasks{}

		for rows.Next() {
			task := Entities.Tasks{}

			if err := rows.Scan(&task.ID, &task.Title, &task.Group); err != nil {
				return nil, err
			}

			task.TimeFrames, err = conn.GetTimeFramesDb(task.ID)
			if err != nil {
				return nil, err
			}

			tasks = append(tasks, task)
		}
		return tasks, nil
	} else { //if by itself
		rows, err := conn.db.Query("SELECT * from tasks")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		tasks := []Entities.Tasks{}

		for rows.Next() {
			task := Entities.Tasks{}

			if err := rows.Scan(&task.ID, &task.Title, &task.Group); err != nil {
				return nil, err
			}

			task.TimeFrames, err = conn.GetTimeFramesDb(task.ID)
			if err != nil {
				return nil, err
			}

			tasks = append(tasks, task)
		}
		return tasks, nil
	}
}

func (conn *DbStruct) PostTask(task *Entities.Tasks) (string, error) {
	err := conn.db.QueryRow("INSERT INTO tasks(title, group_id) VALUES ($1, $2)", task.Title, task.Group).Scan(&task.ID)
	return task.ID, err
}

func (conn *DbStruct) PutTask(task *Entities.Tasks) error {
	_, err := conn.db.Exec("update tasks set title = $1, group_id = $2 where group_id = $3;", task.Title, task.Group, task.ID)
	return err
}

func (conn *DbStruct) DeleteTask(id string) error {
	_, err := conn.db.Exec("DELETE FROM tasks WHERE task_id = $1", id)
	return err
}
