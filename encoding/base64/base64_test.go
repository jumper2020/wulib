package base64

import (
	"encoding/base64"
	"testing"
)

//Test
func TestBase64Coding_Encode(t *testing.T) {

	var bc Base64Coding

	t.Run("test1", func(t *testing.T){
		rst := bc.Encode([]byte("it's a simple test."))
		if string(rst) != "aXQncyBhIHNpbXBsZSB0ZXN0Lg=="{
			t.Errorf("result not expected.")
		}
	})

	t.Run("test2", func(t *testing.T){
		rst := bc.Encode([]byte("it's a complicated test."))
		if string(rst) != "aXQncyBhIGNvbXBsaWNhdGVkIHRlc3Qu"{
			t.Errorf("result not expected.")
		}
	})

	t.Run("test3", func(t *testing.T){
		rst := bc.Encode([]byte("it's a abnormal test.一个非正常测试")) //注意确定编码方式与测试网站上使用的编码一致，这里是utf-8
		if string(rst) != "aXQncyBhIGFibm9ybWFsIHRlc3Qu5LiA5Liq6Z2e5q2j5bi45rWL6K+V"{
			t.Errorf("result not expected.")
		}
	})
}

func TestBase64Coding_Decode(t *testing.T) {
	if testing.Short(){
		t.Skip("skip this test func.")
	}

	var bc Base64Coding

	t.Run("test1", func(t *testing.T){
		rst := bc.Decode([]byte("aXQncyBhIHNpbXBsZSB0ZXN0Lg=="))
		if string(rst) != "it's a simple test."{
			t.Errorf("result not expected.")
		}
	})

	t.Run("test2", func(t *testing.T){
		rst := bc.Decode([]byte("aXQncyBhIGNvbXBsaWNhdGVkIHRlc3Qu"))
		if string(rst) != "it's a complicated test."{
			t.Errorf("result not expected.")
		}
	})

	t.Run("test3", func(t *testing.T){
		rst := bc.Decode([]byte("aXQncyBhIGFibm9ybWFsIHRlc3Qu5LiA5Liq6Z2e5q2j5bi45rWL6K+V"))
		if string(rst) != "it's a abnormal test.一个非正常测试"{
			t.Errorf("result not expected.")
		}
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
