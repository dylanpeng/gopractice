package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"log"
	"time"
)

const mySigningKey string = "sdjfiejisdjfiej"

type MyClaims struct {
	*jwt.RegisteredClaims
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}

func main() {
	myClaims := &MyClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "dylan-api",
			Subject:   "dylan-api-token",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().AddDate(0, 0, 3)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        "1111111",
		},
		Name:   "test",
		Gender: 1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, err := token.SignedString([]byte(mySigningKey))

	if err != nil {
		log.Fatalf("SignedString failed. err: %s", err)
	}

	fmt.Printf("sign success. token:%s\n", tokenStr)

	claims, err := ValidToken(tokenStr)

	if err != nil {
		log.Fatalf("ValidToken failed. err: %s", err)
	}

	fmt.Printf("valid token success:\n%+v", claims)

}

func ValidToken(tokenStr string) (claims *MyClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, GetKey)

	valid := token.Valid

	if !valid {
		fmt.Printf("token not valid.\n")
		err = errors.New("token not valid")
		return
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		err = errors.New("token not valid")
		return
	}

	return
}

func GetKey(token *jwt.Token) (result interface{}, err error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(mySigningKey), nil
}
