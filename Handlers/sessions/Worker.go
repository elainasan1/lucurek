package sessions

//Worker handles commands sent to the session
func (s *Session) Worker() {
	for {
		data, open := <-s.Queue
		if open == false {
			return
		}

		s.Write(data)

	}
}
