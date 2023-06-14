package models

type IEncrypter interface{
	GenerateHash(password []byte) []byte
	Compare(hashedPassword, password []byte) bool
}