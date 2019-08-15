package aes

import "fmt"

//http://www.alonemonkey.com/2016/05/25/aes-and-des/

type EncryptInterf interface {
	Encrypt(src []byte, key []byte) ([]byte, error)
	Decrypt(src []byte, key []byte) ([]byte, error)
}
type EncryptAes struct {
}

var keyScheduleConst [][]uint8
var subBytesConst [256]byte
var keys [][]byte
var mixConst []uint8

func init() {

	keyScheduleConst = [][]uint8{
		[]uint8{0x01, 0x00, 0x00, 0x00},
		[]uint8{0x02, 0x00, 0x00, 0x00},
		[]uint8{0x04, 0x00, 0x00, 0x00},
		[]uint8{0x08, 0x00, 0x00, 0x00},
		[]uint8{0x10, 0x00, 0x00, 0x00},
		[]uint8{0x20, 0x00, 0x00, 0x00},
		[]uint8{0x40, 0x00, 0x00, 0x00},
		[]uint8{0x80, 0x00, 0x00, 0x00},
		[]uint8{0x1b, 0x00, 0x00, 0x00},
		[]uint8{0x36, 0x00, 0x00, 0x00},
	}

	subBytesConst = [256]byte{
		0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
		0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
		0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
		0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75,
		0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84,
		0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf,
		0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8,
		0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2,
		0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73,
		0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb,
		0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79,
		0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08,
		0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a,
		0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e,
		0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf,
		0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16,
	}

	keys = make([][]byte, 10)

	mixConst = []uint8{0x02,0x03,0x01,0x01,0x01,0x02,0x03,0x01,0x01,0x01,0x02,0x03,0x03,0x01,0x01,0x02}
}

func xor(src1 []byte, src2 []byte) []byte {
	if len(src1) != len(src2) {
		return nil
	}

	rst := make([]byte, 0, len(src1))
	for k, v := range src1 {
		rst = append(rst, v^src2[k])
	}

	fmt.Printf("xor: %x\n", rst)
	return rst
}

func subBytes(src []byte) []byte {
	rst := make([]byte, 0, 16)
	for _, v := range src {
		rst = append(rst, subBytesConst[v])
	}
	return rst
}

func keySchedule(src []byte, ksi int) []byte {
	if len(src) != 16 {
		return nil
	}

	col1 := []byte{src[0], src[1*4], src[2*4], src[3*4]}
	col2 := []byte{src[0+1], src[1*4+1], src[2*4+1], src[3*4+1]}
	col3 := []byte{src[0+2], src[1*4+2], src[2*4+2], src[3*4+2]}
	col4 := []byte{src[0+3], src[1*4+3], src[2*4+3], src[3*4+3]}

	tmpCol4 := []byte{src[1*4+3], src[2*4+3], src[3*4+3], src[0+3]}

	//col1 := make([]byte, 4)
	//copy(col1, src[:4])
	//col2 := make([]byte, 4)
	//copy(col2, src[4:8])
	//col3 := make([]byte, 4)
	//copy(col3, src[8:12])
	//col4 := make([]byte, 4)
	//copy(col4, src[12:])
	//tmpCol4 := []byte{col4[3],col4[0],col4[1],col4[2]}

	tmpCol4 = subBytes(tmpCol4)

	fmt.Printf("col: %x, %x, %x, %x\n", col1, col2, col3, col4)
	fmt.Printf("tmpCol4: %x\n", tmpCol4)
	newCol1 := xor(xor(col1, tmpCol4), keyScheduleConst[ksi])
	newCol2 := xor(newCol1, col2)
	newCol3 := xor(newCol2, col3)
	newCol4 := xor(newCol3, col4)

	rst := make([]byte, 16)
	for i := 0; i < 16; i += 4 {
		rst[i] = newCol1[i/4]
		rst[i+1] = newCol2[i/4]
		rst[i+2] = newCol3[i/4]
		rst[i+3] = newCol4[i/4]
	}

	//rst := make([]byte, 0)
	//rst = append(rst, newCol1...)
	//rst = append(rst, newCol2...)
	//rst = append(rst, newCol3...)
	//rst = append(rst, newCol4...)

	return rst
}

func getKeys(src []byte) {

	tmp := make([]byte, len(src))
	copy(tmp, src)

	for i := 0; i < 10; i++ {
		keys[i] = keySchedule(tmp, i)
		tmp = keys[i]
	}
	return
}

func shiftrow(src []byte, shift uint) []byte {
	if len(src) != 4 {
		return nil
	}

	rst := make([]byte, 0)
	i := shift
	count := 0
	for count < len(src) {
		rst = append(rst, src[i])
		i = (i + 1) % uint(len(src))
		count++
	}
	return rst
}

func shiftRows(src []byte) []byte {
	if len(src) != 16 {
		return nil
	}

	rst := make([]byte, 0)
	for i := 0; i < len(src); i += 4 {
		rst = append(rst, shiftrow(src[i:i+4], uint(i/4))...)
	}

	return rst
}

func multiSingle(src1 uint8, src2 uint8) uint8 {
	var rst uint8
	switch src2 {
	case 1:
		rst = src1
	case 2:
		rst = src1 << 1
		if (src1 & 0x80) != 0 {
			rst = rst ^ 0x1b
		}
	case 3:
		rst = src1 << 1
		if (src1 & 0x80) != 0 {
			rst = rst ^ 0x1b
		}
		rst ^= src1
	}
	return rst
}

func multiGroup(src1 []uint8, src2 []uint8) []uint8 {
	if len(src1) != 4 || len(src2) != 16 {
		return nil
	}

	rst := make([]uint8, 0)
	for i := 0; i < len(src2); i += 4 {
		tmp := make([]uint8, 0)
		for j := 0; j < 4; j++ {
			tmp = append(tmp, multiSingle(src1[j], src2[i+j]))
		}
		var tmprst uint8
		for _, v := range tmp {
			tmprst = tmprst ^ v
		}
		rst = append(rst, tmprst)
	}

	return rst
}

func rollMix(src []uint8) []uint8{
	if len(src) != 16{
		return nil
	}
	rst := make([]uint8, 16)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			rst[i*4+j] = src[j*4+i]
		}
	}

	return rst
}

func mixColumns(src []uint8) []uint8{
	if len(src) != 16 {
		return nil
	}

	tmproll := rollMix(src)
	tmprst := make([]uint8, 0, 16)
	for i := 0; i < 16; i+=4 {
		tmprst = append(tmprst, multiGroup(tmproll[i:i+4], mixConst)...)
	}
	rst := rollMix(tmprst)
	return rst
}


func normalRound(src []uint8, i int) []uint8{
	rst := subBytes(src)
	rst = shiftRows(rst)
	rst = mixColumns(rst)
	rst = xor(rst, keys[i])
	return rst
}

func lastRound(src []uint8) []uint8{
	rst := subBytes(src)
	rst = shiftRows(rst)
	rst = xor(rst, keys[9])
	return rst
}

func (self *EncryptAes) Encrypt(src []byte, key []byte) ([]byte, error) {

	tmp := xor(src,key)
	fmt.Printf("tmp: %x\n", tmp)
	for i := 0; i < 9; i++ {
		tmp = normalRound(tmp, i)
	}
	fmt.Printf("tmp: %x\n", tmp)
	rst := lastRound(tmp)

	return rst, nil
}

//func (self *EncryptAes) Decrypt(src []byte, key []byte) ([]byte, error) {
//
//}
