package jaccard

import (
	"math/rand"
	"math"
	"github.com/OneOfOne/xxhash"
)

type CalculateMinhash func(words []string) []uint64

func NewMinhash(elements uint64, seed int64) CalculateMinhash {
	var randomNumbers []uint64
	rand.Seed(seed)
	var i uint64
	for ; i < elements; i++ {
		randomNumbers = append(randomNumbers, rand.Uint64())
	}
	return func(words []string) []uint64{
		var hashes []uint64
		for _, r := range randomNumbers {
			hashes = append(hashes, calculateMinimalHash(words, r))
		}
		return hashes
	}
}

func calculateMinimalHash(words []string, randomNumber uint64) uint64{
	var minimal uint64 = math.MaxUint64
	for _, w := range words {
		hash := xxhash.Checksum64([]byte(w)) ^ randomNumber
		if minimal > hash {
			minimal = hash
		}
	}
	return minimal
}
