package views

import (
	"filec2/Handlers/sessions"
	"time"
	"math/rand"
	"filec2/config"
	"strconv"
	"fmt"
)

//CapchaPage will render the Capcha page
func CapchaPage(session *sessions.Session) error {

	for attempts := 0; attempts <= config.MaxLoginTries; attempts++ {

	num1 := rand.Intn(20)
	num2 := rand.Intn(13)
session.Println("\x1b[0m")
session.Println("\x1b[0m\x1b[38;5;205m                           ╔═╗  ╔═╗  ╔═╗  ╔═╗  ╦ ╦  ╔═╗")
session.Println("\x1b[0m\x1b[38;5;205m                           ║    ╠═╣  ╠═╝  ║    ╠═╣  ╠═╣")
session.Println("\x1b[0m\x1b[38;5;205m                           ╚═╝  ╩ ╩  ╩    ╚═╝  ╩ ╩  ╩ ╩")
session.Println("\x1b[0m\x1b[38;5;205m                      ══════════╦═══════════════╦══════════")
session.Println("\x1b[0m\x1b[38;5;205m               ╔════════════════╩═══════════════╩═════════════════╗")
session.Println("\x1b[0m\x1b[38;5;205m             ╔═╝          \x1b[38;5;220mEnter The Capcha You See Below!         \x1b[38;5;205m╚═╗")                                                                       
session.Println("\x1b[0m\x1b[38;5;205m                                       "+strconv.Itoa(num1)+"+"+strconv.Itoa(num2)+"")
                                     
session.Println("\x1b[0m")
session.Println("\x1b[0m")
session.Println("\x1b[0m")
session.Println("\x1b[0m")
	answer, err := session.ReadCustom("\x1b[38;5;226m⚡\x1b[38;5;127;8;8m ", false, 20)
	if err != nil {
		return err
	}
	if answer == strconv.Itoa(num1+num2) {
		return nil
	}
	if attempts == config.MaxLoginTries-1 {
		break
	}

	cooldown := 3
	for i := 0; i <= cooldown; i++ {
		session.Write([]byte(fmt.Sprintf("\x1b[0m\x1b[38;2;14;161;129m>\x1b[38;5;127;8;8m Retry in %d seconds", cooldown-i)))
		time.Sleep(time.Second)
	}
}
	return LoginTooManyAttempts(session)
}
