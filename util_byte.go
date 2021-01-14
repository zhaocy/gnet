package gnet

import "encoding/binary"

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func BytesToUint16(n []byte) uint16 {
	return binary.BigEndian.Uint16(n)
}
