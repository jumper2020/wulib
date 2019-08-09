package sha256

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha256Hash_Hash(t *testing.T) {
	a := assert.New(t)
	var sha256 Sha256Hash

	t.Run("test1", func(t *testing.T) {
		rst := sha256.Hash([]byte("a"))
		expected := []byte{
			0xca,0x97,0x81,0x12,0xca,0x1b,0xbd,0xca,0xfa,0xc2,0x31,0xb3,0x9a,0x23,0xdc,0x4d,
			0xa7,0x86,0xef,0xf8,0x14,0x7c,0x4e,0x72,0xb9,0x80,0x77,0x85,0xaf,0xee,0x48,0xbb,
		}
		a.Equal(expected, rst, "invalid result.")
	})

	t.Run("test2", func(t *testing.T) {
		rst := sha256.Hash([]byte("abc"))
		expected := []byte{
			0xba,0x78,0x16,0xbf,0x8f,0x01,0xcf,0xea,0x41,0x41,0x40,0xde,0x5d,0xae,0x22,0x23,
			0xb0,0x03,0x61,0xa3,0x96,0x17,0x7a,0x9c,0xb4,0x10,0xff,0x61,0xf2,0x00,0x15,0xad,
		}
		a.Equal(expected, rst, "invalid result.")
	})

	t.Run("test3", func(t *testing.T) {
		rst := sha256.Hash([]byte("message digest"))
		expected := []byte{
			0xf7,0x84,0x6f,0x55,0xcf,0x23,0xe1,0x4e,0xeb,0xea,0xb5,0xb4,0xe1,0x55,0x0c,0xad,
			0x5b,0x50,0x9e,0x33,0x48,0xfb,0xc4,0xef,0xa3,0xa1,0x41,0x3d,0x39,0x3c,0xb6,0x50,
		}
		a.Equal(expected, rst, "invalid result.")
	})


	t.Run("test4", func(t *testing.T) {
		rst := sha256.Hash([]byte("abcdefghijklmnopqrstuvwxyz"))
		expected := []byte{
			0x71,0xc4,0x80,0xdf,0x93,0xd6,0xae,0x2f,0x1e,0xfa,0xd1,0x44,0x7c,0x66,0xc9,0x52,
			0x5e,0x31,0x62,0x18,0xcf,0x51,0xfc,0x8d,0x9e,0xd8,0x32,0xf2,0xda,0xf1,0x8b,0x73,
		}
		a.Equal(expected, rst, "invalid result.")
	})

	t.Run("test5", func(t *testing.T) {
		rst := sha256.Hash([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"))
		expected := []byte{
			0x64,0xcc,0x0a,0xb1,0xa8,0x8e,0xfe,0xac,0xd6,0x4f,0xa7,0x9e,0xce,0x34,0xed,0xe0,
			0x44,0xcd,0x6d,0x1c,0x32,0xc2,0xa1,0xc2,0x79,0x1e,0x5b,0xa2,0x06,0x3c,0x1b,0xea,
		}
		a.Equal(expected, rst, "invalid result.")
	})

	t.Run("test6", func(t *testing.T) {
		rst := sha256.Hash([]byte("8a683566bcc7801226b3d8b0cf35fd97"))
		expected := []byte{
			0x38,0xcd,0x55,0x85,0xf1,0x6e,0x9b,0xf2,0x23,0x96,0xc3,0xc8,0xe8,0x28,0x1b,0xc1,
			0x95,0xf4,0x23,0x05,0xa2,0x1d,0x4f,0x6d,0xb7,0x08,0x1b,0x5d,0xe2,0x54,0xe9,0x03,
		}
		a.Equal(expected, rst, "invalid result.")
	})
}