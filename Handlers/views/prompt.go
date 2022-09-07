package views

import (
	"filec2/Handlers/sessions"
	"log"
	"strings"
  "fmt"
	"filec2/attacks"

	"filec2/Handlers/commands"
)

//PromptPage is where commands are entered
func PromptPage(session *sessions.Session) error {
	if _, err := session.Print(
"     Welcome User, Happy To Have You Here At Sky Services! API'S ONLINE: [1]\r\n",
"\r\n",
"                   ███████ ██   ██ ██    ██      ██████ ██████\r\n",
"                   ██      ██  ██   ██  ██      ██           ██\r\n",
"                   ███████ █████     ████       ██       █████\r\n",
"                        ██ ██  ██     ██        ██      ██\r\n",
"                   ███████ ██   ██    ██         ██████ ███████\r\n",
"                      ══════════╦═══════════════╦══════════\r\n", 
"               ╔════════════════╩═══════════════╩═════════════════╗\r\n", 
"               ║       Welcome To The Start Screen Of Sky C2      ║\r\n",
"               ║              Powered By Sky API                  ║\r\n", 
"               ╚═══════════════╦════════════════╦═════════════════╝\r\n",
"               ╔═══════════════╩════════════════╩═════════════════╗\r\n", 
"               ║    Copyright © 2022 Sky C2 All Rights Reserved   ║\r\n",   
"               ╚══════════════════════════════════════════════════╝\r\n",); err != nil {
		return err
	}

	for {
		if _, err := session.Println("\x1b[0m"); err != nil {
			return err
		}
		command, err := session.ReadCustom("\x1b[38;5;205mSkyC2\x1b[38;5;226m@"+fmt.Sprintf("\x1b[38;5;205m%s", session.User.Username)+"\x1b[16m:\x1b[38;5;205m ", false, 100)
		if err != nil {
			log.Println("CMD", err)
			return err
		}
		if command == "" {
			continue
		}
		if _, err := session.Print("\x1b[2K"); err != nil {
			return err
		}

		executeCommand(session, command)
	}

}

func executeCommand(session *sessions.Session, command string) {
	cmd := strings.Split(command, " ")
	//if cmd[0] == "" {
	//	return
	//}

	c := commands.Get(strings.ToLower(cmd[0]))
	if c == nil {

		method := attacks.Get(cmd[0])
		if method == nil {
			session.Println("\x1b[38;5;205mSkyC2\x1b[38;5;226m@\x1b[196mError_Report\x1b[16m:\x1b[38;5;205m INVALID INPUT")
			return
		}

		if err := AttackPage(session, cmd, method); err != nil {
			log.Println("views/prompt.Attack:", err)
			session.Println("\x1b[38;5;205mSkyC2\x1b[38;5;226m@\x1b[196mError_Report\x1b[16m:\x1b[38;5;205m ERROR OCCURRED - RETRY OR CONTACT AN ADMIN")
			return
		}

		return
	}

	if !c.HasPermission(session) {
		session.Println("\x1b[38;5;205mSkyC2\x1b[38;5;226m@\x1b[196mError_Report\x1b[16m:\x1b[38;5;205m YOU DO NOT HAVE AUTHORISATION TO EXECUTE THIS COMMAND.")
		return
	}

	if err := c.Execute(session, cmd); err != nil {
		session.Println("\x1b[38;5;205mSkyC2\x1b[38;5;226m@\x1b[196mError_Report\x1b[16m:\x1b[38;5;205m ERROR OCCURRED - RETRY OR CONTACT AN ADMIN")

		if session.User.Admin {
			session.Println(err)
		}
	}

}

func paddingLeft(input string, pad string, length int) string {
	if len(input) > length {
		return input
	}

	for i := 1; i <= length; i++ {
		input += pad
	}

	return input
}
