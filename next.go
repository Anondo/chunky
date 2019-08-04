package chunky

// Next pops the chunks sequentially
func (itr *IntIterator) Next() []int {
	var chnk []int
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Int8Iterator) Next() []int8 {
	var chnk []int8
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Int16Iterator) Next() []int16 {
	var chnk []int16
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Int32Iterator) Next() []int32 {
	var chnk []int32
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Int64Iterator) Next() []int64 {
	var chnk []int64
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Float32Iterator) Next() []float32 {
	var chnk []float32
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Float64Iterator) Next() []float64 {
	var chnk []float64
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *StringIterator) Next() []string {
	var chnk []string
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *BoolIterator) Next() []bool {
	var chnk []bool
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *UintIterator) Next() []uint {
	var chnk []uint
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Uint8Iterator) Next() []uint8 {
	var chnk []uint8
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Uint16Iterator) Next() []uint16 {
	var chnk []uint16
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Uint32Iterator) Next() []uint32 {
	var chnk []uint32
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Uint64Iterator) Next() []uint64 {
	var chnk []uint64
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *UintptrIterator) Next() []uintptr {
	var chnk []uintptr
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Complex64Iterator) Next() []complex64 {
	var chnk []complex64
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}

// Next pops the chunks sequentially
func (itr *Complex128Iterator) Next() []complex128 {
	var chnk []complex128
	chnk, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
	return chnk
}
