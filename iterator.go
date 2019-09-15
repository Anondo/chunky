package chunky

type (

	// Iterator ...
	Iterator interface {
		ChunkUp()
		Next() bool
		GetCurrentBlock() interface{}
	}

	//IntIterator ...
	IntIterator struct {
		Chunkable    []int
		ChunkLength  int
		Chunk        [][]int
		CurrentBlock []int
	}
	//Int8Iterator ...
	Int8Iterator struct {
		Chunkable    []int8
		ChunkLength  int
		Chunk        [][]int8
		CurrentBlock []int8
	}
	//Int16Iterator ...
	Int16Iterator struct {
		Chunkable    []int16
		ChunkLength  int
		Chunk        [][]int16
		CurrentBlock []int16
	}
	//Int32Iterator ...
	Int32Iterator struct {
		Chunkable    []int32
		ChunkLength  int
		Chunk        [][]int32
		CurrentBlock []int32
	}
	//Int64Iterator ...
	Int64Iterator struct {
		Chunkable    []int64
		ChunkLength  int
		Chunk        [][]int64
		CurrentBlock []int64
	}
	//UintIterator ...
	UintIterator struct {
		Chunkable    []uint
		ChunkLength  int
		Chunk        [][]uint
		CurrentBlock []uint
	}
	//Uint8Iterator ...
	Uint8Iterator struct {
		Chunkable    []uint8
		ChunkLength  int
		Chunk        [][]uint8
		CurrentBlock []uint8
	}
	//Uint16Iterator ...
	Uint16Iterator struct {
		Chunkable    []uint16
		ChunkLength  int
		Chunk        [][]uint16
		CurrentBlock []uint16
	}
	//Uint32Iterator ...
	Uint32Iterator struct {
		Chunkable    []uint32
		ChunkLength  int
		Chunk        [][]uint32
		CurrentBlock []uint32
	}
	//Uint64Iterator ...
	Uint64Iterator struct {
		Chunkable    []uint64
		ChunkLength  int
		Chunk        [][]uint64
		CurrentBlock []uint64
	}
	//UintptrIterator ...
	UintptrIterator struct {
		Chunkable    []uintptr
		ChunkLength  int
		Chunk        [][]uintptr
		CurrentBlock []uintptr
	}
	//Float32Iterator ...
	Float32Iterator struct {
		Chunkable    []float32
		ChunkLength  int
		Chunk        [][]float32
		CurrentBlock []float32
	}
	//Float64Iterator ...
	Float64Iterator struct {
		Chunkable    []float64
		ChunkLength  int
		Chunk        [][]float64
		CurrentBlock []float64
	}
	//Complex64Iterator ...
	Complex64Iterator struct {
		Chunkable    []complex64
		ChunkLength  int
		Chunk        [][]complex64
		CurrentBlock []complex64
	}
	//Complex128Iterator ...
	Complex128Iterator struct {
		Chunkable    []complex128
		ChunkLength  int
		Chunk        [][]complex128
		CurrentBlock []complex128
	}
	//StringIterator ...
	StringIterator struct {
		Chunkable    []string
		ChunkLength  int
		Chunk        [][]string
		CurrentBlock []string
	}
	//BoolIterator ...
	BoolIterator struct {
		Chunkable    []bool
		ChunkLength  int
		Chunk        [][]bool
		CurrentBlock []bool
	}
)

// NewIntIterator ...
func NewIntIterator(chunkable []int, chunkLength int) Iterator {
	i := &IntIterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewInt8Iterator ...
func NewInt8Iterator(chunkable []int8, chunkLength int) Iterator {
	i := &Int8Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewInt16Iterator ...
func NewInt16Iterator(chunkable []int16, chunkLength int) Iterator {
	i := &Int16Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewInt32Iterator ...
func NewInt32Iterator(chunkable []int32, chunkLength int) Iterator {
	i := &Int32Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewInt64Iterator ...
func NewInt64Iterator(chunkable []int64, chunkLength int) Iterator {
	i := &Int64Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewUintIterator ...
func NewUintIterator(chunkable []uint, chunkLength int) Iterator {
	i := &UintIterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewUint8Iterator ...
func NewUint8Iterator(chunkable []uint8, chunkLength int) Iterator {
	i := &Uint8Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewUint16Iterator ...
func NewUint16Iterator(chunkable []uint16, chunkLength int) Iterator {
	i := &Uint16Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewUint32Iterator ...
func NewUint32Iterator(chunkable []uint32, chunkLength int) Iterator {
	i := &Uint32Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewUint64Iterator ...
func NewUint64Iterator(chunkable []uint64, chunkLength int) Iterator {
	i := &Uint64Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewUintptrIterator ...
func NewUintptrIterator(chunkable []uintptr, chunkLength int) Iterator {
	i := &UintptrIterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewFloat32Iterator ...
func NewFloat32Iterator(chunkable []float32, chunkLength int) Iterator {
	i := &Float32Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewFloat64Iterator ...
func NewFloat64Iterator(chunkable []float64, chunkLength int) Iterator {
	i := &Float64Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewComplex64Iterator ...
func NewComplex64Iterator(chunkable []complex64, chunkLength int) Iterator {
	i := &Complex64Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewComplex128Iterator ...
func NewComplex128Iterator(chunkable []complex128, chunkLength int) Iterator {
	i := &Complex128Iterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewStringIterator ...
func NewStringIterator(chunkable []string, chunkLength int) Iterator {
	i := &StringIterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// NewBoolIterator ...
func NewBoolIterator(chunkable []bool, chunkLength int) Iterator {
	i := &BoolIterator{
		Chunkable:   chunkable,
		ChunkLength: chunkLength,
	}

	i.ChunkUp()

	return i

}

// GetCurrentBlock returns the current block of chunk
func (itr *IntIterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Int8Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Int16Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Int32Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Int64Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *UintIterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock
}

// GetCurrentBlock returns the current block of chunk
func (itr *Uint8Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Uint16Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Uint32Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Uint64Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *UintptrIterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Float32Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Float64Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Complex64Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *Complex128Iterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *StringIterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}

// GetCurrentBlock returns the current block of chunk
func (itr *BoolIterator) GetCurrentBlock() interface{} {

	return itr.CurrentBlock

}
