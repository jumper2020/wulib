package base58

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//首先需要明白这里的 base58　是用于 bitcoin 中的，因此输入是 256bit 的一个大整数，　输出是特定字符集中字符组成的串。
//编码之前是一个十六进制数(使用字符串表示), 因此一定是偶数个字符 (sha256 hash结果为32个字节，可以表示为十六进制超大整数
//编码之前大小写无所谓，因为无论大小写，表示出的整数是一致的
//编码之后大小写需要区分开，因为编码之后为字符串，大小写表示不同字符
//这里使用 math/big 来表示大整数, 该库中与[]byte相互转换时是会忽略最开始的'0'的，也就是说比如在big.Int中存放　0x0123，　
//Text 函数输出为　"123", 因此这里就需要补充一个 '0', 具体参见 Decode 函数
//而github中有些golang　base58的编码还是与bitcoin中的base58不完全相同，　比如bitcoin中需要根据0x00数量来补充'1'
//网上的编解码工具也需要找用于bitcoin中的base58
//编解码测试　参考　http://lenschulwitz.com/base58


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
