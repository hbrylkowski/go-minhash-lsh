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
	var elements []string
	var anotherPointer *StringSet
	if len(s.set) > len(another.set) {
		elements = another.Elements()
		anotherPointer = &s
	} else {
		elements = s.Elements()
		anotherPointer = &another
	}

	for _, element := range elements {
		if anotherPointer.Contains(element) {
			intersection.Add(element)
		}
	}
	return intersection
}

func JaccardsIndex(setA StringSet, setB StringSet) float64 {
	intersectionSize := 0
	var elements []string
	var anotherPointer *StringSet
	if len(setA.set) > len(setB.set) {
		elements = setB.Elements()
		anotherPointer = &setA
	} else {
		elements = setA.Elements()
		anotherPointer = &setB
	}

	for _, element := range elements {
		if anotherPointer.Contains(element) {
			intersectionSize++
		}
	}
	return float64(intersectionSize) / float64(len(setA.set)+len(setB.set)-intersectionSize)
}
