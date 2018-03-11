package jaccard

import "testing"

func TestStringSetAdd(t *testing.T) {
	stringSet := NewStringSet()
	if stringSet.contains("something") {
		t.Fail()
	}
	stringSet.add("something")
	if !stringSet.contains("something") {
		t.Fail()
	}
}
