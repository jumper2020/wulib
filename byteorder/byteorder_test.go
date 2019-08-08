package byteorder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigU16ToBS(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := bigU16ToBS(uint16(0x1234))
		a.Equal([]byte{0x12, 0x34}, rst, "invalid rst")
	})
}

func TestBigU32ToBS(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := bigU32ToBS(uint32(0x12345678))
		a.Equal([]byte{0x12, 0x34, 0x56, 0x78}, rst, "invalid rst")
	})
}


func TestBigU64ToBS(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := bigU64ToBS(uint64(0x1234567887654321))
		a.Equal([]byte{0x12, 0x34, 0x56, 0x78, 0x87, 0x65, 0x43, 0x21}, rst, "invalid rst")
	})
}

func TestBigBSToU16(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := bigBSToU16([]byte{0x12, 0x34})
		a.Equal(uint16(0x1234), rst, "invalid rst")
	})

}

func TestBigBSToU32(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := bigBSToU32([]byte{0x12, 0x34, 0x56, 0x78})
		a.Equal(uint32(0x12345678), rst, "invalid rst")
	})

}

func TestBigBSToU64(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := bigBSToU64([]byte{0x12, 0x34, 0x56, 0x78, 0x87, 0x65, 0x43, 0x21})
		a.Equal(uint64(0x1234567887654321), rst, "invalid rst")
	})
}



func TestLittleU16ToBS(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := littleU16ToBS(uint16(0x1234))
		a.Equal([]byte{0x34, 0x12}, rst, "invalid rst")
	})
}

func TestLittleU32ToBS(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := littleU32ToBS(uint32(0x12345678))
		a.Equal([]byte{0x78, 0x56, 0x34, 0x12}, rst, "invalid rst")
	})
}


func TestLittleU64ToBS(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := littleU64ToBS(uint64(0x1234567887654321))
		a.Equal([]byte{0x21, 0x43, 0x65, 0x87, 0x78, 0x56, 0x34, 0x12}, rst, "invalid rst")
	})
}

func TestLittleBSToU16(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := littleBSToU16([]byte{0x12, 0x34})
		a.Equal(uint16(0x3412), rst, "invalid rst")
	})
}

func TestLittleBSToU32(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := littleBSToU32([]byte{0x12, 0x34, 0x56, 0x78})
		a.Equal(uint32(0x78563412), rst, "invalid rst")
	})
}

func TestLittleBSToU64(t *testing.T){

	t.Run("test1", func(t *testing.T) {
		a := assert.New(t)
		rst := littleBSToU64([]byte{0x12, 0x34, 0x56, 0x78, 0x87, 0x65, 0x43, 0x21})
		a.Equal(uint64(0x2143658778563412), rst, "invalid rst")
	})
}



