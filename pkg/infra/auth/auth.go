package auth

import (
	"dissent-api-service/pkg/infra/db"
	"dissent-api-service/pkg/infra/entities"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenProvider struct {
	HMACSigningKey []byte
	Cache          TokenCache
	DB             *db.DBConn
}

func (t *TokenProvider) NewToken(username string, password string) (string, error) {

	// check user database
	// store user passwords here - the other one just for "data"

	err := t.DB.VerifyPassword(username, password)
	if err != nil {
		return "", fmt.Errorf("error verifying password, err %v", err)
	}

	cl := &entities.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)

	tokenString, err := tok.SignedString(t.HMACSigningKey)
	if err != nil {
		return "", fmt.Errorf("error signing token, err %v", err)
	}

	t.Cache.CacheToken(tokenString)

	return tokenString, nil

}

func (t *TokenProvider) getJWTTokenWithClaims(tok string) (*jwt.Token, error) {

	// https://github.com/MicahParks/keyfunc
	var keyFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return t.HMACSigningKey, nil
	}

	claims := &entities.Claims{}
	return jwt.ParseWithClaims(tok, claims, keyFunc)
}

func (t *TokenProvider) CheckToken(tok string, username string) error {

	tokenSplit := strings.Split(tok, "Bearer ")
	if len(tokenSplit) != 2 {
		return fmt.Errorf("error removing 'Bearer', incorrect length")
	}
	tok = tokenSplit[1]

	err := t.Cache.Exists(tok)
	if err != nil {
		return err
	}

	parsed, err := t.getJWTTokenWithClaims(tok)
	if err != nil {
		return fmt.Errorf("error parsing jwt token, err %v", err)
	}

	claims, ok := parsed.Claims.(*entities.Claims)
	if !ok || !parsed.Valid {
		return fmt.Errorf("error, token is not valid")
	}

	// Check if the token subject (sub) matches the provided username
	if claims.Username != username {
		return fmt.Errorf("error, token subject does not match the provided username")
	}

	return nil
}

func InitialiseTokenProvider(signingKey string, db *db.DBConn) TokenProvider {
	return TokenProvider{
		HMACSigningKey: []byte(signingKey),
		Cache:          NewTokenCache(),
		DB:             db,
	}

}
