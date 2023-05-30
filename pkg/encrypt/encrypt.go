package encrypt

import "golang.org/x/crypto/bcrypt"

func Encrypt(str string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed)
}

func Compare(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
