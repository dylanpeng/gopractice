package main

import "gopractice/projects/jwtdemo/jwtdemo"

func main() {
	// HMAC-SHA256
	jwtdemo.JwtHsDemo()
	// RSA-SHA256
	jwtdemo.JwtRsDemo()
	// ECDSA-SHA256
	jwtdemo.JwtEsDemo()
}
