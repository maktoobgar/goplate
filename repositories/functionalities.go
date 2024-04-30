package repositories

import (
	"regexp"
	g "service/global"
	"service/static_models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func hasPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes)
}

func (u *User) IsHashed() bool {
	// Bcrypt hash pattern
	bcryptPattern := "^\\$2[aby]?(?:a|b)?\\$\\d{2}\\$[./0-9A-Za-z]{53}$"

	// Check if the string matches the bcrypt hash pattern
	matched, _ := regexp.MatchString(bcryptPattern, u.Password)
	return matched
}

func (u *User) HashPassword() {
	u.Password = hasPassword(u.Password)
}

func (u *User) IsSamePassword(rawPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(rawPassword)); err != nil {
		return false
	}

	return true
}

func (u *User) GenerateToken(refreshToken ...bool) Token {
	var isRefreshToken = false
	if len(refreshToken) > 0 {
		isRefreshToken = refreshToken[0]
	}

	var period = g.CFG.AccessTokenLifePeriod
	if isRefreshToken {
		period = g.CFG.RefreshTokenLifePeriod
	}
	expirationTime := time.Now().Add(time.Duration(period) * (time.Hour * 24))

	var tokenType = static_models.AccessTokenType
	if isRefreshToken {
		tokenType = static_models.RefreshTokenType
	}
	claims := &static_models.Claims{
		UserId: u.ID,
		Type:   tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := tkn.SignedString(g.SecretKeyBytes)

	return NewToken(tokenString, isRefreshToken, u.ID, expirationTime)
}
