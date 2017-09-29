package iota

type String struct {
    s string
}

func (s *String) Load() string {
    return ""
}

func (s *String) Store(s string) {
}
