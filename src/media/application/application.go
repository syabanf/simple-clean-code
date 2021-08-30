package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sagara-test/src/common/handler"
	"sagara-test/src/common/infrastructure/db"
	"sagara-test/src/common/interfaces"
	"sagara-test/src/common/utility"
	"sagara-test/src/media/domain"
	"sagara-test/src/media/domain/entity"
	"sagara-test/src/media/infrastructure/repository"

	middleware "sagara-test/src/middleware/application"

	"github.com/fasthttp/router"
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

// MediaApp ...
type MediaApp struct {
	interfaces.IApplication
}

// NewMediaApp ...
func NewMediaApp() *MediaApp {
	// Place where we init infrastructure, repo etc
	a := MediaApp{}
	return &a
}

// Initialize will be called when application run
func (a *MediaApp) Initialize(r *router.Router) {
	a.addRouter(r)
	log.Println("Media App Initialize")
}

// Destroy will be called when app shutdown
func (a *MediaApp) Destroy() {
	// Clean up resource here
	log.Println("Media App released...")
}

// Route declaration
func (a *MediaApp) addRouter(r *router.Router) {
	r.POST("/api/media", middleware.Validate(insertMedia))
	r.POST("/api/media/get", middleware.Validate(getMedia))
	r.PUT("/api/media/get", middleware.Validate(updateMedia))
	r.DELETE("/api/media/get", middleware.Validate(deleteMedia))
}

func getMedia(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")
	request := entity.StructQuery{}
	body := ctx.Request.Body()
	err := json.Unmarshal(body, &request)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload: ", err)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	repo := repository.NewMediaRepository(DB)
	data, _, err := domain.GetMedia(ctx, &repo, request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(fasthttp.StatusMessage(fasthttp.StatusNotFound)))))
		return
	}

	//responseData := []VMMedia{}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(data, nil)))
}

func insertMedia(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")

	request := VMMedia{}
	body := ctx.Request.Body()
	err := json.Unmarshal(body, &request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload: ", err)
		return
	}

	// Response
	ctx.Response.Header.Set("Content-Type", "application/json")

	repo := repository.NewMediaRepository(DB)
	resultData, err := domain.InsertMedia(ctx, &repo, request.ToEntity())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(fasthttp.StatusMessage(fasthttp.StatusInternalServerError)))))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(resultData, nil)))
}

func updateMedia(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")

	request := VMMedia{}
	body := ctx.Request.Body()
	err := json.Unmarshal(body, &request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload: ", err)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	repo := repository.NewMediaRepository(DB)
	err = domain.UpdateMedia(ctx, &repo, request.ToEntity())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(request, nil)))
}

func deleteMedia(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")
	guid := ctx.QueryArgs().Peek("guid")

	req := entity.ModelMedia{
		GUID: string(guid),
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	repo := repository.NewMediaRepository(DB)
	err := domain.RemoveMedia(ctx, &repo, string(guid))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(fasthttp.StatusMessage(fasthttp.StatusInternalServerError)))))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(req, nil)))
}
