package util

// CODE IS FROM EITHER THE BLACK HAT GO BOOK OR BLACK HAT GO REPO WITH MINOR MODIFICATION

import(
	"bytes"
	"fmt"
	"io"
	"log"
	"models"
	"os"
	"strconv"
)

//Writes data to a particular offset
func WriteData(r *bytes.Reader, c *models.CmdLineOpts, b []byte) {
	offset, err := strconv.ParseInt(c.Offset, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	w, err := os.Create(c.Output)
	if err != nil {
		log.Fatal("Fatal: Issue writing to the output file")
	}
	r.Seek(0, 0)

	var buff = make([]byte, offset)
	r.Read(buff)
	w.Write(buff)
	w.Write(b)

	if c.Decode {
		r.Seek(int64(len(b)), 1)
	}

	_, err = io.Copy(w, r)
	if err == nil {
		fmt.Printf("Success: %s created \n", c.Output)
	}
}