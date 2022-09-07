package views

import (
	"filec2/Handlers/sessions"
	"time"
)

//Userbanned will render the Userbanned page
func Userbanned(session *sessions.Session) error {

	session.Print("\x1b[101m\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n")
	session.Println("\x1b[13A                                  \x1b[37;1mAccess Denied\x1b[0m\x1b[101m")
	session.Println("              Your access has been revoked to this system! Contact")
	session.Println("                                a administrator.")

	time.Sleep(time.Second * 30)
	return ErrUnauthorised
}
