package chunky

import (
	"errors"
)

// Iterator ...
type Iterator struct {
	Chunkable   interface{}
	ChunkLength int
	chunk       interface{}
}

// NewIterator ...
func NewIterator(c interface{}, cl int) (*Iterator, error) {
	if !isProperType(c) {
		return nil, errors.New("Must be a valid slice")
	}
	chnk := getAssertedSlice(c)
	if chnk == nil {
		return nil, errors.New("Could not assert interface into a valid slice")
	}

	i := &Iterator{
		Chunkable:   chnk,
		ChunkLength: cl,
	}

	return i, nil
}

// NewIntIterator ...
func NewIntIterator(c []int, cl int) *Iterator {
	return &Iterator{
		Chunkable:   c,
		ChunkLength: cl,
	}
}
