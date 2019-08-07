package base58

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)



func TestBase58Coding_Encode(t *testing.T) {
	var bc Base58Coding
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		result := bc.Encode([]byte("0123"))
		a.Equal("62", string(result), "result not expected.")
	})

	t.Run("test2", func(t *testing.T) {
		result := bc.Encode([]byte("000123"))
		a.Equal("162", string(result), "result not expected.")
	})

	t.Run("test3", func(t *testing.T) {
		result := bc.Encode([]byte("00000123"))
		a.Equal("1162", string(result), "result not expected.")
	})

	t.Run("test4", func(t *testing.T) {
		result := bc.Encode([]byte("0005f573"))
		a.Equal("13162", string(result), "result not expected.")
	})

	t.Run("test5", func(t *testing.T) {
		result := bc.Encode([]byte("3a8c0d9f2d"))
		a.Equal("7c7Tz6x", string(result), "result not expected.")
	})

	t.Run("test6", func(t *testing.T) {
		result := bc.Encode([]byte("3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c"))
		a.Equal("4wYWbSnD3Z5NVuc81NZbDbHnhnS4Uv2DxdWw17gcC4iF", string(result), "result not expected.")
	})

	t.Run("test7", func(t *testing.T) {
		result := bc.Encode([]byte("00000d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c"))
		a.Equal("113jJq7UDBhQryMur1y33T7Dykzi3Usvs2Gfzipv1Zu", string(result), "result not expected.")
	})

	//3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c
}

func TestBase58Coding_Decode(t *testing.T) {

	var bc Base58Coding
	a := assert.New(t)

	t.Run("test1", func(t *testing.T) {
		result := bc.Decode([]byte("62"))
		a.Equal("0123", strings.ToLower(string(result)), "result not expected.")
	})

	t.Run("test2", func(t *testing.T) {
		result := bc.Decode([]byte("162"))
		a.Equal("000123", strings.ToLower(string(result)), "result not expected.")
	})

	t.Run("test3", func(t *testing.T) {
		result := bc.Decode([]byte("1162"))
		a.Equal("00000123", strings.ToLower(string(result)), "result not expected.")
	})

	t.Run("test4", func(t *testing.T) {
		result := bc.Decode([]byte("13162"))
		a.Equal("0005f573", strings.ToLower(string(result)), "result not expected.")
	})

	t.Run("test5", func(t *testing.T) {
		result := bc.Decode([]byte("7c7Tz6x"))
		a.Equal("3a8c0d9f2d", strings.ToLower(string(result)), "result not expected.")
	})

	t.Run("test6", func(t *testing.T) {
		result := bc.Decode([]byte("4wYWbSnD3Z5NVuc81NZbDbHnhnS4Uv2DxdWw17gcC4iF"))
		a.Equal("3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c", strings.ToLower(string(result)), "result not expected.")
	})

	t.Run("test7", func(t *testing.T) {
		result := bc.Decode([]byte("113jJq7UDBhQryMur1y33T7Dykzi3Usvs2Gfzipv1Zu"))
		a.Equal("00000d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c", strings.ToLower(string(result)), "result not expected.")
	})
}

func BenchmarkBase58Coding_Encode(b *testing.B) {
	var bc Base58Coding
	src := []byte("3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c0d9f2d3a8c")
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		_ = string(bc.Encode(src))
	}
}


func BenchmarkBase58Coding_Decode(b *testing.B) {
	var bc Base58Coding
	src := []byte("4wYWbSnD3Z5NVuc81NZbDbHnhnS4Uv2DxdWw17gcC4iF")
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		bc.Decode(src)
	}
}
