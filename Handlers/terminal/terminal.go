package terminal

import (
	"io"
	"sync"
)

//Terminal handles reading and writing
type Terminal struct {
	wr io.ReadWriter

	mutex sync.Mutex
}
