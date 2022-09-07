package commands

import (
	"filec2/Handlers/sessions"
	"time"
)

func init() {
	Register(&Command{
		Name: "exit",
		Execute: func(session *sessions.Session, cmd []string) error {

			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}

			if _, err := session.Println("                                ████"); err != nil {
				return err
			}
			if _, err := session.Println("                                █████████"); err != nil {
				return err
			}
			if _, err := session.Println("                                System Logoff"); err != nil {
				return err
			}
			if _, err := session.Println("                                ███████"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}

			if _, err := session.Println("                                █████"); err != nil {
				return err
			}
			if _, err := session.Println("                                ████████"); err != nil {
				return err
			}
			if _, err := session.Println("                                System Logoff"); err != nil {
				return err
			}
			if _, err := session.Println("                                ████████"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}

			if _, err := session.Println("                                ████"); err != nil {
				return err
			}
			if _, err := session.Println("                                ███████"); err != nil {
				return err
			}
			if _, err := session.Println("                                System Logoff"); err != nil {
				return err
			}
			if _, err := session.Println("                                █████████"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)
			if err := session.Clear(); err != nil {
				return err
			}

			if _, err := session.Print("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"); err != nil {
				return err
			}

			if _, err := session.Println("                                ██████"); err != nil {
				return err
			}
			if _, err := session.Println("                                ███████"); err != nil {
				return err
			}
			if _, err := session.Println("                                System Logoff"); err != nil {
				return err
			}
			if _, err := session.Println("                                ██████████"); err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 300)

			return session.Conn.Close()
		},
	})
}
