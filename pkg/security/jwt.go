package security

import (
	"errors"
	"fmt"
	"log"
	"os"
	"thelastking/kingseafood/model"
	redisdb "thelastking/kingseafood/pkg/redisDB"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var keyJwt string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	keyJwt = os.Getenv("KEY_JWT")
	if keyJwt == "" {
		log.Fatal("KEY_JWT is not set")
	}
}

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

func UpdateToken(userToken string) (string, error) {
	//kiểm tra xem refresh token có hợp lệ hay không
	claims, err := ValidateToken(userToken)
	if err != nil {
		return "", errors.New("refreshtoken khong hop le")
	}
	//kiem tra refresh token con han hay khong
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return "", errors.New("refresh token het han ban phai dang nhap lai")
	}
	//neu con han thi cap nhap moi lai 1 asscesstoken
	newClaimsAccess := model.Token{
		UserID: claims.UserID,
		Role:   claims.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	asscessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaimsAccess).SignedString([]byte(keyJwt))
	if err != nil {
		return "", err
	}
	return asscessToken, nil
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
		//check token redis
		redisInstance := redisdb.GetInstanceRedis()
		exists, err := redisInstance.CheckRefreshToken()
		if err != nil {
			return nil, fmt.Errorf("error checking refresh token: %v", err)
		}

		if !exists {
			fmt.Println("Refresh token invalid on Redis.")
			return nil, fmt.Errorf("refresh token not found")
		}
		return claims, nil
	} else {
		panic("invalid claims")
	}
}
