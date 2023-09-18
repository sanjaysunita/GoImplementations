package main

import (
	"fmt"

	"math/rand"
	"hash/fnv"  
	"time"
)


// A go program that implements the BloomFilter algorithm


type BloomFilter struct {
	bitArray  []bool
	hashCount uint
	fpProb    float64
}

func NewBloomFilter(size uint, hashCount uint, fpProb float64) *BloomFilter {
	return &BloomFilter{
		bitArray:  make([]bool, size),
		hashCount: hashCount,
		fpProb:    fpProb,
	}
}

func (bf *BloomFilter) Add(item string) {
	for i := uint(0); i < bf.hashCount; i++ {
		hash := fnvHash(item + string(i))
		index := hash % uint32(len(bf.bitArray))
		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) Check(item string) bool {
	for i := uint(0); i < bf.hashCount; i++ {
		hash := fnvHash(item + string(i))
		index := hash % uint32(len(bf.bitArray))
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	size := uint(1000)
	hashCount := uint(5)
	fpProb := 0.05

	bf := NewBloomFilter(size, hashCount, fpProb)

	wordPresent := []string{
		"abound", "abounds", "abundance", "abundant", "accessible",
		"bloom", "blossom", "bolster", "bonny", "bonus", "bonuses",
		"coherent", "cohesive", "colorful", "comely", "comfort",
		"gems", "generosity", "generous", "generously", "genial",
	}

	wordAbsent := []string{
		"bluff", "cheater", "hate", "war", "humanity",
		"racism", "hurt", "nuke", "gloomy", "facebook",
		"geeksforgeeks", "twitter",
	}

	for _, item := range wordPresent {
		bf.Add(item)
	}

	rand.Shuffle(len(wordPresent), func(i, j int) {
		wordPresent[i], wordPresent[j] = wordPresent[j], wordPresent[i]
	})

	rand.Shuffle(len(wordAbsent), func(i, j int) {
		wordAbsent[i], wordAbsent[j] = wordAbsent[j], wordAbsent[i]
	})

	testWords := append(wordPresent[:10], wordAbsent...)
	rand.Shuffle(len(testWords), func(i, j int) {
		testWords[i], testWords[j] = testWords[j], testWords[i]
	})

	for _, word := range testWords {
		if bf.Check(word) {
			if contains(wordAbsent, word) {
				fmt.Printf("'%s' is a false positive!\n", word)
			} else {
				fmt.Printf("'%s' is probably present!\n", word)
			}
		} else {
			fmt.Printf("'%s' is definitely not present!\n", word)
		}
	}
}

func fnvHash(s string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hash.Sum32()
}

func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
