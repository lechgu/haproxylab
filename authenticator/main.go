package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lechgu/authenticator/dtos"
)

func main() {
	r := gin.Default()
	r.POST("", issueToken)
	r.Run()
}

func issueToken(ctx *gin.Context) {
	var req dtos.AuthenticateRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.Fatalln(err)
	}
	key, err := loadPrivateKey("keys/key.pem")
	if err != nil {
		log.Fatalln(err)
	}
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"iss": "emily",
		"aud": req.Audience,
		"sub": req.APIKey,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatalln(err)
	}

	ctx.String(http.StatusOK, tokenString)
}

func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	priv, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	privPem, _ := pem.Decode(priv)
	var parsedKey interface{}
	parsedKey, err = x509.ParsePKCS1PrivateKey(privPem.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}
	return privateKey, nil
}
