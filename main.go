package main

import (
	"fmt"
	"os"
	"sync"

	"ttg/input"
	"ttg/trigrams"
)

func main() {
	var wg sync.WaitGroup

	// Either process stdin or file path inputs
	if input.IsPiped() {
		wg.Add(1)
		go trigrams.Produce(os.Stdin, &wg)
	} else {
		if len(os.Args[1:]) > 0 {
			// Make sure file paths exist, skipp not present ones
			files := input.File(os.Args[1:])

			for _, file := range files {
				f, err := os.Open(file)
				if err != nil {
					fmt.Printf("Failed to open '%v': %v\n", f, err)
				} else {
					wg.Add(1)
					go trigrams.Produce(f, &wg)
				}
				defer f.Close()
			}
		} else {
			fmt.Println("No input files")
		}
	}

	wg.Wait()
	fmt.Print("\nCompleted!\n\n")
}
