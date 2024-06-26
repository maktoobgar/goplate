package static_models

import "github.com/dgrijalva/jwt-go"

const (
	AccessTokenType  = "Access"
	RefreshTokenType = "Refresh"
)

type Claims struct {
	Id            int32
	UserId        int32
	Type          string
	AccessTokenId int32
	jwt.StandardClaims
}
