package terminal

//Clear will clear the terminal screen
func (term *Terminal) Clear() error {
	_, err := term.Write([]byte("\033c"))
	return err
}
