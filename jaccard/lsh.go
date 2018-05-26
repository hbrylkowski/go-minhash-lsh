package jaccard

import (
	"fmt"
	"strings"
	"github.com/OneOfOne/xxhash"
)

type LSH struct {
	minHash CalculateMinhash
	bucketsNumber uint64
	elementsInBucket uint64
	Buckets []map[uint64][]uint64
}

func (lsh LSH) GetDuplicateCandidates() [][]uint64 {
	var duplicates [][]uint64
	for _, bucket := range lsh.Buckets {
		for _, row := range bucket {
			if len(row) > 1 {
				duplicates = append(duplicates, row)
			}
		}
	}
	return duplicates
}

func (lsh *LSH) Add(s []string, identifier uint64)  {
	minHash := lsh.minHash(s)
	var b uint64
	for ; b < uint64(len(lsh.Buckets)); b++ {
		numbersHash := minHash[b * lsh.elementsInBucket: (b+1) * lsh.elementsInBucket]
		stringToHash := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbersHash)), ","), "[]")
		hash := xxhash.Checksum64([]byte(stringToHash))
		lsh.Buckets[b][hash] = append(lsh.Buckets[b][hash], identifier)
	}
}

type UintToIdentifiers struct {
	set map[uint64][]uint64
}

func NewLSH(bucketsCount uint64, elementsInBucket uint64, seed int64) LSH {
	buckets := make([]map[uint64][]uint64, bucketsCount)
	var b uint64
	for ; b < uint64(bucketsCount); b++ {
		buckets[b] = make(map[uint64][]uint64)
	}
	return LSH{
		minHash:          NewMinhash(bucketsCount* elementsInBucket, seed),
		bucketsNumber:    bucketsCount,
		elementsInBucket: elementsInBucket,
		Buckets:          buckets,
	}
}