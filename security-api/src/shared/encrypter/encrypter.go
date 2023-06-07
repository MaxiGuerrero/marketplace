package encrypter

import "golang.org/x/crypto/bcrypt"

type Encrypter struct {
	cost int
}

func CreateEncrypter(cost int)*Encrypter{
	return &Encrypter{cost}
}

func (e Encrypter) GenerateHash(password []byte) ([]byte, error){
	return bcrypt.GenerateFromPassword(password,e.cost)
}

func (e Encrypter) Compare(hashedPassword, password []byte) error{
	return bcrypt.CompareHashAndPassword(hashedPassword,password)
}