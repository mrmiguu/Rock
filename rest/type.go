package rest

import (
	"sync"
)

type Bool struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type String struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Int struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Int8 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Int16 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Int32 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Int64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Uint struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Uint8 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Uint16 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Uint32 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Uint64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Uintptr struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Byte struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Bytes struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Rune struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Float32 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Float64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Complex64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
type Complex128 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}
}
