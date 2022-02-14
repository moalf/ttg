package trigrams

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
)

func CleanUp(b *[]byte) []string {
	// Ignore punctuation
	punctuationRegex := regexp.MustCompile(`(?i)[’"”“_;:,!'@#%&<>\(\)\+\{\}\[\]\*\.\?]`)
	tempData := punctuationRegex.ReplaceAllString(string(*b), "")

	// Replace the following with 1 space: a) 2 or more consecutive white spaces, b) 1 hyphen,
	// c) new line, d) carriage return
	multipleSpacesRegex := regexp.MustCompile(`\s\s+|-|\r|\n`)
	tempData = multipleSpacesRegex.ReplaceAllString(tempData, " ")

	return strings.Split(tempData, " ")
}

func Produce(f *os.File, wg *sync.WaitGroup) {
	defer wg.Done()

	r := bufio.NewReader(f)
	b := make([]byte, 0, 4*1024)

	trigrams := make(map[string]int)

	for {
		n, err := r.Read(b[:cap(b)])
		b = b[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
		}

		data := CleanUp(&b)

		var key string
		l := len(data)

		// Iterate over the words data to collect/generate the trigrams
		for i := 0; i < l-2; i++ {
			// Generate trigrams, e.g.: "the white whale"
			key = fmt.Sprintf("%s %s %s", data[i], data[i+1], data[i+2])
			key = strings.ToLower(key)
			key = strings.TrimSpace(key)

			// Add the trigram if non-existent, increase its count otherwise
			if v, ok := trigrams[key]; ok {
				trigrams[key] = v + 1
			} else {
				trigrams[key] = 1
			}
		}
	}

	sequences := Sort(&trigrams)

	most := 100
	fmt.Printf("\nFile: %s\n", f.Name())
	for i, s := range sequences {
		if i == most {
			break
		}
		fmt.Printf("%3d: %q - %d\n", i+1, s.key, s.val)
	}
}
