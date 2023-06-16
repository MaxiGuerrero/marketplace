package models

type IEncrypter interface{
	Compare(hashedPassword, password []byte) bool
}