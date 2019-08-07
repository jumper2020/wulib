package main

/* This file defines the interface of encoding implementation.*/

type EncodingInterf interface {
	Encode(src []byte) []byte
	Decode(src []byte) []byte
}


