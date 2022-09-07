package db

import (
	"database/sql"
	"log"
)

//User holds information on a user
type User struct {
	ID       int
	Username string
	Password string

	Admin  bool
	Banned bool

	Cooldown    int
	MaxTime     int
	Conns       int
	Expiry      int64
	MaxSessions int
}

//NewUser will create a new User in the database
func NewUser(user *User) error {

	rows, err := SQL.Query("INSERT INTO `users` (`username`, `password`, `admin`, `banned`, `cooldown`, `maxTime`, `conns`, `expiry`, `maxSessions`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.Username,
		user.Password,

		user.Admin,
		user.Banned,

		user.Cooldown,
		user.Conns,
		user.MaxTime,
		user.Expiry,
		user.MaxSessions,
	)
	if rows != nil {
		rows.Close()
	}

	return err
}

//EditUser will modify a new User in the database
func EditUser(user *User) error {

	rows, err := SQL.Query("UPDATE `users` set `username` = ?, `password` = ?, `admin` = ?, `banned` = ?, `cooldown` = ?, `maxTime` = ?, `conns` = ?, `expiry` = ?, `maxSessions` = ? WHERE `id` = ?",
		user.Username,
		user.Password,

		user.Admin,
		user.Banned,

		user.Cooldown,
		user.Conns,
		user.MaxTime,
		user.Expiry,
		user.MaxSessions,

		user.ID,
	)
	if rows != nil {
		rows.Close()
	}

	return err
}

//GetUser will get a user via their username and return a user of nil if it was not found
func GetUser(username string) (*User, error) {

	user := &User{}
	err := scanUser(SQL.QueryRow("SELECT `id`, `username`, `password`, `admin`, `banned`, `cooldown`, `maxTime`, `conns`, `expiry`, `maxSessions` FROM  `users` where `username` = ?", username), user)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return user, err
	}

	return user, nil
}

//GetUsers will return a list of all users
func GetUsers() ([]*User, error) {

	var users []*User
	rows, err := SQL.Query("SELECT `id`, `username`, `password`, `admin`, `banned`, `cooldown`, `maxTime`, `conns`, `expiry`, `maxSessions` FROM  `users`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &User{}
		if err := scanUsers(rows, user); err != nil {
			log.Println("db/getUsers:", err)
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func scanUser(row *sql.Row, user *User) error {

	return row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,

		&user.Admin,
		&user.Banned,

		&user.Cooldown,
		&user.Conns,
		&user.MaxTime,
		&user.Expiry,
		&user.MaxSessions,
	)
}

func scanUsers(row *sql.Rows, user *User) error {

	return row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,

		&user.Admin,
		&user.Banned,

		&user.Cooldown,
		&user.Conns,
		&user.MaxTime,
		&user.Expiry,
		&user.MaxSessions,
	)
}

//Delete will remove the user from the database
func (user *User) Delete() error {
	rows, err := SQL.Query(
		"DELETE FROM `users` WHERE `id` = ?",
		user.ID,
	)

	rows.Close()

	return err
}
