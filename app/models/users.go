package models

// import golang packages
import (
	"time"
)

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Email             string    `json:"email"`
	Enabled           bool      `json:"enabled"`
	Lastpasswordreset time.Time `json:"lastpasswordreset"`
}

func AllUsers() ([]*User, error) {
	rows, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Enabled, &user.Lastpasswordreset)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
