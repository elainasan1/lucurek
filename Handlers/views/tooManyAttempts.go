package views

import (
	"filec2/Handlers/sessions"
	"time"
)

//LoginTooManyAttempts will render the LoginTooManyAttempts page
func LoginTooManyAttempts(session *sessions.Session) error {

	session.Print("\x1b[101m\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n")
	session.Println("\x1b[13A                                  \x1b[37;1mAccess Denied\x1b[0m\x1b[101m")
	session.Println("            You have been kicked for too many failed login attempts")

	time.Sleep(time.Second * 30)
	return ErrUnauthorised
}
