package hash

/* This file defines the interface of hash implementation.*/

type HashInterf interface {
	Hash(src []byte) []byte
}


