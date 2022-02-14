// input provides distinction between the 2 ways data can be feed into the program: stdin and file paths.
package input

import (
	"fmt"
	"log"
	"os"
	"ttg/utils"
)

// IsPiped confirms wether or not the input received is from stdin. It returns true in case stdin is used,
// otherwise, it returns false.
func IsPiped() bool {
	fi, err := os.Stdin.Stat()

	if err != nil {
		log.Fatal(err)
	}

	if (fi.Mode() & os.ModeNamedPipe) != 0 {
		return true
	}

	return false
}

// File receives an array of strings containing the file paths that will be processed. It checks if those
// paths really exist, returning exclusively those which were confirmed to exist.
func File(args []string) []string {
	var files []string

	for _, v := range args {
		exists, err := utils.FileExists(v)

		if exists {
			files = append(files, v)
		} else {
			fmt.Printf("Skipping '%v': %v\n", v, err)
		}
	}
	return files
}
