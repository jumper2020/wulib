package main

import (
	"encoding/binary"
	"fmt"
)

//1. 如下一些需要明白
//	0x12345678 中左边为高位，右边为低位
//	地址 0xa1 0xa2 0xa3 0xa4 从低址到高址
//  big endian: 低址存高位，高址存低位
//	0xa1 - 0x12  0xa2 - 0x34  0xa3 - 0x56  0xa4 - 0x78
//	little endian: 低址存低位，高址存高位
//	0xa1 - 0x78  0xa2 - 0x56  0xa3 - 0x34  0xa4 - 0x12

//2. 将大变量赋值给小变量时，数据截断是从低址开始的
//	这一点可以用于判断cpu大端小端, 见 checkEndian


//0. check big endian / little endian
//1. big endian: uint - []byte
//2. big endian: []byte - uint
//3. little endian: uint - []byte
//4. little endian: []byte - uint

//read package binary

func main() {
	x := uint16(0x1000)

	fmt.Printf("x: %x, uint(x): %x\n", x, uint8(x))
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, x)
	fmt.Printf("b0: %v, b1: %v", b[0], b[1])
}

//return true means big endian, false little endian
func checkEndian() bool {
	test32 := uint32(0x12345678)
	test8 := uint8(test32)
	if test8 == 0x12{
		return true
	}else{
		return false
	}
}

func bigU8ToBS(src uint8) byte {
	return byte(src)
}

func bigU16ToBS(src uint16) []byte {
	rst := make([]byte, 2)
	rst[0] = byte(src>>8)
	rst[1] = byte(src)
	return rst
}
