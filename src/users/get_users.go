package users

import (
	"log"
)

func (m *Module) GetAllDataUsers() (result ResponseJSON) {
	users := []Users{}
	rows, err := m.queries.GetAllDataUsers.Query()
	if err != nil {
		log.Printf("queries.GetAllDataUsers err:%+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		u := Users{}
		err := rows.Scan(&u.Username, &u.UserID, &u.Name, &u.Password, &u.LastLogin, &u.BirthDate, &u.Address, &u.RoleID)
		if err != nil {
			log.Printf("scan err:%+v\n", err)
			continue
		}
		users = append(users, u)
	}
	result = ResponseJSON{
		Data: users,
	}
	return
}
