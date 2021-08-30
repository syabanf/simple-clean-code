package main

import (
	"log"
	"sagara-test/src/common/infrastructure"
	"sagara-test/src/common/infrastructure/web"
	"sagara-test/src/common/interfaces"

	"github.com/fasthttp/router"

	Authenticate "sagara-test/src/auth/application"
	Media "sagara-test/src/media/application"
	Product "sagara-test/src/product/application"
)

func main() {
	// App and Routing Initialization
	var apps = map[string]interfaces.IApplication{}
	router := web.NewRouter()
	initialize(apps, router)

	// Turn on Web API Server
	// ports, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	ws, _ := web.NewWebServer(router, 8081)
	go ws.Listen()

	ctx := infrastructure.CaptureSignal()
	// stop serving API
	ws.Shutdown(ctx)
	<-ctx.Done()
	// Clean up each app
	destroy(apps)

	log.Println("🟢 app-init app has been shut down successfully.")
}

func initialize(apps map[string]interfaces.IApplication, router *router.Router) {

	// Register application to run
	apps["auhtenticate"] = Authenticate.NewAuthApp()
	apps["product"] = Product.NewProductApp()
	apps["media"] = Media.NewMediaApp()

	for _, v := range apps {
		v.Initialize(router)
	}
}

func destroy(apps map[string]interfaces.IApplication) {
	for _, v := range apps {
		v.Destroy()
	}
}
