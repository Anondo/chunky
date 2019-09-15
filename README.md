# CHUNKY

[![Build Status](https://travis-ci.org/Anondo/chunky.svg?branch=master)](https://travis-ci.org/Anondo/chunky)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/Anondo/chunky/blob/master/LICENSE)
[![Project status](https://img.shields.io/badge/version-1.0.1-green.svg)](https://github.com/Anondo/chunky/releases)

Not the package you needed but the package you deserved. Chunk the slices up by creating an iterator specifically
for each data types with ease. Because you are lazy maybe?

### Installing
```console
go get -u github.com/Anondo/chunky

```

### Usage

**Import The Package**

```go
import "github.com/Anondo/chunky"

```

**Create An Iterator & Iterate Over The Chunks**

```go
data := []int{}

for i := 0; i < 10000; i++ {
  data = append(data, i)
}

itr := chunky.IntIterator{
  Chunkable:   data,
  ChunkLength: 450,
}

itr.ChunkUp()

log.Println("Chunks:")

for itr.Next() {
  log.Println(itr.CurrentBlock)
}

```

**OR**

Use the whole chunk

```go

log.Println(itr.Chunk)

```

**Using The Factory Function**

```go
data := []int{}

for i := 0; i < 10000; i++ {
	data = append(data, i)
}

itr := chunky.NewIntIterator(data, 450)

log.Println("Chunks:")

for itr.Next() {
	log.Println(itr.GetCurrentBlock().([]int)) // Need to type cast because GetCurrentBlock returns an empty interface
}

```

You need to type cast the chunks to use it, as the ```factory functions``` return an **Iterator** interface.
So to play with the actual chunk:

```go

log.Println(itr.(*chunky.IntIterator).Chunk)

```

Try out the iterators for other data types as well.


### Contributing

See the contributions guide [here](CONTRIBUTING.md).

### License

Chunky is licensed under the [MIT License](LICENSE).
