package repositories

import "golang.org/x/crypto/bcrypt"

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 16)
	u.Password = string(bytes)
}
