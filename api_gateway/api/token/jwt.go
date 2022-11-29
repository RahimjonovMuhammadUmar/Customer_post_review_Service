package token

import (
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTHandler struct {
	Sub       string
	Iss       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	SignInKey string
	Log       logger.Logger
	Token     string
}

type CustomClaims struct {
	*jwt.Token
	Sub  string
	Iss  string
	Exp  float64
	Iat  float64
	Aud  []string
	Role string
}

// GenerateAuthJWT ...
func (jwtHandler *JWTHandler) GenerateAuthJWT() (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["iss"] = jwtHandler.Iss
	claims["sub"] = jwtHandler.Sub
	claims["exp"] = time.Now().Add(time.Hour * 500).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role
	claims["aud"] = jwtHandler.Aud
	access, err = accessToken.SignedString([]byte("UmarSecret"))
	if err != nil {
		jwtHandler.Log.Error("error access, err = accessToken.SignedString([]byte(UmarSecret)", logger.Error(err))
		return "", "", err
	}

	refresh, err = refreshToken.SignedString([]byte("UmarSecret"))
	if err != nil {
		jwtHandler.Log.Error("error refresh, err = accessToken.SignedString([]byte(jwtHandler.SignInKey)", logger.Error(err))
		return "", "", err
	}

	return access, refresh, nil
}

// ExtractClaims ...
func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	c := config.Load()
	token, err := jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(c.SignInKey), nil
	})
	if err != nil {
		fmt.Println("error while parsing SignInKey token.go 72")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		jwtHandler.Log.Error("invalid jwt token")
		return nil, err
	}
	return claims, nil
}
func ExtractClaims(tokenStr string, signInKey []byte) (jwt.MapClaims, error) {
	fmt.Println(tokenStr)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signInKey, nil
	})
	if err != nil {
		fmt.Println("Error jwt.go token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil

}
