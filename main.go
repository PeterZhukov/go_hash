package main

import (
	"fmt"
	"go_hash/hashing"
	"go_hash/rehashing"
)

func main() {
	hashFunc := func(data []byte) uint32 {
		return hashing.HashMurMur3(data)
	}

	fmt.Println("==============Hashing==============")
	hash := make(hashing.Hash, hashing.NumOfBuckets)

	hashPhones(hash, hashFunc)
	fmt.Println("Phone Numbers:")
	printHashResult(hash)

	hash = make(hashing.Hash, hashing.NumOfBuckets)

	hashWords(hash, hashFunc)
	fmt.Println("Words:")
	printHashResult(hash)

	hash = make(hashing.Hash, hashing.NumOfBuckets)

	hashEmails(hash, hashFunc)
	fmt.Println("Emails:")
	printHashResult(hash)

	fmt.Println("==============Rehashing==============")
	rehash := make(rehashing.Rehash, hashing.NumOfBuckets)

	rehashPhones(rehash, hashFunc)
	fmt.Println("Phone Numbers:")
	printRehashResult(rehash)

	rehash = make(rehashing.Rehash, hashing.NumOfBuckets)

	rehashWords(rehash, hashFunc)
	fmt.Println("Words:")
	printRehashResult(rehash)

	rehash = make(rehashing.Rehash, hashing.NumOfBuckets)

	rehashEmails(rehash, hashFunc)
	fmt.Println("Emails:")
	printRehashResult(rehash)
}

func printHashResult(hash hashing.Hash) {
	res := make([]uint32, len(hash))
	var total uint64 = 0

	for k, v := range hash {
		res[k] = v
		total += uint64(v)
	}

	fmt.Println("total=", total)
	for k, v := range res {
		fmt.Println("bucket=", k, " nums=", v)
	}
	fmt.Println()
}

func printRehashResult(rehash rehashing.Rehash) {
	res := make([]rehashing.Total, len(rehash))

	for oldBucket, newHash := range rehash {
		var total uint64 = 0
		var totalStayed uint64 = 0
		var totalMoved uint64 = 0

		for newBucket, count := range newHash {
			if oldBucket == newBucket {
				totalStayed += uint64(count)
			} else {
				totalMoved += uint64(count)
			}

			total += uint64(count)
		}

		res[oldBucket] = rehashing.Total{
			TotalMoved:   totalMoved,
			TotalStayied: totalStayed,
			Total:        total,
		}
	}

	for oldBucket, totals := range res {
		fmt.Println("===================")
		fmt.Println("bucket=", oldBucket)
		fmt.Println("totalMoved=", totals.TotalMoved)
		fmt.Println("totalStayed=", totals.TotalStayied)
		fmt.Println("total=", totals.Total)
		fmt.Println("===================")
	}
}

func hashPhones(hash hashing.Hash, hashFunc func(data []byte) uint32) {
	for phone := range hashing.GeneratePhones() {
		phoneHash := hashFunc([]byte(fmt.Sprintf("%v", phone)))

		bucket := phoneHash % hashing.NumOfBuckets

		hash.Increase(bucket)
	}
}

func rehashPhones(rehash rehashing.Rehash, hashFunc func(data []byte) uint32) {
	for phone := range hashing.GeneratePhones() {
		phoneHash := hashFunc([]byte(fmt.Sprintf("%v", phone)))

		bucketOld := phoneHash % hashing.NumOfBuckets
		bucketNew := phoneHash % rehashing.NumOfBuckets

		rehash.Increase(bucketOld, bucketNew)
	}
}

func hashWords(hash hashing.Hash, hashFunc func(data []byte) uint32) {
	for word := range hashing.GenerateWords() {
		wordHash := hashFunc([]byte(word))

		bucket := wordHash % hashing.NumOfBuckets

		hash.Increase(bucket)
	}
}

func rehashWords(rehash rehashing.Rehash, hashFunc func(data []byte) uint32) {
	for word := range hashing.GenerateWords() {
		wordHash := hashFunc([]byte(word))

		bucketOld := wordHash % hashing.NumOfBuckets
		bucketNew := wordHash % rehashing.NumOfBuckets

		rehash.Increase(bucketOld, bucketNew)
	}
}

func hashEmails(hash hashing.Hash, hashFunc func(data []byte) uint32) {
	for email := range hashing.GenerateEmails() {
		emailHash := hashFunc([]byte(email))

		bucket := emailHash % hashing.NumOfBuckets

		hash.Increase(bucket)
	}
}

func rehashEmails(rehash rehashing.Rehash, hashFunc func(data []byte) uint32) {
	for email := range hashing.GenerateEmails() {
		emailHash := hashFunc([]byte(email))

		bucketOld := emailHash % hashing.NumOfBuckets
		bucketNew := emailHash % rehashing.NumOfBuckets

		rehash.Increase(bucketOld, bucketNew)
	}
}
