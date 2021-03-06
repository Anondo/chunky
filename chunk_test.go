package chunky

import (
	"reflect"
	"testing"
)

func TestIntIterator(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cl := 5
	i := IntIterator{
		Chunkable:   data,
		ChunkLength: cl,
	}

	i.ChunkUp()

	chunk1 := data[:cl]
	chunk2 := data[cl:]
	numberOfChunks := 2

	if len(i.Chunk) != numberOfChunks {
		t.Errorf("Unwanted number of chunks: wanted:%d , got:%d", numberOfChunks, len(i.Chunk))
	}

	if !reflect.DeepEqual(i.Chunk[0], chunk1) || !reflect.DeepEqual(i.Chunk[1], chunk2) {
		t.Errorf("Invalid chunks: wanted:%v & %v, got: %v & %v", chunk1, chunk2, i.Chunk[0], i.Chunk[1])
	}

}

func TestFloatIterator(t *testing.T) {

	data := []float64{1.3, 3.4, 2.3, 5.6, 3.44, 2.444, 7.65, 34.6, 11.2}
	cl := 5
	i := Float64Iterator{
		Chunkable:   data,
		ChunkLength: cl,
	}
	i.ChunkUp()

	chunk1 := data[:cl]
	chunk2 := data[cl:]
	numberOfChunks := 2

	if len(i.Chunk) != numberOfChunks {
		t.Errorf("Unwanted number of chunks: wanted:%d , got:%d", numberOfChunks, len(i.Chunk))
	}

	if !reflect.DeepEqual(i.Chunk[0], chunk1) || !reflect.DeepEqual(i.Chunk[1], chunk2) {
		t.Errorf("Invalid chunks: wanted:%v & %v, got: %v & %v", chunk1, chunk2, i.Chunk[0], i.Chunk[1])
	}
}

func TestStringIterator(t *testing.T) {

	data := []string{"Hello", "World", "This", "is", "chunky"}
	cl := 3
	i := StringIterator{
		Chunkable:   data,
		ChunkLength: cl,
	}

	i.ChunkUp()

	chunk1 := data[:cl]
	chunk2 := data[cl:]
	numberOfChunks := 2

	if len(i.Chunk) != numberOfChunks {
		t.Errorf("Unwanted number of chunks: wanted:%d , got:%d", numberOfChunks, len(i.Chunk))
	}

	if !reflect.DeepEqual(i.Chunk[0], chunk1) || !reflect.DeepEqual(i.Chunk[1], chunk2) {
		t.Errorf("Invalid chunks: wanted:%v & %v, got: %v & %v", chunk1, chunk2, i.Chunk[0], i.Chunk[1])
	}

}

func TestComplexIterator(t *testing.T) {
	data := []complex64{complex(1, 2), complex(2, 2), complex(3, 2), complex(4, 2),
		complex(5, 2), complex(6, 2), complex(7, 2), complex(8, 2), complex(9, 2)}
	cl := 3

	i := Complex64Iterator{
		Chunkable:   data,
		ChunkLength: cl,
	}

	i.ChunkUp()

	chunk1 := data[:cl]
	chunk2 := data[cl : cl*2]
	chunk3 := data[cl*2:]
	numberOfChunks := 3

	if len(i.Chunk) != numberOfChunks {
		t.Errorf("Unwanted number of chunks: wanted:%d , got:%d", numberOfChunks, len(i.Chunk))
	}

	if !reflect.DeepEqual(i.Chunk[0], chunk1) || !reflect.DeepEqual(i.Chunk[1], chunk2) ||
		!reflect.DeepEqual(i.Chunk[2], chunk3) {
		t.Errorf("Invalid chunks: wanted:%v & %v & %v, got: %v & %v & %v", chunk1, chunk2, chunk3,
			i.Chunk[0], i.Chunk[1], i.Chunk[2])
	}

}
