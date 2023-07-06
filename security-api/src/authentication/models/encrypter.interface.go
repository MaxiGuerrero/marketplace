package models

// Interface to implement methods about hash management.
type IEncrypter interface{
	Compare(hashedPassword, password []byte) bool
}