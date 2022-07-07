package utils

import "golang.org/x/crypto/bcrypt"

func BcryptHash(pwd string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(password)
}

func BcryptCheck(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
