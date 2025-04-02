package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"time"
)

type jwtProvider struct {
	secret string
	prefix string
}

func NewTokenJWTProvider(prefix string, secret string) *jwtProvider {
	return &jwtProvider{prefix: prefix, secret: secret}
}

type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {

	// generate the JWT
	now := time.Now()

	claims := myClaims{
		Payload: common.TokenPayload{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        fmt.Sprintf("%d", now.UnixNano()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// Return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(token string) (tokenprovider.TokenPayload, error) {

	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return claims.Payload, nil
}
