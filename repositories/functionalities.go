package repositories

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) IsHashed() bool {
	// Bcrypt hash pattern
	bcryptPattern := "^\\$2[aby]?(?:a|b)?\\$\\d{2}\\$[./0-9A-Za-z]{53}$"

	// Check if the string matches the bcrypt hash pattern
	matched, _ := regexp.MatchString(bcryptPattern, u.Password)
	return matched
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 16)
	u.Password = string(bytes)
}
