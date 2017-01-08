package authentication

import (
	. "github.com/alternative-vote/server/generated"
	"github.com/dgrijalva/jwt-go"
)

var Secret = []byte("boggleboggle1234")

type CustomClaims struct {
	Email string
	jwt.StandardClaims
}

type Controller struct {
}

func (o *Controller) Login(req *LoginRequest) *LoginResponse {
	email, password := req.Body.Email, req.Body.Password

	//hardcoded single user for now
	if email != "test@fake.com" || password != "test" {
		panic(HttpError(401))
	}

	//stuff whatever info we want into the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Email: email,
	})

	//sign the token with the hardcoded secret
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		panic(err)
	}

	return &LoginResponse{
		StatusCode: 200,
		Body: map[string]string{
			"token": tokenString,
		},
	}

}
