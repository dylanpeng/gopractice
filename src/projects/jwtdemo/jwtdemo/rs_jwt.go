package jwtdemo

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func JwtRsDemo() {
	initJWT()

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

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, myClaims)
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		log.Fatalf("SignedString failed. err: %s", err)
	}

	fmt.Printf("sign success. token:%s\n", tokenString)
	claims, err := RsValidToken(tokenString)

	if err != nil {
		log.Fatalf("HsValidToken failed. err: %s", err)
	}

	fmt.Printf("valid token success:\n%+v", claims)
}

func initJWT() {
	signBytes, err := os.ReadFile("./jwt_key/rsa_private.key")
	if err != nil {
		log.Fatal(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := os.ReadFile("./jwt_key/rsa_public.pem")
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func RsValidToken(tokenString string) (claims *MyClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, RsGetSignKey)

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

func RsGetSignKey(token *jwt.Token) (result interface{}, err error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return verifyKey, nil
}
