package main

import (
	"github.com/hbrylkowski/lsh/jaccard"
	"github.com/tjarratt/babble"
	"math/rand"
	"strings"
	"fmt"
	"math"
)

func main() {
}

func greatestCommonDivisor(a int, b int) int {
	if b == 0{
		return a
	}
	return greatestCommonDivisor(b, a%b)
}

func getSmallestMultiplier(b float64) int {
	precision := len(fmt.Sprint(b)) - strings.Index(fmt.Sprint(b), ".") - 1
	limit := int(math.Pow10(precision))
	divisor := greatestCommonDivisor(int(b* float64(limit)), limit)
	return limit / divisor
}

func newBabble() func(int) string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	return func(a int) string {
		babbler.Count = a
		return babbler.Babble()
	}
}


func createSentencesPair(similarity float64) (jaccard.StringSet, jaccard.StringSet){
	n := getSmallestMultiplier(similarity) * ((rand.Int() % 1) + 1)
	x := int(float64(n) * similarity)
	a := rand.Int() % (n-x)
	b := n - x - a

	babbler := newBabble()
	core := babbler(x)
	first := core
	second := core
	if a > 0 {
		first += " " + babbler(a)
	}
	if b > 0 {
		second += " " + babbler(b)
	}
	return sentenceToSet(first), sentenceToSet(second)

}

func sentenceToSet(sentence string) jaccard.StringSet {
	set := jaccard.NewStringSet()
	for _, word := range strings.Split(sentence, " ") {
		set.Add(word)
	}
	return set
}

