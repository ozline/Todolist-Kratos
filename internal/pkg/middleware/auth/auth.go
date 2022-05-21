package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

var (
	secret string = "MTAxNTkwMTg1Mw=="
)

type JWTClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Status   int    `json:"status" example:"0"`
	jwt.StandardClaims
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			spew.Dump(secret)
			//获取JWTTOKEN
			var JWTToken string
			if md, ok := metadata.FromServerContext(ctx); ok {
				JWTToken = md.Get("x-md-global-token")
				spew.Dump(JWTToken)
			} else if tr, ok := transport.FromServerContext(ctx); ok {
				JWTToken = strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)[1]
			} else {
				return nil, errors.New("Missing Token")
			}

			claims, err := JWTParse(JWTToken, secret)

			if err != nil {
				return nil, err
			}

			if claims.ExpiresAt < time.Now().Unix() {
				return nil, errors.New("Invaild Token")
			}

			ctx = context.WithValue(ctx, "username", claims.Username)
			ctx = context.WithValue(ctx, "id", claims.Id)
			ctx = context.WithValue(ctx, "status", claims.Status)

			spew.Dump(claims)
			return handler(ctx, req)
		}
	}
}

// func FromContext(ctx context.Context) *CurrentUser {
// 	return ctx.Value(currentUserKey).(*CurrentUser)
// }

func JWTParse(token string, secret string) (*JWTClaims, error) { //解析JWT
	tmp, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// return []byte(conf.Auth.Secret), nil
		// #BASE64,JWTToken
		return []byte(secret), nil
	})
	if tmp != nil {
		claims, result := tmp.Claims.(*JWTClaims)
		if result && tmp.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWTGenerate(claims JWTClaims, secret string) (string, error) { //生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
