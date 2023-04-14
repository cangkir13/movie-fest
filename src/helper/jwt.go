package helper

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	jwt.StandardClaims
}

/*
* @SignIn login jwt service
 */
func SignIn(props JWTClaim) (tokenString string, err error) {
	jwtKey := os.Getenv("JWT_SECRET")
	expirationTime := time.Now().Add(time.Hour * 24 * 2)
	claims := &JWTClaim{
		UserID:   props.UserID,
		Username: props.Username,
		RoleID:   props.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))
	return
}

type DataClaims struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
}

func DecodeJWTAdmin(strToken string) (data DataClaims, errData error) {
	token, _, err := new(jwt.Parser).ParseUnverified(strToken, jwt.MapClaims{})
	if err != nil {
		errData = fmt.Errorf("jwt error")
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		id, okInt := claims["id"].(float64)
		if !okInt {
			errData = fmt.Errorf("id is not integer")
			return
		}

		name, ok := claims["username"].(string)
		if !ok {
			errData = fmt.Errorf("name is not string")
			return
		}

		role, ok := claims["role_id"].(float64)
		if !ok {
			errData = fmt.Errorf("id is not integer")
			return
		}

		data = DataClaims{
			UserID:   int(id),
			Username: name,
			RoleID:   int(role),
		}
		return
	} else {
		errData = fmt.Errorf("false jwt")
		return
	}
}

func ValidateToken(singedToken string) (err error) {
	jwtKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(
		singedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
