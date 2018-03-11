package jaccard

type StringSet struct {
	set map[string]bool
}

func NewStringSet() StringSet {
	m := make(map[string]bool)
	return StringSet{m}
}

func (s *StringSet) Add(another string) {
	s.set[another] = true
}

func (s StringSet) Contains(another string) bool {
	_, found := s.set[another]
	return found
}

func (s StringSet) Elements() []string {
	var keys []string
	for key := range s.set {
		keys = append(keys, key)
	}
	return keys
}

func (s StringSet) Intersection(another StringSet) StringSet {
	intersection := NewStringSet()
	for _, element := range s.Elements() {
		if another.Contains(element) {
			intersection.Add(element)
		}
	}
	return intersection
}
