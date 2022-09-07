package sessions

//Remove deletes the session
func (s *Session) Remove() {
	mutex.Lock()
	defer mutex.Unlock()

	close(s.Queue)

	delete(sessions, s.ID)
}
