package bd

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword func para encriptar pass al momento de registrar un usuario*/
func EncryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
