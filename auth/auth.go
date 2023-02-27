package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"voltunes-chick-api-master-product/exception"
	"voltunes-chick-api-master-product/helper"
)

type CreateAuthFunc func(userID uint, tokenDetails *TokenDetails)

type AccessDetails struct {
	UserID uint
	Role   string
}

type TokenDetails struct {
	AccessToken        string
	RefreshToken       string
	AccessUUID         string
	RefreshUUID        string
	UserAgent          string
	RemoteAddress      string
	AtExpired          int64
	RefeshTokenExpired int64
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check JWT Token
		tokenAuth, err := ExtractTokenMetadata(c.Request)
		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		// Check Permission User
		// if !helper.Contains(roles, tokenAuth.Role) {
		// 	helper.PanicIfError(exception.ErrPermissionDenied)
		// }

		next(c, tokenAuth)
	}
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			UserID: uint(userId),
			Role:   claims["role"].(string),
		}, nil
	}
	return nil, err
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
