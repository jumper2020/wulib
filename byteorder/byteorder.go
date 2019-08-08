package byteorder

import (
	"fmt"
	"unsafe"
)

//1. 如下一些需要明白
//	0x12345678 中左边为高位，右边为低位
//	地址 0xa1 0xa2 0xa3 0xa4 从低址到高址
//  big endian: 低址存高位，高址存低位
//	0xa1 - 0x12  0xa2 - 0x34  0xa3 - 0x56  0xa4 - 0x78
//	little endian: 低址存低位，高址存高位
//	0xa1 - 0x78  0xa2 - 0x56  0xa3 - 0x34  0xa4 - 0x12
//https://itimetraveler.github.io/2018/01/18/%E5%A6%82%E4%BD%95%E5%88%A4%E6%96%ADCPU%E6%98%AF%E5%A4%A7%E7%AB%AF%E8%BF%98%E6%98%AF%E5%B0%8F%E7%AB%AF%E6%A8%A1%E5%BC%8F/
//http://www.ruanyifeng.com/blog/2016/11/byte-order.html

//2. 注意:
//大小端是指在内存中的存储方式，在进行运算的时候（左右移或者类型转换）都是将内存中的数据读入到寄存器中进行运算的。
//此时无论内存中是大端存储还是小端存储，并不会影响寄存器中的处理过程，
//比如x>>8 内存中是大端，读入寄存器中直接进行>>8, 然后再以大端方式存入到内存
//内存中是小端，读入寄存器中转为大端, >>8, 再转为小端存入到内存
//总之，大小端为内存中的存储方式，不会影响运算。应该可以认为无论大小端，在执行运算的时候都是同一个形式，比如 0x1234
//参见
//https://stackoverflow.com/questions/7184789/does-bit-shift-depend-on-endianness
//https://bbs.csdn.net/topics/390567539
//因此 对于疑问 https://stackoverflow.com/questions/7184789/does-bit-shift-depend-on-endianness
//其中以内存中的存储形式来执行 >> 是错误的。
//byte(xxx) 是在寄存器中执行运算， 都是获取的低位字节。
//因此 >> << byte() 等这些运算都是在寄存器中执行，不受大小端影响的。

//要想判断内存中是大端还是小端方式，应该使用 unsafe.Pointer, 如下 checkEndian 函数
//参考： https://stackoverflow.com/questions/51332658/any-better-way-to-check-endianness-in-go


//function:
//0. check big endian / little endian
//1. big endian: uint - []byte
//2. big endian: []byte - uint
//3. little endian: uint - []byte
//4. little endian: []byte - uint


//return true means big endian, false little endian
func checkEndian() bool {
	test16 := uint16(0x1234)
	test8 := *(*byte)(unsafe.Pointer(&test16))
	if test8 == 0x12{
		fmt.Printf("big")
		return true
	}else{
		fmt.Printf("little")
		return false
	}
}

func bigU16ToBS(src uint16) []byte {
	rst := make([]byte, 2)
	rst[0] = byte(src>>8)
	rst[1] = byte(src)
	return rst
}


func bigU32ToBS(src uint32) []byte {
	rst := make([]byte, 4)
	rst[0] = byte(src>>24)
	rst[1] = byte(src>>16)
	rst[2] = byte(src>>8)
	rst[3] = byte(src)

	return rst
}


func bigU64ToBS(src uint64) []byte {
	rst := make([]byte, 8)
	rst[0] = byte(src>>56)
	rst[1] = byte(src>>48)
	rst[2] = byte(src>>40)
	rst[3] = byte(src>>32)
	rst[4] = byte(src>>24)
	rst[5] = byte(src>>16)
	rst[6] = byte(src>>8)
	rst[7] = byte(src)

	return rst
}

func bigBSToU16(src []byte) uint16{
	return uint16( uint16(src[0])<<8 | uint16(src[1]))
}

func bigBSToU32(src []byte) uint32{
	return uint32( uint32(src[0]) << 24 | uint32(src[1]) << 16 | uint32(src[2]) << 8 | uint32(src[3]))
}

func bigBSToU64(src []byte) uint64{
	return uint64( uint64(src[0]) << 56 | uint64(src[1])<<48 | uint64(src[2])<<40 | uint64(src[3])<<32 |
		uint64(src[4])<<24 | uint64(src[5])<<16 | uint64(src[6])<<8 | uint64(src[7]))
}

func littleU16ToBS(src uint16) []byte{
	rst := make([]byte, 2)
	rst[0] = byte(src)
	rst[1] = byte(src>>8)
	return rst
}

func littleU32ToBS(src uint32) []byte{
	rst := make([]byte, 4)
	rst[0] = byte(src)
	rst[1] = byte(src>>8)
	rst[2] = byte(src>>16)
	rst[3] = byte(src>>24)

	return rst
}

func littleU64ToBS(src uint64) []byte{
	rst := make([]byte, 8)
	rst[0] = byte(src)
	rst[1] = byte(src>>8)
	rst[2] = byte(src>>16)
	rst[3] = byte(src>>24)
	rst[4] = byte(src>>32)
	rst[5] = byte(src>>40)
	rst[6] = byte(src>>48)
	rst[7] = byte(src>>56)

	return rst
}

func littleBSToU16(src []byte) uint16{
	return  uint16(src[0]) | uint16(src[1])<<8
}

func littleBSToU32(src []byte) uint32{
	return uint32(src[0]) | uint32(src[1])<<8 | uint32(src[2])<<16 | uint32(src[3])<<24
}

func littleBSToU64(src []byte) uint64{
	return uint64(src[0]) | uint64(src[1])<<8 | uint64(src[2])<<16 | uint64(src[3])<<24 |
		uint64(src[4])<<32 | uint64(src[5])<<40 | uint64(src[6])<<48 | uint64(src[7])<<56
}
