package terminal

import "fmt"

//Write will write directly out to the connection
func (term *Terminal) Write(data []byte) (int, error) {
	return term.wr.Write(data)
}

//Println will write a new line
func (term *Terminal) Println(d ...interface{}) (int, error) {
	return term.wr.Write([]byte(fmt.Sprint(append(d, "\r\n")...)))
}

//Print will write to the terminal
func (term *Terminal) Print(d ...interface{}) (int, error) {
	return term.wr.Write([]byte(fmt.Sprint(d...)))
}
