package security

import (
	"fmt"
	"thelastking/kingseafood/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const keyJwt = "asfdgrhtheerte"

func JwtToken(data *model.Users) (string, string, error) {
	//thanh phan cua 1 token
	newClaimsAccess := &model.Token{
		UserID: data.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	newClaimsRefresh := &model.Token{
		UserID: data.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	//khoi tao token
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaimsAccess).SignedString([]byte(keyJwt))
	if err != nil {
		panic(err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaimsRefresh).SignedString([]byte(keyJwt))
	if err != nil {
		panic(err)
	}
	return accessToken, refreshToken, nil
}

func UpdateToken(data *model.Users) (string, string, error) {
	//new token
	newAccessToken, newRefreshToken, err := JwtToken(data)
	if err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}

func ValidateToken(userToken string) (*model.Token, error) {
	//xac thuc token
	token, err := jwt.ParseWithClaims(
		userToken,
		&model.Token{},
		func(token *jwt.Token) (interface{}, error) { //callback cung cap khoa bi mat(keyJwt) xac thuc token
			return []byte(keyJwt), nil
		},
	)
	if err != nil {
		return nil, err
	}
	//trich xuat cac claims duoc token xac thuc
	if claims, ok := token.Claims.(*model.Token); ok {
		// Check token expiration
		if claims.ExpiresAt < time.Now().Local().Unix() {
			fmt.Println("claims token expires")
		}
		return claims, nil
	} else {
		panic("invalid claims")
	}
}
