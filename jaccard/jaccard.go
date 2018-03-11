package jaccard

type StringSet struct {
	set map[string]bool
}

func NewStringSet() StringSet {
	m := make(map[string]bool)
	return StringSet{m}
}

func (s *StringSet) add(another string) {
	s.set[another] = true
}

func (s StringSet) contains(another string) bool {
	_, found := s.set[another]
	return found
}
