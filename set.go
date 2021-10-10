package app

type Set struct {
	size int
}

func NewSet() *Set {
	return &Set{0}
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}
