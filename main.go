package main

import (
	"strings"
	"github.com/hbrylkowski/lsh/jaccard"
	"github.com/tjarratt/babble"
	"math/rand"
	"github.com/hbrylkowski/lsh/minhash"
	"time"
	"fmt"
	"math"
)
const minHashSize = 1000
var minHash = minhash.NewMinhash(minHashSize, time.Now().UTC().UnixNano())

func main() {
	var diff float64 = 0
	tries := 1000
	for i := 0; i < tries; i++ {
		first, second := createSentencesPair()
		diff += math.Abs(jaccard.JaccardsIndex(first, second) - multipleMinHashes(first, second))
	}
	fmt.Println(float64(diff) / float64(tries))

}
func createSentencesPair() (jaccard.StringSet, jaccard.StringSet){
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = rand.Int() % 90 + 10
	first := babbler.Babble()
	babbler.Count += rand.Int() % 20
	second := first + " " + babbler.Babble()
	return sentenceToSet(first), sentenceToSet(second)

}

func sentenceToSet(sentence string) jaccard.StringSet {
	set := jaccard.NewStringSet()
	for _, word := range strings.Split(sentence, " ") {
		set.Add(word)
	}
	return set
}

func multipleMinHashes(sentenceA jaccard.StringSet, sentenceB jaccard.StringSet) float64{
	commonMinHashElements := [minHashSize + 1]int64{}
	for i := 0; i < 10; i++{
		minHashA := minHash(sentenceA.Elements())
		minHashB := minHash(sentenceB.Elements())
		agreed := 0
		for i, h := range minHashA {
			if h == minHashB[i]{
				agreed++
			}
		}
		commonMinHashElements[agreed]++
	}
	var sum float64 = 0
	for i, s := range commonMinHashElements {
		sum += float64(i)/float64(minHashSize) * float64(s)
	}
	return sum / float64(10)
}
