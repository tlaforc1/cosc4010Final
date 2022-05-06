// CODE IS FROM EITHER THE BLACK HAT GO BOOK OR BLACK HAT GO REPO WITH MINOR MODIFICATION

// why...
package util

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

//Reads to buffer from file
func PreProcessImage(dat *os.File) (*bytes.Reader, error) {
	stats, err := dat.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var size = stats.Size()
	b := make([]byte, size)

	bufR := bufio.NewReader(dat)

	if _, err := bufR.Read(b); err != nil {
		log.Fatal(err)
	}

	bReader := bytes.NewReader(b)

	return bReader, err
}
