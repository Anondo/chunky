package chunky

import (
	"reflect"
)

//the slice types
const (
	Int        = "[]int"
	Int8       = "[]int8"
	Int16      = "[]int16"
	Int32      = "[]int32"
	Int64      = "[]int64"
	Uint       = "[]uint"
	Uint8      = "[]uint8"
	Uint16     = "[]uint16"
	Uint32     = "[]uint32"
	Uintptr    = "[]uintptr"
	Float32    = "[]float32"
	Float64    = "[]float64"
	Complex64  = "[]complex64"
	Complex128 = "[]complex128"
	String     = "[]string"
	Bool       = "[]bool"
)

var (
	types = map[string]bool{
		Int:        true,
		Int8:       true,
		Int16:      true,
		Int32:      true,
		Int64:      true,
		Uint:       true,
		Uint8:      true,
		Uint16:     true,
		Uint32:     true,
		Uintptr:    true,
		Float32:    true,
		Float64:    true,
		Complex64:  true,
		Complex128: true,
		String:     true,
		Bool:       true,
	}
)

func isProperType(i interface{}) bool {
	typeStr := reflect.TypeOf(i).String()
	_, ok := types[typeStr]
	return ok
}

func getAssertedSlice(i interface{}) interface{} {
	switch i.(type) {
	case []int:
		return i.([]int)
	case []int8:
		return i.([]int8)
	case []int16:
		return i.([]int16)
	case []int32:
		return i.([]int32)
	case []int64:
		return i.([]int64)
	case []uint:
		return i.([]uint)
	case []uint8:
		return i.([]uint8)
	case []uint16:
		return i.([]uint16)
	case []uint32:
		return i.([]uint32)
	case []uint64:
		return i.([]uint64)
	case []uintptr:
		return i.([]uintptr)
	case []float32:
		return i.(float32)
	case []float64:
		return i.(float64)
	case []complex64:
		return i.(complex64)
	case []complex128:
		return i.(complex128)
	case []string:
		return i.([]string)
	case []bool:
		return i.([]bool)
	}

	return nil
}
