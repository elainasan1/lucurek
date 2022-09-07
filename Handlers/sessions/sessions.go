package sessions

import (
	"net"
	"sync"
	"time"

	"filec2/db"
	"filec2/Handlers/terminal"
)

var (
	sessions = make(map[uint64]*Session)
	mutex    sync.RWMutex

	queueSize = 60
)

//Session holds information on a current session
type Session struct {
	ID uint64

	User *db.User

	*terminal.Terminal
	Conn net.Conn

	Queue chan []byte

	Created time.Time
}

//Count returns the amount of online users
func Count() int {
	mutex.RLock()
	defer mutex.RUnlock()

	return len(sessions)
}

//Broadcast forwards a payload to all clients able to accept it
func Broadcast(payload []byte) {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, session := range sessions {
		if len(session.Queue) >= queueSize-1 {
			continue
		}

		session.Queue <- payload
	}
}

//Clone will copy the sessions map and return it
func Clone() []Session {
	mutex.RLock()
	defer mutex.RUnlock()

	var list []Session

	for _, session := range sessions {
		list = append(list, *session)
	}

	return list
}
