package des

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//http://www.hankcs.com/security/des-algorithm-illustrated.html

func TestTransform(t *testing.T) {

	//func Transform(src []byte, trans []uint8) []byte
	a := assert.New(t)

	t.Run("pc1", func(t *testing.T) {

		//pc1 := []uint8{
		//	57, 49, 41, 33, 25, 17, 9,
		//	1, 58, 50, 42, 34, 26, 18,
		//	10, 2, 59, 51, 43, 35, 27,
		//	19, 11, 3, 60, 52, 44, 36,
		//	63, 55, 47, 39, 31, 23, 15,
		//	7, 62, 54, 46, 38, 30, 22,
		//	14, 6, 61, 53, 45, 37, 29,
		//	21, 13, 5, 28, 20, 12, 4,
		//}

		src := []byte{
			0x13, 0x34, 0x57, 0x79, 0x9b, 0xbc, 0xdf, 0xf1,
		}

		rst := transform(src, pc1)

		a.Equal(
			[]byte{0xf0, 0xcc, 0xaa, 0xf5, 0x56, 0x67, 0x8f},
			rst,
			"invalid rst.")
	})

	t.Run("pc2", func(t *testing.T) {
		src := []byte{
			0xe1, 0x99, 0x55, 0xfa, 0xac, 0xcf, 0x1e,
		}

		rst := transform(src, pc2)
		a.Equal(
			[]byte{
				0x1b, 0x02, 0xef, 0xfc, 0x70, 0x72,
			},
			rst,
			"invalid rst")
	})

	t.Run("ip", func(t *testing.T) {
		src := []byte{
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
		}

		rst := transform(src, ip)
		a.Equal(
			[]byte{
				0xcc, 0x00, 0xcc, 0xff, 0xf0, 0xaa, 0xf0, 0xaa,
			},
			rst,
			"invalid rst")
	})

	t.Run("E", func(t *testing.T) {
		src := []byte{
			//11110000 10101010 11110000 10101010
			0xf0, 0xaa, 0xf0, 0xaa,
		}

		rst := transform(src, E)
		a.Equal(
			[]byte{
				//01111010, 00010101, 01010101,  01111010, 00010101, 01010101
				0x7a, 0x15, 0x55, 0x7a, 0x15, 0x55,
			},
			rst,
			"invalid rst")
	})

	t.Run("P", func(t *testing.T) {
		src := []byte{
			0x5c, 0x82, 0xb5, 0x97,
		}

		rst := transform(src, P)
		a.Equal(
			[]byte{
				0x23, 0x4a, 0xa9, 0xbb,
			},
			rst,
			"invalid rst")
	})

	t.Run("ip_1", func(t *testing.T) {

		src := []byte{
			0x0a, 0x4c, 0xd9, 0x95, 0x43, 0x42, 0x32, 0x34,
		}

		rst := transform(src, ip_1)
		a.Equal(
			[]byte{
				0x85, 0xe8, 0x13, 0x54, 0x0f, 0x0a, 0xb4, 0x05,
			},
			rst,
			"invalid rst")
	})
}

func TestRotateLeftShift(t *testing.T) {

	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		num := uint64(0x35)
		bitLen := uint8(6)
		shift := uint8(2)
		rst := rotateLeftShift(num, bitLen, shift)
		a.Equal(uint64(0x17), rst, "invalid rst")
	})

	t.Run("test2", func(t *testing.T) {
		num := uint64(0x0f12)
		bitLen := uint8(12)
		shift := uint8(6)
		rst := rotateLeftShift(num, bitLen, shift)
		a.Equal(uint64(0x04bc), rst, "invalid rst")
	})
}

func TestCalculateSubKeys(t *testing.T){
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		key := []byte{
			//00010011 00110100 01010111 01111001 10011011 10111100 11011111 11110001
			0x13, 0x34, 0x57, 0x79, 0x9b, 0xbc, 0xdf, 0xf1,
		}
		rst := calculateSubKeys(key)
		fmt.Printf("rst: %v\n", rst)
		a.Equal([]byte{
			//00011011 00000010 11101111 11111100 01110000 01110010
			0x1b, 0x02, 0xef, 0xfc, 0x70, 0x72,
		}, rst[0], "invalid rst.")
	})
}