package main

/* This file defines the interface of encoding implementation.*/

//Input of "Encode" function - 'estr' is a binary sequence, return base64 result
//Input of "Decode" function - 'dstr' is a base64 result, return binary sequence
type EncodingInterf interface {
	Encode(src []byte) []byte
	Decode(src []byte) []byte
}


