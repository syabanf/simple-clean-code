package application

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"sagara-test/src/common/constant"
	"sagara-test/src/common/handler"
	"sagara-test/src/common/utility"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

var (
	bearerAuthPrefix = []byte("Bearer")
	// DB ...

)

// init db
func init() {

}

// Validate ...
func Validate(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Get Bearer Authorization
		auth := ctx.Request.Header.Peek("Authorization")
		if bytes.HasPrefix(auth, bearerAuthPrefix) {
			_, err := checkToken(string(auth))
			if err == nil {
				next(ctx)
				return
			}
		}

		ctx.Response.Header.Set("WWW-Authenticate", "Bearer realm=Restricted")
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Unauthorization"))))

	})
}

func checkToken(token string) (authClaim AuthClaims, err error) {
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		log.Println("Failed parse token")
	}

	signedToken := strings.TrimSpace(splitToken[1])

	tokenDecode, err := jwt.ParseWithClaims(signedToken, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWTKey), nil
	})

	if err != nil {
		log.Println("Error token")
		return
	}

	if claims, ok := tokenDecode.Claims.(*AuthClaims); ok && tokenDecode.Valid {

		// Check User is exist

		authClaim.ID = claims.ID
		authClaim.RoleID = claims.RoleID
		authClaim.IsVerified = claims.IsVerified
		authClaim.ExpiresAt = claims.ExpiresAt

	} else {
		log.Println("Error token:", err)
		return
	}

	return

}

// GetAuthClaim ...
func GetAuthClaim(ctx *fasthttp.RequestCtx) (authClaim AuthClaims, err error) {
	// Get Bearer Authorization
	auth := ctx.Request.Header.Peek("Authorization")
	if bytes.HasPrefix(auth, bearerAuthPrefix) {
		authClaim, err := checkToken(string(auth))
		return authClaim, err
	}
	return
}
