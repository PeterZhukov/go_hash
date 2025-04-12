package hashing

import (
	"bufio"
	"fmt"
	"github.com/spaolacci/murmur3"
	"hash/crc32"
	"iter"
	"math/rand"
	"os"
)

type Hash map[uint32]uint32

func (h Hash) Increase(key uint32) {
	h[key]++
}

var NumOfBuckets uint32 = 10

func HashCrc32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func HashMurMur3(data []byte) uint32 {
	return murmur3.Sum32(data)
}

func GeneratePhones() iter.Seq[uint64] {
	return func(yield func(uint64) bool) {
		phonesPrefixes := []uint64{7903, 7905, 7910, 7916, 7926}

		for _, prefix := range phonesPrefixes {
			for i := 0; i <= 999_99_99; i++ {
				num := prefix*uint64(10000000) + uint64(i)

				if !yield(num) {
					return
				}
			}
		}
	}
}

func GenerateWords() iter.Seq[string] {
	return func(yield func(string) bool) {
		f, err := os.Open("words_alpha.txt")

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()

			if !yield(line) {
				return
			}
		}
	}
}

func GenerateEmails() iter.Seq[string] {
	return func(yield func(string) bool) {
		f, err := os.Open("words_alpha.txt")

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		defer f.Close()

		lines1 := make([]string, 0)
		lines2 := make([]string, 0)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()

			lines1 = append(lines1, line)
			lines2 = append(lines2, line)
		}

		rand.Shuffle(len(lines1), func(i, j int) { lines1[i], lines1[j] = lines1[j], lines1[i] })
		rand.Shuffle(len(lines2), func(i, j int) { lines2[i], lines2[j] = lines2[j], lines2[i] })

		lines := make([]string, 3*len(lines1))

		for i := range lines1 {
			lines = append(lines, lines1[i]+lines2[i])
			lines = append(lines, lines1[i]+"_"+lines2[i])
			lines = append(lines, lines1[i]+"."+lines2[i])
		}

		emailSuffixes := []string{"@gmail.com", "@yandex.ru", "@mail.ru"}

		for _, line := range lines {

			for _, suffix := range emailSuffixes {
				currLine := line
				currLine += suffix

				if !yield(currLine) {
					return
				}
			}
		}
	}
}
