package jwtdemo

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

const mySigningKey string = "sdjfiejisdjfiej"

type MyClaims struct {
	*jwt.RegisteredClaims
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}

func (e *MyClaims) String() string {
	str, _ := json.Marshal(e)
	return fmt.Sprintf("%s", str)
}
