package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"ylink/ext/globalkey"
	"ylink/ext/jwtdata"
)

type Player2CtxMiddleware struct {
}

func NewPlayer2CtxMiddleware() *Player2CtxMiddleware {
	return &Player2CtxMiddleware{}
}

func (m *Player2CtxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		ctx := r.Context()
		token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
			return []byte(globalkey.AccessSecret), nil
		})
		// 将获取的token中的Claims强转为MapClaims
		claims, _ := token.Claims.(jwt.MapClaims)
		playerId := claims[jwtdata.JwtKeyPlayerId]
		gameId := claims[jwtdata.JwtKeyGameId]
		ctx = context.WithValue(ctx, jwtdata.JwtKeyPlayerId, playerId)
		ctx = context.WithValue(ctx, jwtdata.JwtKeyGameId, gameId)
		next(w, r.WithContext(ctx))
	}
}
