#!/bin/sh
cd ./
openssl genrsa -out rsa_private.key 2048
openssl rsa -in rsa_private.key -pubout -out rsa_public.pem
openssl ecparam -genkey -name secp256r1 -out ecdsa_private.key
openssl ec -in ecdsa_private.key -pubout -out ecdsa_public.pem