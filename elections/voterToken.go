package elections

import (
	"github.com/alternative-vote/server/consts"
	"github.com/dgrijalva/jwt-go"
)

type VoterClaims struct {
	ElectionId   string
	EmailAddress string
	jwt.StandardClaims
}

func GetVoterToken(electionId, emailAddress string) string {

	//stuff whatever info we want into the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, VoterClaims{
		ElectionId:   electionId,
		EmailAddress: emailAddress,
	})

	//sign the token with the hardcoded secret
	tokenString, err := token.SignedString(consts.Secret)
	if err != nil {
		panic(err)
	}

	return tokenString
}
