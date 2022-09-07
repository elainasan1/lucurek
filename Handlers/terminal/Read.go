package terminal

//ReadLine will read the input
func (term *Terminal) ReadLine(prompt string) (string, error) {

	return term.readLine(prompt, false, 200)
}

//ReadPassword will read the input but also mask it
func (term *Terminal) ReadPassword(prompt string) (string, error) {

	return term.readLine(prompt, true, 200)
}

//ReadCustom will read the input but allow you set advanced paramaters manually
func (term *Terminal) ReadCustom(prompt string, password bool, maxLen int) (string, error) {
	return term.readLine(prompt, password, maxLen)
}

func (term *Terminal) readLine(prompt string, password bool, maxLen int) (string, error) {
	term.mutex.Lock()
	defer term.mutex.Unlock()

	var line []byte

	if _, err := term.Write([]byte("\r")); err != nil {
		return "", nil
	}

	if _, err := term.Write([]byte(prompt)); err != nil {
		return "", nil
	}

	index := 0

	for {
		buf := make([]byte, 1)
		_, err := term.wr.Read(buf)
		if err != nil {
			return "", err
		}

		switch buf[0] {
		case 27: //Control
			sequence := make([]byte, 2)
			_, err := term.wr.Read(sequence)
			if err != nil {
				return "", err
			}

			break
		case 13:
			break
		case 10, 9:
			term.Write([]byte("\r\n"))
			return string(line), nil
		case 127:
			if len(line) == 0 {
				break
			}

			index--
			term.Write([]byte{127})
			line = line[:len(line)-1]
			break
		default:

			// line = append(line[:index], append(buf, line[:index]...)...)

			if len(line) < maxLen {
				if password {
					term.Write([]byte("•"))
				} else {
					term.Write(buf)
				}

				line = append(line, buf...)
				index++
			}
		}
	}
}
