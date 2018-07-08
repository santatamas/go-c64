package main

import (
	"bytes"
	"encoding/binary"
)

func toInt16(data []byte) (ret uint16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func toInt16_2(hi byte, lo byte) (ret uint16) {
	buf := bytes.NewBuffer([]byte{hi, lo})
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func getLO(value uint16) byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, value)
	return bs[0]
}

func getHI(value uint16) byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, value)
	return bs[1]
}
