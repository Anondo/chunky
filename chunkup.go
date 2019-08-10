package chunky

// ChunkUp divides the slices into chunks
func (itr *IntIterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Int8Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Int16Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Int32Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Int64Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *UintIterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Uint8Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Uint16Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Uint32Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Uint64Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *UintptrIterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Float32Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Float64Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Complex64Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *Complex128Iterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *StringIterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}

// ChunkUp divides the slices into chunks
func (itr *BoolIterator) ChunkUp() {

	for i := 0; i < len(itr.Chunkable); i += itr.ChunkLength {
		end := i + itr.ChunkLength
		if end > len(itr.Chunkable) {
			end = len(itr.Chunkable)
		}
		itr.Chunk = append(itr.Chunk, itr.Chunkable[i:end])
	}

}
