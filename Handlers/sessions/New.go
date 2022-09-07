package sessions

import (
	"crypto/rand"
	"encoding/binary"
)

//New creates a new session
func New(session *Session) error {
	mutex.Lock()
	defer mutex.Unlock()

	buf := make([]byte, 8)
	for {
		if _, err := rand.Read(buf); err != nil {
			return err
		}

		id := binary.BigEndian.Uint64(buf)
		if _, ok := sessions[id]; ok == false {
			session.ID = id
			session.Queue = make(chan []byte, queueSize)
			sessions[id] = session
			return nil
		}
	}
}
