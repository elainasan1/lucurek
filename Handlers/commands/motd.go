package commands

import (
	"filec2/Handlers/sessions"
	"sync"
)

var (
	motd      string
	motdMutex sync.RWMutex
)

// Motd returns the message of the day
func Motd() string {
	motdMutex.RLock()
	defer motdMutex.RUnlock()

	return motd
}

func init() {
	Register(&Command{
		Name:  "motd",
		Admin: true,
		Execute: func(session *sessions.Session, cmd []string) error {

			if _, err := session.Write([]byte("\x1b[0m    Motd> \x1b[47m\x1b[30m                                         \r\n\x1b[1A")); err != nil {
				return err
			}

			motdSet, err := session.ReadCustom("\x1b[0m    Motd> \x1b[47m\x1b[30m", false, 40)
			if err != nil {
				return err
			}

			motdMutex.Lock()
			defer motdMutex.Unlock()
			motd = motdSet

			if _, err := session.Write([]byte("\x1b[0m\r\n    Message Of The Day Set\r\n")); err != nil {
				return err
			}

			return nil
		},
	})
}
