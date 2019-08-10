package chunky

// Next pops the chunks sequentially
func (itr *IntIterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Int8Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Int16Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Int32Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Int64Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Float32Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Float64Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *StringIterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *BoolIterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *UintIterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Uint8Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Uint16Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Uint32Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Uint64Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *UintptrIterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Complex64Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}

// Next pops the chunks sequentially
func (itr *Complex128Iterator) Next() bool {
	if len(itr.Chunk) > 0 {
		itr.CurrentBlock, itr.Chunk = itr.Chunk[0], itr.Chunk[1:]
		return true
	}
	return false
}
