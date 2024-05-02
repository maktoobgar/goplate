package repositories

import (
	"database/sql"
	"regexp"
	g "service/global"
	"service/pkg/copier"
	"service/static_models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golodash/godash/strings"
	"github.com/kataras/iris/v12"
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

func (u *User) GenerateAccessToken(ctx iris.Context, db *sql.DB) (Token, string) {
	token := NewToken(u.ID)
	token, _ = New(db).CreateToken(ctx, copier.Copy(&CreateTokenParams{}, &token))

	expirationTime := time.Now().Add(time.Duration(g.CFG.AccessTokenLifePeriod) * (time.Hour * 24))
	claims := &static_models.Claims{
		Id:     token.ID,
		UserId: u.ID,
		Type:   static_models.AccessTokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := tkn.SignedString(g.SecretKeyBytes)

	return token, tokenString
}

func (u *User) GenerateRefreshToken(ctx iris.Context, db *sql.DB, relatedAccessTokenId int32) (Token, string) {
	token := NewToken(u.ID)
	token, _ = New(db).CreateToken(ctx, copier.Copy(&CreateTokenParams{}, &token))

	expirationTime := time.Now().Add(time.Duration(g.CFG.RefreshTokenLifePeriod) * (time.Hour * 24))
	claims := &static_models.Claims{
		Id:            token.ID,
		UserId:        u.ID,
		Type:          static_models.RefreshTokenType,
		AccessTokenId: relatedAccessTokenId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := tkn.SignedString(g.SecretKeyBytes)

	return token, tokenString
}

func (u *User) Reformat() {
	if u.Avatar.Valid {
		if !strings.StartsWith(u.Avatar.String, g.CFG.Domain) {
			u.Avatar.String = g.CFG.Domain + u.Avatar.String
		}
	}
}
