package commands

import "filec2/Handlers/sessions"

func init() {
	Register(&Command{
		Name: "clear",
		Execute: func(session *sessions.Session, cmd []string) error {

			if err := session.Clear(); err != nil {
				return err
			}

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

			return nil
		},
	})
}
