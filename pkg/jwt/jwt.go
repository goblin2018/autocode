package jwt

import (
	"auto/pkg/e"

	"github.com/dgrijalva/jwt-go"
)

const issuer = "goblin-is-best"
const key = "a-nice-future"

type Claims struct {
	Id        string `json:"idd,omitempty"`
	Phone     string `json:"pho,omitempty"`
	Role      string `json:"rol,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Issuer    string `json:"iss,omitempty"`
}

func (c Claims) Valid() error {
	if c.Id == "" || c.Issuer != issuer {
		return e.TokenError
	}
	return nil
}

func GenToken(c Claims) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, _ := token.SignedString([]byte(key))
	return tokenStr

}

func ParseToken(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if t == nil {
		return nil, err
	}

	if cliams, ok := t.Claims.(*Claims); ok && t.Valid {
		return cliams, nil
	}
	return nil, err
}
