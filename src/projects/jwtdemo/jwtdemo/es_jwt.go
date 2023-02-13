package jwtdemo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var (
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
)

func InitEsKey() {
	rankdKey := rand.Reader

	prk, err := ecdsa.GenerateKey(elliptic.P256(), rankdKey)
	if err != nil {
		log.Fatalf("GenerateKey failed. err: %s", err)
		return
	}

	puk := prk.PublicKey
	fmt.Printf("prkD=%X\n", prk.D)
	fmt.Printf("prkX=%X\n", prk.X)
	fmt.Printf("prkY=%X\n", prk.Y)
	fmt.Println("prk", prk, " \npbk", puk)

	publicKey = &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     prk.X,
		Y:     prk.Y,
	}

	privateKey = &ecdsa.PrivateKey{D: prk.D, PublicKey: *publicKey}
}

func JwtEsDemo() {
	InitEsKey()

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

	token := jwt.NewWithClaims(jwt.SigningMethodES256, myClaims)
	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		log.Fatalf("SignedString failed. err: %s", err)
	}

	fmt.Printf("sign success. token:%s\n", tokenString)
	claims, err := EsValidToken(tokenString)

	if err != nil {
		log.Fatalf("HsValidToken failed. err: %s", err)
	}

	fmt.Printf("valid token success:\n%+v", claims)
}

func EsValidToken(tokenString string) (claims *MyClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, EsGetSignKey)

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

func EsGetSignKey(token *jwt.Token) (result interface{}, err error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return publicKey, nil
}
