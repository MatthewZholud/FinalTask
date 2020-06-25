package DbService

import "github.com/MatthewZholud/FinalTask/TimeTracker/Entities"

func (conn *DbStruct) GetTimeFramesDb(TasID string) ([]Entities.TimeFrames, error) {
	rows, err := conn.db.Query("SELECT * from timeframes where task_id = $1", TasID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	timeframes := []Entities.TimeFrames{}
	for rows.Next() {
		timeframe := Entities.TimeFrames{}
		if err := rows.Scan(&timeframe.TaskID, &timeframe.From, &timeframe.To); err != nil {
			return nil, err
		}
		timeframes = append(timeframes, timeframe)
	}
	return timeframes, nil
}

func (conn *DbStruct) PostTimeFrames(timeframe *Entities.TimeFrames) error {
	_, err := conn.db.Query("INSERT INTO timeframes(time_from, time_to, task_id) VALUES ($1, $2, $3)", timeframe.From, timeframe.To, timeframe.TaskID)
	return err
}

func (conn *DbStruct) DeleteTimeframes(id string) error {
	_, err := conn.db.Exec("DELETE FROM timeframes WHERE task_id = $1", id)
	return err
}
