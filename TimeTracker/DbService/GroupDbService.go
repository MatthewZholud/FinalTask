package DbService

import "github.com/MatthewZholud/FinalTask/TimeTracker/Entities"

func (conn *DbStruct) GetGroupsDb() ([]Entities.Groups, error) {
	rows, err := conn.db.Query("SELECT * from groups")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := []Entities.Groups{}

	for rows.Next() {
		group := Entities.Groups{}

		if err := rows.Scan(&group.ID, &group.Title); err != nil {
			return nil, err
		}
		group.Tasks, err = conn.GetTasksDb(group.ID, true)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (conn *DbStruct) PostGroup(group *Entities.Groups) (string, error) {
	err := conn.db.QueryRow("INSERT INTO groups(title) VALUES ($1) RETURNING group_id", group.Title).Scan(&group.ID)
	return group.ID, err
}

func (conn *DbStruct) PutGroup(group *Entities.Groups) error {
	_, err := conn.db.Exec("update groups set title = $1 where group_id = $2;", group.Title, group.ID)
	return err
}

func (conn *DbStruct) DeleteGroup(id string) error {
	_, err := conn.db.Exec("DELETE FROM groups WHERE group_id = $1", id)
	return err
}
