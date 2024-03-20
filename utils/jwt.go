package utils

import "github.com/dgrijalva/jwt-go"

// Specify the encryption key
var jwtSecret = []byte("2d9a0da267bee9c14d8e7aaedeca907c")

// A claim is the state and additional metadata of some entity (usually a user).
type Claims struct {
	UserId      int    `json:"id"`
	ClientId    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	jwt.StandardClaims
}

// Obtain the claims object information according to the passed token value, (and then obtain the username and password in it)
func ParseToken(token string) (*Claims, error) {

	// The method is mainly used to parse the authentication claim, and the internal process of decoding and verification is mainly the specific decoding and verification process, and finally returns the *token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// We get the Claims object from Token Claims and use assertions to convert that object into Claims that we define ourselves
		// To pass in pointers, structs in the project are passed with pointers, saving space.
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
