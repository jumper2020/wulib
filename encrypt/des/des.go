package des

import (
	"encoding/binary"
	"errors"
)

//reference:
//http://www.hankcs.com/security/des-algorithm-illustrated.html
//http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Bdes/
//https://www.cnblogs.com/52fhy/p/6863133.html

type EncryptDes struct {
}

var pc1 []uint8
var pc2 []uint8
var ip []uint8
var E []uint8
var P []uint8
var ip_1 []uint8
var sdata [][]uint8
var k []uint64

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

	sdata = [][]uint8{
		{
			14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7,
			0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8,
			4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0,
			15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13,
		},
		{

			15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10,
			3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5,
			0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15,
			13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9,
		},
		{
			10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8,
			13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1,
			13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7,
			1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12,
		},
		{
			7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15,
			13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9,
			10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4,
			3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14,
		},
		{
			2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9,
			14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6,
			4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14,
			11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3,
		},
		{
			12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11,
			10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8,
			9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6,
			4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13,
		},
		{
			4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1,
			13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6,
			1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2,
			6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12,
		},
		{
			13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7,
			1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2,
			7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8,
			2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11,
		},
	}

}



func s(src uint64) uint64 {

	src6 := make([]uint8, 0)
	for i := 16; i < 64; i += 6 {
		src6 = append(src6, uint8((src<<uint(i))>>58))
	}

	rst := make([]uint8, 0)
	for k, v := range src6 {
		row := (v>>5)<<1 | (v & 0x01)
		col := ((v << 3) >> 4)
		s := sdata[k][row*16+col]
		rst = append(rst, s)
	}

	//need to transform four bit one group to eight bit one group
	r := make([]byte, 8)
	for i := 0; i < len(rst); i += 2 {
		tmp := (rst[i] << 4) | (rst[i+1])
		r[4+i/2] = tmp
	}

	return binary.BigEndian.Uint64(r)
}

func transform(src uint64, srcLen uint8, trans []uint8, destLen uint8) (block uint64) {
	srcPad := 64 - srcLen
	desPad := 64 - destLen
	for k, v := range trans {
		bit := ((src >> (64 - srcPad - v)) & 0x01)
		block |= bit << (64 - desPad - uint8(k+1))
	}
	return
}


//num is the num, bitLen is the actual bit length of num.
func rotateLeftShift(num uint64, bitLen uint8, shift uint8) uint64 {
	tmp := uint64(0x01)<<shift - 1
	tmpLow := uint64(tmp & (num >> (bitLen - shift)))
	tmpHigh := uint64(num & (^(tmp << (bitLen - shift))))
	rst := (tmpHigh << shift) | tmpLow
	return rst
}

func calculateSubKeys(key uint64) []uint64 {

	rst := make([]uint64, 0)

	shiftTable := []uint8{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	newKey := transform(key, 64, pc1, 56)

	c0 := newKey >> 28
	d0 := newKey & (uint64(0x01)<<28 - 1)
	tmpc := c0
	tmpd := d0

	for i := 0; i < 16; i++ {

		tmpc = rotateLeftShift(tmpc, 28, shiftTable[i])
		tmpd = rotateLeftShift(tmpd, 28, shiftTable[i])
		tmpK := uint64((tmpc << 28) | tmpd)

		ki := transform(tmpK, 56, pc2, 48)
		rst = append(rst, ki)
	}

	return rst
}

func calculateLoop(src uint64, decrypt bool) (uint32, uint32) {
	tmpIp := transform(src,64, ip, 64)
	l_1 := uint32(tmpIp>>32)
	r_1 := uint32(tmpIp)
	var l uint32
	var r uint32

	for i := 0; i < 16; i++ {
		l = r_1
		var frst uint32
		if !decrypt{
			frst = f(uint64(r_1), k[i])
		}else{
			frst = f(uint64(r_1), k[15-i])
		}
		r = l_1 ^ frst

		l_1 = l
		r_1 = r
	}

	return l, r
}

func f(srcR uint64, srcK uint64) uint32 {

	Er := transform(srcR, 32, E, 48)
	xRst := uint64(srcK ^ Er)
	c := s(xRst)
	rst := transform(c, 32,  P, 32)
	return uint32(rst)
}

func (self *EncryptDes) Encrypt(src []byte, key []byte) ([]byte, error) {
	if len(src) != 8 || len(key) != 8{
		return nil, errors.New("invalid param.")
	}

	keyNum := binary.BigEndian.Uint64(key)
	srcNum := binary.BigEndian.Uint64(src)

	k = calculateSubKeys(keyNum)
	l, r := calculateLoop(srcNum, false)
	bs := uint64(uint64(r)<<32 | uint64(l))
	rst := transform(bs, 64, ip_1, 64)
	block := make([]byte, 8)
	binary.BigEndian.PutUint64(block, rst)
	return block, nil
}

func (self *EncryptDes) Decrypt(src []byte, key []byte) ([]byte, error){

	keyNum := binary.BigEndian.Uint64(key)
	srcNum := binary.BigEndian.Uint64(src)

	k = calculateSubKeys(keyNum)
	l, r := calculateLoop(srcNum, true)
	bs := uint64(uint64(r)<<32 | uint64(l))
	rst := transform(bs, 64, ip_1, 64)
	block := make([]byte, 8)
	binary.BigEndian.PutUint64(block, rst)

	return block, nil
}
