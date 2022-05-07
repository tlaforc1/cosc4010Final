// CODE IS FROM EITHER THE BLACK HAT GO BOOK OR BLACK HAT GO REPO WITH MINOR MODIFICATION

package pnginfo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"log"
	"models"
	"strconv"
	"util"
)

// Header of the PNG file
type Header struct {
	Header uint64
}

//Chunk = data byte chunk segment
type Chunk struct {
	Size uint32
	Type uint32
	Data []byte
	CRC  uint32
}

//MetaChunk = chunk with offset
type MetaChunk struct {
	Chk    Chunk
	Offset int64
}

//Validate if file is a PNG
func (mc *MetaChunk) validate(b *bytes.Reader) {
	var header Header
	if err := binary.Read(b, binary.BigEndian, &header.Header); err != nil {
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

//Read the sequence of chunks that makes up the file
func (mc *MetaChunk) ProcessImage(b *bytes.Reader, c *models.CmdLineOpts) {
	mc.validate(b)

	if (c.Offset != "") && (c.Encode == false && c.Decode == false) {
		var m MetaChunk
		m.Chk.Data = []byte(c.Payload)
		m.Chk.Type = m.strToInt(c.Type)
		m.Chk.Size = m.createChunkSize()
		m.Chk.CRC = m.createChunkCRC()

		bm := m.marshalData()
		bmb := bm.Bytes()

		fmt.Printf("Payload Original: % X\n", []byte(c.Payload))
		fmt.Printf("Payload: % X\n", m.Chk.Data)
		util.WriteData(b, c, bmb)
	}

	if (c.Offset != "") && c.Encode {
		var m MetaChunk
		m.Chk.Data = util.XorEncode([]byte(c.Payload), c.Key)
		m.Chk.Type = m.strToInt(c.Type)
		m.Chk.Size = m.createChunkSize()
		m.Chk.CRC = m.createChunkCRC()

		bm := m.marshalData()
		bmb := bm.Bytes()

		fmt.Printf("Payload Original: % X\n", []byte(c.Payload))
		fmt.Printf("Payload Encode: % X\n", m.Chk.Data)
		util.WriteData(b, c, bmb)
	}

	if (c.Offset != "") && c.Decode {
		var m MetaChunk
		offset, _ := strconv.ParseInt(c.Offset, 10, 64)
		b.Seek(offset, 0)

		m.readChunk(b)
		origData := m.Chk.Data
		m.Chk.Data = util.XorDecode(m.Chk.Data, c.Key)
		m.Chk.CRC = m.createChunkCRC()

		bm := m.marshalData()
		bmb := bm.Bytes()

		fmt.Printf("Payload Original: % X\n", origData)
		fmt.Printf("Payload Decode: % X\n", m.Chk.Data)
		util.WriteData(b, c, bmb)
	}

	count := 1
	chunkType := ""
	endChunkType := "IEND"

	// for chunkType != endChunkType {
	// 	var chk MetaChunk
	// 	fmt.Println("---- Chunk # " + strconv.Itoa(count) + " ----")
	// 	offset := chk.getOffset(b)
	// 	fmt.Printf("Chunk Offset: %#02x\n", offset)
	// 	chk.readChunk(b)
	// 	chunkType = chk.chunkTypeToString()
	// 	count++
	// }
}

//Get the offset of a chunk
func (mc *MetaChunk) getOffset(b *bytes.Reader) {
	offset, _ := b.Seek(0, 1)
	mc.Offset = offset
}

//Reads individual chunks
func (mc *MetaChunk) readChunk(b *bytes.Reader) {
	mc.readChunkSize(b)
	mc.readChunkType(b)
	mc.readChunkBytes(b, mc.Chk.Size)
	mc.readChunkCRC(b)
}

//Helper functions to read indvidual chunks
func (mc *MetaChunk) readChunkSize(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Size); err != nil {
		log.Fatal(err)
	}
}

func (mc *MetaChunk) readChunkType(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Type); err != nil {
		log.Fatal(err)
	}
}

func (mc *MetaChunk) readChunkBytes(b *bytes.Reader, cLen uint32) {
	mc.Chk.Data = make([]byte, cLen)
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Data); err != nil {
		log.Fatal(err)
	}
}

func (mc *MetaChunk) readChunkCRC(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.CRC); err != nil {
		log.Fatal(err)
	}
}

//Helper functions to create chunks
func (mc *MetaChunk) strToInt(s string) uint32 {
	t := []byte(s)
	return binary.BigEndian.Uint32(t)
}

func (mc *MetaChunk) createChunkSize() uint32 {
	return uint32(len(mc.Chk.Data))
}

func (mc *MetaChunk) createChunkCRC() uint32 {
	bytesMSB := new(bytes.Buffer)
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Type); err != nil {
		log.Fatal(err)
	}

	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Data); err != nil {
		log.Fatal(err)
	}

	return crc32.ChecksumIEEE(bytesMSB.Bytes())
}

func (mc *MetaChunk) marshalData() *bytes.Buffer {
	bytesMSB := new(bytes.Buffer)

	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Size); err != nil {
		log.Fatal(err)
	}

	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Type); err != nil {
		log.Fatal(err)
	}

	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Data); err != nil {
		log.Fatal(err)
	}

	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.CRC); err != nil {
		log.Fatal(err)
	}

	return bytesMSB
}
