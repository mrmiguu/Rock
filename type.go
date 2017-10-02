package rock

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

	n struct {
		sync.RWMutex
		c chan int
	}
}

type String struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Int struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Int8 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Int16 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Int32 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Int64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Uint struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Uint8 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Uint16 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Uint32 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Uint64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Uintptr struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Byte struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Bytes struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Rune struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Float32 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Float64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Complex64 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}

type Complex128 struct {
	Name string
	Len  uint

	w, r struct {
		sync.RWMutex
		c chan []byte
	}

	n struct {
		sync.RWMutex
		c chan int
	}
}
