package gcjwt

import (
	"errors"
	"fmt"
	"golang-core/config/JWTConfig"
	"golang-core/models/UserModel"
	"golang-core/responseCodes"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)



type Claims struct {
	Username string `json:"username"`
	UID      uint `json:"uid"`
	jwt.StandardClaims
}

var jwtkey =[]byte (os.Getenv("JWT_SECRET_KEY"))

func CreateToken(user UserModel.User)(token string, err error){

	expirationTime := time.Now().Add(JWTConfig.TTL * time.Minute)
	claims := &Claims{
		Username: user.Username,
		UID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(tokenData)
	fmt.Println(jwtkey)
	

	token, err = tokenData.SignedString(jwtkey)
	fmt.Println(token)
	return
}

func RefreshToken(token string)(newToken string, err error){
	
	claims := &Claims{}

	tkn,tknErr := jwt.ParseWithClaims(token, claims, func(newToken *jwt.Token) (interface{}, error) {
		return jwtkey,nil
	})
	if !tkn.Valid {
		err = errors.New(string(responseCodes.StatusUnauthorized))
		return
	}

	if tknErr != nil {

	}
	
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now())>30*time.Second{
		err = errors.New(string(responseCodes.StatusBadRequest))
		return
	}

	expirationTime := time.Now().Add(JWTConfig.TTL * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newTokenData := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	newToken, err = newTokenData.SignedString(jwtkey)
	
	return
}