package jaccard

import (
	"testing"
)

func TestStringSetAdd(t *testing.T) {
	stringSet := NewStringSet()
	if stringSet.Contains("something") {
		t.Fail()
	}
	stringSet.Add("something")
	if !stringSet.Contains("something") {
		t.Fail()
	}
}

func TestStringSetElements(t *testing.T) {
	stringSet := NewStringSet()
	stringSet.Add("1")
	stringSet.Add("2")
	elements := stringSet.Elements()
	if len(elements) != 2 {
		t.Fail()
	}
	if !((elements[0] == "1" && elements[1] == "2") || (elements[0] == "2" && elements[1] == "1")) {
		t.Fail()
	}
}

func TestStringSetIntersection(t *testing.T) {
	stringSet := NewStringSet()
	stringSet.Add("1")
	stringSet.Add("2")
	anotherSet := NewStringSet()
	anotherSet.Add("3")
	anotherSet.Add("2")
	intersection := stringSet.Intersection(anotherSet)
	if intersection.Elements()[0] != "2" {
		t.Fail()
	}
	if len(intersection.Elements()) != 1 {
		t.Fail()
	}
}
