package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sagara-test/src/auth/infrastructure/helper"
	"sagara-test/src/common/handler"
	"sagara-test/src/common/infrastructure/db"
	"sagara-test/src/common/interfaces"
	"sagara-test/src/common/utility"
	"strings"

	"github.com/fasthttp/router"
	"github.com/miguelpragier/handy"
	"github.com/valyala/fasthttp"
)

var (
	// DB ...
	DB *db.ConnectTo
)

// init db
func init() {
	DB = db.NewDBConnectionFactory(1)
}

// AuthApp ...
type AuthApp struct {
	interfaces.IApplication
}

// NewAuthApp ...
func NewAuthApp() *AuthApp {
	// Place where we init infrastructure, repo etc
	a := AuthApp{}
	return &a
}

// Initialize will be called when application run
func (a *AuthApp) Initialize(r *router.Router) {
	a.addRouter(r)
	log.Println("Auth App Initialize")
}

// Destroy will be called when app shutdown
func (a *AuthApp) Destroy() {
	// Clean up resource here
	log.Println("Auth App released...")
}

// Route declaration
func (a *AuthApp) addRouter(r *router.Router) {
	r.POST("/api/login", login)
}

func login(ctx *fasthttp.RequestCtx) {
	// Request
	ctx.Request.Header.Add("Content-Type", "application/json")
	request := VMAuthRequest{}
	body := ctx.Request.Body()
	err := json.Unmarshal(body, &request)

	// Response
	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload: ", err)
		return
	}

	// Check format email
	if !handy.CheckEmail(request.Username) || !strings.Contains(request.Username, ".") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(helper.ResponseEmailFormat))))
		return
	}

	data, err := loginService(ctx, request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(helper.ResponseAuthenticationInvalid+":"+err.Error()))))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(data, nil)))

}
