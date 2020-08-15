package got

import (
	"sync"
)

// Chunk is a partial content range.
type Chunk struct {

	// Chunk start pos.
	Start uint64

	// Chunk end pos.
	End uint64

	// Path name where this chunk downloaded.
	Path string

	// Done to check is this chunk downloaded.
	Done chan struct{}
}

// Reset resets the chunk item to the initial state.
func (c *Chunk) Reset() {
	c.Start = 0
	c.End = 0
	c.Path = ""
	c.Done = make(chan struct{})
}

// ChunkPool helps in multi *Download files.
var ChunkPool = &sync.Pool{
	New: func() interface{} {
		return new(Chunk)
	},
}
