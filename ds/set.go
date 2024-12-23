package ds

type Set[T comparable] struct {
	set map[T]bool
}

func (s *Set[T]) Add(elem T) {
	s.set[elem] = true
}

func (s *Set[T]) Size() int {
	return len(s.set)
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{make(map[T]bool)}
}
