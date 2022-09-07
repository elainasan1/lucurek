package handlers

import (
	"filec2/Handlers/commands"
	"filec2/Handlers/sessions"
	"fmt"
	"log"
	"time"
)

func titleWorker() {
	log.Printf(" [Title worker started]\n")

	for {
		sessions.Broadcast([]byte(fmt.Sprintf("\033]0;Pika C2 | API'S Online [5] | Users [%d] | MOTD: [%s]\007",  sessions.Count(), commands.Motd())))

		time.Sleep(time.Second)
	}
}
