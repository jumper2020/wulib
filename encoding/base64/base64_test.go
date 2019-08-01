package base64

import (
	"encoding/base64"
	"testing"
)

//Test
func TestBase64Coding_Encode(t *testing.T) {

	var bc Base64Coding
	rst := bc.Encode([]byte("it's a simple test."))
	if string(rst) != "aXQncyBhIHNpbXBsZSB0ZXN0Lg=="{
		t.Errorf("result not expected.")
	}


}

func TestBase64Coding_Decode(t *testing.T) {

	var bc Base64Coding
	rst := bc.Decode([]byte("aXQncyBhIHNpbXBsZSB0ZXN0Lg=="))
	if string(rst) != "it's a simple test."{
		t.Errorf("result not expected.")
	}
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
