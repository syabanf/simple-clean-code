package interfaces

import "github.com/fasthttp/router"

type IApplication interface {
	Initialize(r *router.Router)
	Destroy()
}
