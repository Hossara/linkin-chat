package hash

import (
	"golang.org/x/crypto/bcrypt"

	userService "github.com/Hossara/linkin-chat/internal/user"
)

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) HashPassword(password string) (string, error) {
	if len(password) > 72 {
		return "", userService.ErrPasswordTooLong
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (b *BcryptHasher) Validate(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
