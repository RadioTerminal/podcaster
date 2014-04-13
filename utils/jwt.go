package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// User defines all the functions necessary to work with the user's authentication.
// The caller should implement these functions for whatever system of authentication
// they choose to use

func LoginRequired(publicKey []byte) martini.Handler {
	return func(context martini.Context, res http.ResponseWriter, req *http.Request) {
		token, err := jwt.ParseFromRequest(req, func(t *jwt.Token) ([]byte, error) { return publicKey, nil })
		if token != nil && err == nil {
			if token.Valid {
				context.Map(token)
				context.Next()
			} else {
				http.Error(res, err.Error(), http.StatusForbidden)
			}

		} else {
			http.Error(res, err.Error(), http.StatusUnauthorized)
		}
	}
}

func GenerateAuthToken(claims map[string]interface{}, privateKey []byte) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = claims
	tokenString, err := token.SignedString(privateKey)
	if err == nil {
		return fmt.Sprintf("Bearer %v", tokenString), nil
	} else {
		return tokenString, err
	}
}

func GenKeyPairIfNone(privateName string, publicName string) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	privatekey := filepath.Join(dir, privateName)
	publickey := filepath.Join(dir, publicName)

	if _, err := os.Stat(string(privatekey)); os.IsNotExist(err) {

		log.Println("Generating JWT private key at ", string(privatekey))
		k, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Fatal(err)
		}
		var private pem.Block
		private.Type = "RSA PRIVATE KEY"
		private.Bytes = x509.MarshalPKCS1PrivateKey(k)
		pp := new(bytes.Buffer)
		pem.Encode(pp, &private)
		err = ioutil.WriteFile(string(privatekey), pp.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Generating JWT public key at ", string(privatekey))
		var public pem.Block
		public.Type = "RSA PUBLIC KEY"
		public.Bytes, _ = x509.MarshalPKIXPublicKey(&k.PublicKey)
		ps := new(bytes.Buffer)
		pem.Encode(ps, &public)
		err = ioutil.WriteFile(string(publickey), ps.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

}

var GetKey = func(file string) []byte {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	key, err := ioutil.ReadFile(filepath.Join(dir, file))
	if err != nil {
		log.Fatal(err)
	}
	return key
}
