package gnet

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Uint16(b []byte) uint16 {
	_ = b[1]
	return uint16(b[0]) | uint16(b[1])<<8
}

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func BytesToUint16(n []byte) uint16 {
	return binary.BigEndian.Uint16(n)
}

func BytesToUint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func Uint64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |        uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

//遇到0字节结束
func ByteStringToString(bytes []byte) string {
	for i := 0; i < len(bytes); i = i + 1 {
		if bytes[i] == 0 {
			return string(bytes[0:i])
		}
	}
	return string(bytes)
}

//字符串连接
func BytesConnect(b1,b2 []byte)[]byte{
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write(b1)
	buffer.Write(b2)
	return buffer.Bytes()
}

//字符串连接
func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

//isSymbol表示有无符号
func BytesToInt(b []byte, isSymbol bool)  (int, error){
	if isSymbol {
		return bytesToIntS(b)
	}
	return bytesToIntU(b)
}


//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0},b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0,fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}



//字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0},b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0,fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}


//整形转换成字节
func IntToBytes(n int,b byte) ([]byte,error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(),nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(),nil
	case 3,4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(),nil
	}
	return nil,fmt.Errorf("IntToBytes b param is invaild")
}

//解析C2S的Head byte
func HeadByteC2S(byte []byte) []byte{
	cmdLenByte:= byte[:4]
	cmdLen := BytesToUint32(cmdLenByte)
	headLen:= 4+cmdLen+2  //4表示cmd len uint32的字节长度  2表示bodyDataLen uint16
	return byte[:headLen]
}