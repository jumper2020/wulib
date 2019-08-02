package base64

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Test

func TestBase64Coding_Encode(t *testing.T) {

	a := assert.New(t)

	var bc Base64Coding
	t.Run("test1", func(t *testing.T){
		rst := bc.Encode([]byte("it's a simple test."))
		a.Equal("aXQncyBhIHNpbXBsZSB0ZXN0Lg==", string(rst), "result not expected.")
	})

	t.Run("test2", func(t *testing.T){
		rst := bc.Encode([]byte("it's a complicated test."))
		a.Equal("aXQncyBhIGNvbXBsaWNhdGVkIHRlc3Qu", string(rst), "result not expected.")
	})

	t.Run("test3", func(t *testing.T){
		rst := bc.Encode([]byte("it's a abnormal test.一个非正常测试")) //注意确定编码方式与测试网站上使用的编码一致，这里是utf-8
		a.Equal("aXQncyBhIGFibm9ybWFsIHRlc3Qu5LiA5Liq6Z2e5q2j5bi45rWL6K+V", string(rst), "result not expected.")
	})
}

func TestBase64Coding_Decode(t *testing.T) {
	a := assert.New(t)

	var bc Base64Coding
	t.Run("test1", func(t *testing.T){
		rst := bc.Decode([]byte("aXQncyBhIHNpbXBsZSB0ZXN0Lg=="))
		a.Equal("it's a simple test.", string(rst), "result not expected.")
	})

	t.Run("test2", func(t *testing.T){
		rst := bc.Decode([]byte("aXQncyBhIGNvbXBsaWNhdGVkIHRlc3Qu"))
		a.Equal("it's a complicated test.", string(rst), "result not expected.")
	})

	t.Run("test3", func(t *testing.T){
		rst := bc.Decode([]byte("aXQncyBhIGFibm9ybWFsIHRlc3Qu5LiA5Liq6Z2e5q2j5bi45rWL6K+V"))
		a.Equal("it's a abnormal test.一个非正常测试", string(rst), "result not expected.")
	})

}



//Benchmark
func BenchmarkBase64Coding_Encode(b *testing.B) {
	var bc Base64Coding
	src := []byte("it's a simple test.")
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		_ = string(bc.Encode(src))
	}
}


func BenchmarkBase64Coding_Decode(b *testing.B) {
	var bc Base64Coding
	src := []byte("aXQncyBhIHNpbXBsZSB0ZXN0Lg==")
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		bc.Decode(src)
	}
}


func BenchmarkEncode(b *testing.B){

	src := []byte("it's a simple test.")
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		base64.StdEncoding.EncodeToString(src)
	}
}


func BenchmarkDecode(b *testing.B){

	src := "aXQncyBhIHNpbXBsZSB0ZXN0Lg=="
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		base64.StdEncoding.DecodeString(src)
	}
}
