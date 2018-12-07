package md5

import (
	"bytes"
	"encoding/binary"
)

func padding(input []byte) (message []uint32) {
	K := len(input) * 8

	// 填充P位数据
	P := (448 - (K % 512) + 512) % 512
	if P == 0 {
		P = 512
	}
	input = append(input, byte(128))
	for i,times := 1,P/8; i < times; i++ {
		input = append(input, byte(0))
	}

	// 附加K值的低64位
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(K))
	input = append(input, b...)

	// 转成uint32数组
	for i,times := 0,len(input)/4; i < times; i++ {
		b := []byte{input[i*4], input[i*4+1], input[i*4+2], input[i*4+3]}
		bytesBuffer := bytes.NewBuffer(b)
		var x uint32
		binary.Read(bytesBuffer, binary.LittleEndian, &x)
		message = append(message, x)
	}

	return message
}