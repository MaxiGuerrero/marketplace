package models

type IEncrypter interface{
	GenerateHash(password []byte) ([]byte, error)
	Compare(hashedPassword, password []byte) error
}