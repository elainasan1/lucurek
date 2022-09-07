package terminal

import "io"

//New creates a new terminal
func New(wr io.ReadWriter) *Terminal {
	return &Terminal{
		wr: wr,
	}
}
