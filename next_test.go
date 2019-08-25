package chunky

import (
	"reflect"
	"testing"
)

func TestIntIteratorNext(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cl := 5
	itr := NewIntIterator(data, cl)

	chunk1 := data[:cl]
	chunk2 := data[cl:]
	chunks := [][]int{chunk1, chunk2}

	for i := 0; itr.Next(); i++ {
		cb := itr.GetCurrentBlock().([]int)
		if !reflect.DeepEqual(cb, chunks[i]) {
			t.Errorf("The Next() method didn't return the chunk block expected, wanted:%v got %v",
				chunk1, cb)
		}
	}

}

func TestFloatIteratorNext(t *testing.T) {

	data := []float64{1.3, 3.4, 2.3, 5.6, 3.44, 2.444, 7.65, 34.6, 11.2}
	cl := 5
	itr := NewFloat64Iterator(data, cl)

	chunk1 := data[:cl]
	chunk2 := data[cl:]
	chunks := [][]float64{chunk1, chunk2}

	for i := 0; itr.Next(); i++ {
		cb := itr.GetCurrentBlock().([]float64)
		if !reflect.DeepEqual(cb, chunks[i]) {
			t.Errorf("The Next() method didn't return the chunk block expected, wanted:%v got %v",
				chunk1, cb)
		}
	}

}

func TestStringIteratorNext(t *testing.T) {

	data := []string{"Hello", "World", "This", "is", "chunky"}
	cl := 3
	itr := NewStringIterator(data, cl)

	chunk1 := data[:cl]
	chunk2 := data[cl:]
	chunks := [][]string{chunk1, chunk2}

	for i := 0; itr.Next(); i++ {
		cb := itr.GetCurrentBlock().([]string)
		if !reflect.DeepEqual(cb, chunks[i]) {
			t.Errorf("The Next() method didn't return the chunk block expected, wanted:%v got %v",
				chunk1, cb)
		}
	}

}

func TestComplexIteratorNext(t *testing.T) {

	data := []complex64{complex(1, 2), complex(2, 2), complex(3, 2), complex(4, 2),
		complex(5, 2), complex(6, 2), complex(7, 2), complex(8, 2), complex(9, 2)}
	cl := 3

	itr := NewComplex64Iterator(data, cl)

	chunk1 := data[:cl]
	chunk2 := data[cl : cl*2]
	chunk3 := data[cl*2:]
	chunks := [][]complex64{chunk1, chunk2, chunk3}

	for i := 0; itr.Next(); i++ {
		cb := itr.GetCurrentBlock().([]complex64)
		if !reflect.DeepEqual(cb, chunks[i]) {
			t.Errorf("The Next() method didn't return the chunk block expected, wanted:%v got %v",
				chunk1, cb)
		}
	}

}
