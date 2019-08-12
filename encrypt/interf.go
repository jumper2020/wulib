package encrypt

/* This file defines the interface of hash implementation.*/

type EncryptInterf interface {
	Encrypt(src []byte, key []byte) ([]byte, error)
	Decrypt(src []byte, key []byte) ([]byte, error)
}


