package rock

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
)

var (
	Addr     string
	IsClient = js.Global != nil

	v = []byte(V)

	started struct {
		sync.RWMutex
		b bool
	}

	boolDict struct {
		sync.RWMutex
		m map[string]*Bool
	}
	stringDict struct {
		sync.RWMutex
		m map[string]*String
	}
	intDict struct {
		sync.RWMutex
		m map[string]*Int
	}
	int8Dict struct {
		sync.RWMutex
		m map[string]*Int8
	}
	int16Dict struct {
		sync.RWMutex
		m map[string]*Int16
	}
	int32Dict struct {
		sync.RWMutex
		m map[string]*Int32
	}
	int64Dict struct {
		sync.RWMutex
		m map[string]*Int64
	}
	uintDict struct {
		sync.RWMutex
		m map[string]*Uint
	}
	uint8Dict struct {
		sync.RWMutex
		m map[string]*Uint8
	}
	uint16Dict struct {
		sync.RWMutex
		m map[string]*Uint16
	}
	uint32Dict struct {
		sync.RWMutex
		m map[string]*Uint32
	}
	uint64Dict struct {
		sync.RWMutex
		m map[string]*Uint64
	}
	uintptrDict struct {
		sync.RWMutex
		m map[string]*Uintptr
	}
	byteDict struct {
		sync.RWMutex
		m map[string]*Byte
	}
	bytesDict struct {
		sync.RWMutex
		m map[string]*Bytes
	}
	runeDict struct {
		sync.RWMutex
		m map[string]*Rune
	}
	float32Dict struct {
		sync.RWMutex
		m map[string]*Float32
	}
	float64Dict struct {
		sync.RWMutex
		m map[string]*Float64
	}
	complex64Dict struct {
		sync.RWMutex
		m map[string]*Complex64
	}
	complex128Dict struct {
		sync.RWMutex
		m map[string]*Complex128
	}
)
