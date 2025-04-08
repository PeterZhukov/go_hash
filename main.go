package main

import (
	"bufio"
	"fmt"
	"github.com/spaolacci/murmur3"
	//"hash/crc32"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	phonesPrefixes := []int{7903, 7905, 7910, 7916, 7926}
	numBuckets := uint32(10)

	total := 0
	m := make(map[uint32]uint32, numBuckets)

	phonesPrefixes = phonesPrefixes
	checkPhones(phonesPrefixes, m, numBuckets, &total)
	fmt.Println("Phone Numbers:")
	printResult(m, total, numBuckets)

	total = 0
	m = make(map[uint32]uint32, numBuckets)

	checkWords(m, numBuckets, &total)
	fmt.Println("Words:")
	printResult(m, total, numBuckets)

	checkEmails(m, numBuckets, &total)
	fmt.Println("Emails:")
	printResult(m, total, numBuckets)
}

func hash(data []byte) uint32 {
	//return crc32.ChecksumIEEE(data)
	return murmur3.Sum32(data)
}

func printResult(m map[uint32]uint32, total int, numBuckets uint32) {
	res := make([]uint32, numBuckets)
	for k, v := range m {
		res[k] = v
	}

	fmt.Println("total=", total)
	for k, v := range res {
		fmt.Println("bucket=", k, " nums=", v)
	}
	fmt.Println()
}

func checkPhones(phonesPrefixes []int, m map[uint32]uint32, numBuckets uint32, total *int) {
	for _, prefix := range phonesPrefixes {
		for i := 0; i <= 999_99_99; i++ {
			num := prefix*10000000 + i

			hash := hash([]byte(strconv.Itoa(num)))

			bucket := hash % numBuckets

			m[bucket]++
			*total++
		}
	}
}

func checkWords(m map[uint32]uint32, numBuckets uint32, total *int) {
	f, err := os.Open("words_alpha.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		hash := hash([]byte(line))

		bucket := hash % numBuckets

		m[bucket]++
		*total++
	}
}

func checkEmails(m map[uint32]uint32, numBuckets uint32, total *int) {
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

	lines := make([]string, 0)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines1), func(i, j int) { lines1[i], lines1[j] = lines1[j], lines1[i] })
	rand.Shuffle(len(lines2), func(i, j int) { lines2[i], lines2[j] = lines2[j], lines2[i] })

	for i := range lines1 {
		lines = append(lines, lines1[i]+lines2[i])
		lines = append(lines, lines1[i]+"_"+lines2[i])
		lines = append(lines, lines1[i]+"."+lines2[i])
	}

	emailSuffixes := []string{"@gmail.com", "@yandex.ru", "@mail.ru"}

	for _, line := range lines {

		for _, suffix := range emailSuffixes {
			line += suffix

			hash := hash([]byte(line))

			bucket := hash % numBuckets

			m[bucket]++
			*total++
		}
	}
}
