package models

// Interface to implement methods about hash management.
type IEncrypter interface{
	GenerateHash(password []byte) []byte
	Compare(hashedPassword, password []byte) bool
}