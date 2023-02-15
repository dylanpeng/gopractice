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
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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
	//var err error
	//signBytes := []byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA9djdZJ1UEL5p6KPxr7AHIxDglTwT7giKzE+qYqA9oxacbUlV\nTkZGAK8443Hr5gG7vcNWZvTb1wmTAobOJCwo99tW40DTPyvrtVFIPFt5Nmo/CICf\n+NAN5b+cR1hR2SVWtyRZSEpnD9/cXKf7cCHkJcs0Mej8t2kgRYQ3W0CLYwO5sNni\n23f0TYGvtXYXC2A4F21ip9elRQt73h4x4vSgxCltwC2TP8gBYrAPksgutx8iLcDf\nv+MPMA5KlaxDjdzkuVjXxhFQKenSOPOqcd+zuar3VhihM6nX8ip5x81+RoXXtjp1\nLG3pL9+qwihYQ5Xxd9VIGKVCHAJJDxySb1uS4wIDAQABAoIBAQDlVYLZC8ZSxD2p\nrd2T5SITPPgzXlK9FqzbgGlSDWbSDxKnA+SW2wkMNGheC3RiIDXRBDpCWqIFC8Je\ndgAwUB17cNmxrlQhNshvYL6Ax1fgQeZA+TPBd9uu+TpAd4wKg0FMIJVE0VsovMwk\nhvMPnB3mf5NWB6BPO7rF/lthPWmJVyk+TMp0S4SKuPG9vDk24eLygcdgqdC5o2xd\nbmvCut1RRmbviiObvV81lE3zAjfu2K4Pccc5HZvfc9hpG10tTlEHNxCQ3nt2i2h/\n1HLzbs7Tqbkuxrppv8uBo8cKDCpMXT3WZiVGBMWahKjFabDJuDxusXlsI5726/w9\nPREUHx2pAoGBAPy4Ov1SYHQ1zK5ib31euvPl3SMKi+BeobExWvFairZM2VpwQ3qc\nGA4TvnzYtvvb1lMDPx3mzr9sZR3aWPqW8dES09alWJ35UC9EDLfJwXBWVNSbXoMp\nqUZ3ohNsrZuwwYPEzX1m1NF2h7YCpYhsLsoBZqexjP81UpkdjSIda8jlAoGBAPkJ\nzCfZAVqFkj0uXnQKku0RWs8GRk6cZ1AN4y0qTnRRA/jFtzBR6R0TAaeqvp0q2hKJ\nl5Eed1YncvTGZd0SbHA5dx2T4WR8wgY8amsIsLbTZyqrwMhw+q39PBBa0Wknr1Ki\nR1YqLjN8UISWOqgyc15KDLkxre7dcdzv1XOfiJgnAoGBAOnHPwJxvrohvnseogX2\nqLjQTbWJnwVqZOb2Qit8V072XiZ0LWfxl6sGBrOVAgiQP35BRZTSmzSnAA8SmjcN\nhRqj8QThpc1VASEIMT+eymux4P1f0JlC481FA9A2O48Hfqv3VSQJCRvPKxFq91fw\nw4Oosh60dzrqR8NOe+0wDDIlAoGBAJ8X1jdil03H5Nt24tpI4wHVw2hb/tA7dHic\n1pNE4qfGFb54OIYC3eQ3/yeomWr4NCYBhjUr/FqqivK6R9rJ6UJsQ58+mI/Eb4Li\nV62W+KVjOhX1cQvbuRkrnJJqIjuGIaetidsOyUMU2K9K9Z/70t3aenRYu1/MUfAt\nuvPJZ86jAoGAbfjjiGbYQKj1o/8AN+jnFHq6xu+utGWGyDzWoQM94n0H6Fin/mHx\nAdtqib6DPZVgT4wnmVaCLYb7FifW8kjPpZMp88G9tyiAexVdBjuJ97w50lWnpY/c\njdUTFvL2MiGIj0DXwz7clFuxpf6qXcFFHUdNsCsGmT4RtSYrJ5LlY74=\n-----END RSA PRIVATE KEY-----")

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	//verifyBytes := []byte("-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9djdZJ1UEL5p6KPxr7AH\nIxDglTwT7giKzE+qYqA9oxacbUlVTkZGAK8443Hr5gG7vcNWZvTb1wmTAobOJCwo\n99tW40DTPyvrtVFIPFt5Nmo/CICf+NAN5b+cR1hR2SVWtyRZSEpnD9/cXKf7cCHk\nJcs0Mej8t2kgRYQ3W0CLYwO5sNni23f0TYGvtXYXC2A4F21ip9elRQt73h4x4vSg\nxCltwC2TP8gBYrAPksgutx8iLcDfv+MPMA5KlaxDjdzkuVjXxhFQKenSOPOqcd+z\nuar3VhihM6nX8ip5x81+RoXXtjp1LG3pL9+qwihYQ5Xxd9VIGKVCHAJJDxySb1uS\n4wIDAQAB\n-----END PUBLIC KEY-----")
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
