package jaccard

import (
	"testing"
	"fmt"
)

func TestNewLSH(t *testing.T) {
	lsh := NewLSH(10, 5, 1)
	lsh.Add([]string{"cow", "eats"}, 1)
	lsh.Add([]string{"cow", "eats"}, 2)
	lsh.Add([]string{"cat", "dies"}, 3)
	fmt.Printf("%#v",lsh.Buckets)
}

