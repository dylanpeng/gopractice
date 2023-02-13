package jwtdemo

import "github.com/golang-jwt/jwt/v4"

const mySigningKey string = "sdjfiejisdjfiej"

type MyClaims struct {
	*jwt.RegisteredClaims
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}
