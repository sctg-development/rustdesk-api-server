package utils

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userId int, username, token, clientId, uuid string) (string, error) {

	claims := Claims{
		UserId: userId,
		//Username: username,
		//Password: password,
		AccessToken: token,
		ClientId:    clientId,
		//Uuid:     uuid,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: 0,
			// 指定token发行人
			Issuer: "baozier",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//This method generates a signature string internally, which is then used to obtain a complete and signed token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
