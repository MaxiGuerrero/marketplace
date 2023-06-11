package encrypter

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Encrypter struct {
	cost int
}

func CreateEncrypter(cost int)*Encrypter{
	return &Encrypter{cost}
}

func (e Encrypter) GenerateHash(password []byte) []byte{
	hash, err := bcrypt.GenerateFromPassword(password,e.cost)
	if err != nil {
		log.Panicf("Cannot encrypt it, %v",err.Error())
	}
	return hash
}

func (e Encrypter) Compare(hashedPassword, password []byte) bool{
	equal := bcrypt.CompareHashAndPassword(hashedPassword,password)
	return equal == nil
}