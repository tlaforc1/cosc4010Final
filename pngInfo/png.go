// CODE IS FROM EITHER THE BLACK HAT GO BOOK OR BLACK HAT GO REPO WITH MINOR MODIFICATION

import (
	"log"
	"bytes"
	"enconding/binary"
	"strconv"
	"strings"
)

// Header of the PNG file
type Header struct {
	Header uint64
}

//Chunk = data byte chunk segment
type Chunk strut {
	Size uint32
	Type uint32
	Data []byte
	CRC uint32
}

//Validate if file is a PNG
func (mc *MetaChunk) validate(b *bytes.Reader) {
	var header Header
	if err := binary.Read(b, binaryBigEndian, &header.Header); err != nil {
		log.Fatal(err)
	}

	bArr := make([]byte, 8)
	binary.BigEndian.PutUint64(bArr, header.Header)

	if string(bArr[1:4]) != "PNG" {
		log.Fatal("Provided file is not a valid PNG format")
	} else {
		fmt.Println("Valid PNG")
	}
}

//Read chunk
func (mc *MetaChunk) ProcessIMage(b *bytes.Reader, c *models.CmdLineOpts) {
	count := 1
	chunkType := ""
	endChunkType := "IEND"
	
	for chunkType != endChunkType {
		fmt.Println("---- Chunk # " + strconv.Itoa(count) + " ----")
	}
}