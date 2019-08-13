package des

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestTransform(t *testing.T) {

	//func Transform(src []byte, trans []uint8) []byte
	a := assert.New(t)

	t.Run("pc1", func(t *testing.T) {

		rst := transform(uint64(0x133457799bbcdff1), 64, pc1, 56)
		a.Equal(
			uint64(0xf0ccaaf556678f),
			rst,
			"invalid rst.")
	})

	t.Run("pc2", func(t *testing.T) {

		rst := transform(uint64(0xe19955faaccf1e), 56, pc2, 48)
		a.Equal(
			uint64(0x1b02effc7072),
			rst,
			"invalid rst")
	})

	t.Run("ip", func(t *testing.T) {

		rst := transform(uint64(0x0123456789abcdef), 64, ip, 64)
		a.Equal(
			uint64(0xcc00ccfff0aaf0aa),
			rst,
			"invalid rst")
	})

	t.Run("E", func(t *testing.T) {

		rst := transform(uint64(0xf0aaf0aa), 32, E, 48)
		a.Equal(
			uint64(0x7a15557a1555),
			rst,
			"invalid rst")
	})

	t.Run("P", func(t *testing.T) {

		rst := transform(uint64(0x5c82b597), 32, P, 32)
		a.Equal(
			uint64(0x234aa9bb),
			rst,
			"invalid rst")
	})

	t.Run("ip_1", func(t *testing.T) {

		rst := transform(uint64(0x0a4cd99543423234), 64, ip_1, 64)
		a.Equal(
			uint64(0x85e813540f0ab405),
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

func TestCalculateSubKeys(t *testing.T) {
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {

		rst := calculateSubKeys(uint64(0x133457799bbcdff1))
		a.Equal(
			uint64(0x1b02effc7072),
			rst[0], "invalid rst.")
	})
}

func TestS2(t *testing.T) {
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		rst := s(uint64(0x6117ba866527))
		a.Equal(uint64(0x5c82b597), rst, "invalid rst")
	})
}

func TestCalculateLoop(t *testing.T) {
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		k = calculateSubKeys(uint64(0x133457799bbcdff1))
		l, r := calculateLoop(0x0123456789abcdef, false)
		a.Equal(uint32(0x43423234), l, "invalid rst")
		a.Equal(uint32(0x0a4cd995), r, "invalid rst")

		k = k[:0]
	})
}

func TestEncryptDes_Encrypt(t *testing.T) {
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		var des EncryptDes
		key := []byte{
			0x13, 0x34, 0x57, 0x79, 0x9b, 0xbc, 0xdf, 0xf1,
		}
		rst, _ := des.Encrypt([]byte{
			0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		}, key)
		rstNum := binary.BigEndian.Uint64(rst)
		a.Equal(uint64(0x85E813540F0AB405), rstNum, "invalid rst")
	})

}

func TestEncryptDes_Decrypt(t *testing.T) {
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		var des EncryptDes
		key := []byte{
			0x13, 0x34, 0x57, 0x79, 0x9b, 0xbc, 0xdf, 0xf1,
		}
		rst, _ := des.Decrypt([]byte{
			0x85, 0xE8, 0x13, 0x54, 0x0F, 0x0A, 0xB4, 0x05,
		}, key)
		rstNum := binary.BigEndian.Uint64(rst)
		a.Equal(uint64(0x0123456789ABCDEF), rstNum, "invalid rst")
	})
}
