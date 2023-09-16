package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	publicKey, privateKey string
}

func NewAuth(publicKey, privateKey string) *AuthMiddleware {
	return &AuthMiddleware{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (a *AuthMiddleware) AuthHeader(ctx *gin.Context) {
	headerPublicKey := ctx.GetHeader("PUBLIC-KEY")
	privateKey := ctx.GetHeader("PRIVATE-KEY")

	if a.publicKey != headerPublicKey || a.privateKey != privateKey {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access"))
	}
}
