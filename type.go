package rock

import (
	"sync"
)

type Bool struct {
	Name string
	Len  uint

	p private
}

type String struct {
	Name string
	Len  uint

	p private
}

type Int struct {
	Name string
	Len  uint

	p private
}

type Int8 struct {
	Name string
	Len  uint

	p private
}

type Int16 struct {
	Name string
	Len  uint

	p private
}

type Int32 struct {
	Name string
	Len  uint

	p private
}

type Int64 struct {
	Name string
	Len  uint

	p private
}

type Uint struct {
	Name string
	Len  uint

	p private
}

type Uint8 struct {
	Name string
	Len  uint

	p private
}

type Uint16 struct {
	Name string
	Len  uint

	p private
}

type Uint32 struct {
	Name string
	Len  uint

	p private
}

type Uint64 struct {
	Name string
	Len  uint

	p private
}

type Uintptr struct {
	Name string
	Len  uint

	p private
}

type Byte struct {
	Name string
	Len  uint

	p private
}

type Bytes struct {
	Name string
	Len  uint

	p private
}

type Rune struct {
	Name string
	Len  uint

	p private
}

type Float32 struct {
	Name string
	Len  uint

	p private
}

type Float64 struct {
	Name string
	Len  uint

	p private
}

type Complex64 struct {
	Name string
	Len  uint

	p private
}

type Complex128 struct {
	Name string
	Len  uint

	p private
}

type private struct {
	w, r struct {
		sync.Once
		c chan []byte
	}
	n struct {
		sync.Once
		c chan int
	}
}
