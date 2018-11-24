package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	byteOrder := binary.LittleEndian

	binary.Write(buf, byteOrder, uint32(92301))
	fmt.Printf("uint32: %x\n", buf.Bytes())

	// buf.Reset()
	binary.Write(buf, byteOrder, uint16(65535))
	fmt.Printf("uint16: %x\n", buf.Bytes())

	// buf.Reset()
	binary.Write(buf, byteOrder, float32(0.0012))
	fmt.Printf("float: %x\n", buf.Bytes())
}
