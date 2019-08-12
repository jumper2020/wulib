package des

import (
	"encoding/binary"
	"fmt"
)

type EncryptDes struct {
}

var pc1 []uint8
var pc2 []uint8
var ip []uint8
var E []uint8
var P []uint8
var ip_1 []uint8

func init() {

	//64位的秘钥首先根据表格PC-1进行变换
	pc1 = []uint8{
		57, 49, 41, 33, 25, 17, 9,
		1, 58, 50, 42, 34, 26, 18,
		10, 2, 59, 51, 43, 35, 27,
		19, 11, 3, 60, 52, 44, 36,
		63, 55, 47, 39, 31, 23, 15,
		7, 62, 54, 46, 38, 30, 22,
		14, 6, 61, 53, 45, 37, 29,
		21, 13, 5, 28, 20, 12, 4,
	}

	pc2 = []uint8{
		14, 17, 11, 24, 1, 5,
		3, 28, 15, 6, 21, 10,
		23, 19, 12, 4, 26, 8,
		16, 7, 27, 20, 13, 2,
		41, 52, 31, 37, 47, 55,
		30, 40, 51, 45, 33, 48,
		44, 49, 39, 56, 34, 53,
		46, 42, 50, 36, 29, 32,
	}

	ip = []uint8{
		58, 50, 42, 34, 26, 18, 10, 2,
		60, 52, 44, 36, 28, 20, 12, 4,
		62, 54, 46, 38, 30, 22, 14, 6,
		64, 56, 48, 40, 32, 24, 16, 8,
		57, 49, 41, 33, 25, 17, 9, 1,
		59, 51, 43, 35, 27, 19, 11, 3,
		61, 53, 45, 37, 29, 21, 13, 5,
		63, 55, 47, 39, 31, 23, 15, 7,
	}

	E = []uint8{
		32, 1, 2, 3, 4, 5,
		4, 5, 6, 7, 8, 9,
		8, 9, 10, 11, 12, 13,
		12, 13, 14, 15, 16, 17,
		16, 17, 18, 19, 20, 21,
		20, 21, 22, 23, 24, 25,
		24, 25, 26, 27, 28, 29,
		28, 29, 30, 31, 32, 1,
	}

	P = []uint8{
		16, 7, 20, 21,
		29, 12, 28, 17,
		1, 15, 23, 26,
		5, 18, 31, 10,
		2, 8, 24, 14,
		32, 27, 3, 9,
		19, 13, 30, 6,
		22, 11, 4, 25,
	}

	ip_1 = []uint8{
		40, 8, 48, 16, 56, 24, 64, 32,
		39, 7, 47, 15, 55, 23, 63, 31,
		38, 6, 46, 14, 54, 22, 62, 30,
		37, 5, 45, 13, 53, 21, 61, 29,
		36, 4, 44, 12, 52, 20, 60, 28,
		35, 3, 43, 11, 51, 19, 59, 27,
		34, 2, 42, 10, 50, 18, 58, 26,
		33, 1, 41, 9, 49, 17, 57, 25,
	}
}

func transform(src []byte, trans []uint8) []byte {

	rstLen := len(trans)
	if rstLen%8 != 0 {
		return nil
	}
	rst := make([]byte, rstLen/8)

	for k, v := range trans {
		index := v - 1
		sIndex := index / 8
		sShift := index % 8
		fmt.Printf("v: %d\n", v)
		tmp := src[sIndex] & (0x01 << (7 - sShift))
		if tmp != 0 {
			i := k
			kIndex := i / 8
			kShift := i % 8
			rst[kIndex] |= uint8(0x01 << uint8(7-kShift))
		}
	}
	return rst
}

//num is the num, bitLen is the actual bit length of num.
func rotateLeftShift(num uint64, bitLen uint8, shift uint8) uint64 {
	tmp := uint64(0x01)<<shift - 1
	tmpLow := uint64(tmp & (num >> (bitLen - shift)))
	tmpHigh := uint64(num & (^(tmp << (bitLen - shift))))
	rst := (tmpHigh << shift) | tmpLow
	return rst
}

func calculateSubKeys(key []byte) [][]byte {

	//if len(key) != 8 {
	//	return nil
	//}

	rst := make([][]byte, 0)

	shiftTable := []uint8{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	newKey := transform(key, pc1)
	tmpKey := make([]byte, 0)
	tmpKey = append(tmpKey, byte(0x00))
	tmpKey = append(tmpKey, newKey...)


	fmt.Printf("tmpKey: %v\n", tmpKey)
	tmp := binary.BigEndian.Uint64(tmpKey)
	c0 := tmp >> 28
	d0 := tmp & (uint64(0x01)<<28 - 1)
	tmpc := c0
	tmpd := d0

	for i := 0; i < 16; i++ {

		tmpc = rotateLeftShift(tmpc, 28, shiftTable[i])
		tmpd = rotateLeftShift(tmpd, 28, shiftTable[i])

		tmpK := make([]byte, 8)
		binary.BigEndian.PutUint64(tmpK, uint64(uint64(tmpc)<<28|uint64(tmpd)))
		tmpK = tmpK[:len(tmpK)-1]
		fmt.Printf("tmpK: %v\n",tmpK)

		ki := transform(tmpK, pc2)
		rst = append(rst, ki)
		fmt.Printf("i: %d, rst: %v\n", i, rst)
		//? todo: 是否需要删除前面多余的8个0
	}

	return rst
}

//func (self *EncryptDes) Encrypt(src []byte, key []byte) ([]byte, error){
//
//}
//
//func (self *EncryptDes) Decrypt(src []byte, key []byte) ([]byte, error){
//
//}
