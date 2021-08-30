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

	middleware "sagara-test/src/middleware/application"
	"sagara-test/src/product/domain"
	"sagara-test/src/product/domain/entity"
	"sagara-test/src/product/infrastructure/repository"

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

// ProductApp ...
type ProductApp struct {
	interfaces.IApplication
}

// NewProductApp ...
func NewProductApp() *ProductApp {
	// Place where we init infrastructure, repo etc
	a := ProductApp{}
	return &a
}

// Initialize will be called when application run
func (a *ProductApp) Initialize(r *router.Router) {
	a.addRouter(r)
	log.Println("Product App Initialize")
}

// Destroy will be called when app shutdown
func (a *ProductApp) Destroy() {
	// Clean up resource here
	log.Println("Product App released...")
}

// Route declaration
func (a *ProductApp) addRouter(r *router.Router) {
	r.POST("/api/product", middleware.Validate(insertProduct))
	r.POST("/api/product/get", middleware.Validate(getProduct))
	r.PUT("/api/product/get", middleware.Validate(updateProduct))
	r.DELETE("/api/product/get", middleware.Validate(deleteProduct))
}

func getProduct(ctx *fasthttp.RequestCtx) {
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
	repo := repository.NewProductRepository(DB)
	data, _, err := domain.GetProduct(ctx, &repo, request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(fasthttp.StatusMessage(fasthttp.StatusNotFound)))))
		return
	}

	//responseData := []VMProduct{}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(data, nil)))
}

func insertProduct(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")

	request := VMProduct{}
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

	repo := repository.NewProductRepository(DB)
	resultData, err := domain.InsertProduct(ctx, &repo, request.ToEntity())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(fasthttp.StatusMessage(fasthttp.StatusInternalServerError)))))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(resultData, nil)))
}

func updateProduct(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")

	request := VMProduct{}
	body := ctx.Request.Body()
	err := json.Unmarshal(body, &request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New("Bad JSON Payload"))))
		log.Println("Error Bad Request JSON Payload: ", err)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	repo := repository.NewProductRepository(DB)
	err = domain.UpdateProduct(ctx, &repo, request.ToEntity())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, err)))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(request, nil)))
}

func deleteProduct(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Add("Content-Type", "application/json")
	guid := ctx.QueryArgs().Peek("guid")

	req := entity.ModelProduct{
		GUID: string(guid),
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	repo := repository.NewProductRepository(DB)
	err := domain.RemoveProduct(ctx, &repo, string(guid))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(nil, errors.New(fasthttp.StatusMessage(fasthttp.StatusInternalServerError)))))
		return
	}

	fmt.Fprintf(ctx, utility.PrettyPrint(handler.DefaultResponse(req, nil)))
}
