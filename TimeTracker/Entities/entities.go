package Entities

type Groups struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tasks []Tasks
}

type Tasks struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Group      string `json:"group"`
	TimeFrames  []TimeFrames
}

type TimeFrames struct {
	TaskID string  `json:"task_id"`
	From string `json:"from"`
	To   string `json:"to"`
}
