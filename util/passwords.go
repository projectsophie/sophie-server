package util

import "golang.org/x/crypto/bcrypt"

// GenerateHash encrypts password and
// generates hash and salt.
func GenerateHash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password.")
	}
	return string(hash)
}

// VerifyPassword compares hashed password from database
// with given plain password.
func VerifyPassword(hashedPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plainPassword)
	return err == nil
}
