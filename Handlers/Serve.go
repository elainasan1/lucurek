package handlers

import (
	"filec2/config"
	"log"
	"net"
)

//Serve listens for opperator connections
func Serve() {
	l, err := net.Listen("tcp", config.MasterHost)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [Master server starting] [%s]\n", config.MasterHost)
	go titleWorker()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("handlers/serve:", err)
			continue
		}

		go handle(conn)
	}
}
