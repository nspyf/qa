package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTObj struct {
	SignKey string `json:"sign_key"`
	Issuer  string `json:"issuer"`
}

type Claims struct {
	Public interface{}
	jwt.StandardClaims
}

func (obj *JWTObj) Init(path string) {
	err := ReadJSON(path,&obj)
	DoErr(err)
	return
}

func (obj *JWTObj) GenerateToken(data interface{},validTime int64) string {
	now := time.Now().Unix()

	claims := Claims{
		Public: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now+validTime,
			IssuedAt: now,
			Issuer: obj.Issuer,
			NotBefore: now,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(obj.SignKey))
	DoErr(err)

	return token
}

func (obj *JWTObj) ParseToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.New("未携带token，拒绝访问")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface {}, error) {
		return []byte(obj.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, errors.New("token解析为空")
	}

	if claims,ok := token.Claims.(*Claims);ok&&token.Valid {
		return claims,nil
	}

	return nil, errors.New("token解析失败(UnknownError)")
}
