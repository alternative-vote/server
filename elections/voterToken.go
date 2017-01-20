package elections

import "github.com/dgrijalva/jwt-go"

type VoterClaims struct {
	ElectionId   string
	EmailAddress string
	jwt.StandardClaims
}
