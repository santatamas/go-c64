package numeric

import (
	"bytes"
	"encoding/binary"
)

func ToInt16(data []byte) (ret uint16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func ToInt16_2(hi byte, lo byte) (ret uint16) {
	buf := bytes.NewBuffer([]byte{hi, lo})
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func GetLO(value uint16) byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, value)
	return bs[0]
}

func GetHI(value uint16) byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, value)
	return bs[1]
}

func GetBit(value byte, bitNumber byte) bool {
	return (value & (1 << bitNumber)) == (1 << bitNumber)
}
