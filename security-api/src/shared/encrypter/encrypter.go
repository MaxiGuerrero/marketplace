package encrypter

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Responsable to implement the logical encrypter. It implement the encrypter interface.
type Encrypter struct {
	cost int
}

// Create a instance of the encrypter struct.
func CreateEncrypter(cost int)*Encrypter{
	return &Encrypter{cost}
}

// Implementation to generate a hash.
func (e Encrypter) GenerateHash(password []byte) []byte{
	hash, err := bcrypt.GenerateFromPassword(password,e.cost)
	if err != nil {
		log.Panicf("Cannot encrypt it, %v",err.Error())
	}
	return hash
}

// Implementation to compare a plain password with hashed password gotten from database.
func (e Encrypter) Compare(hashedPassword, password []byte) bool{
	equal := bcrypt.CompareHashAndPassword(hashedPassword,password)
	return equal == nil
}