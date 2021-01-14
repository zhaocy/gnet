package gnet

import (
	"bytes"
	"encoding/binary"
)

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func BytesToUint16(n []byte) uint16 {
	return binary.BigEndian.Uint16(n)
}


func BytesConnect(b1,b2 []byte)[]byte{
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write(b1)
	buffer.Write(b2)
	return buffer.Bytes()  //得到了b1+b2的结果
}