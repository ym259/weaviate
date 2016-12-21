package devices


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveDevicesCreateLocalAuthTokensHandlerFunc turns a function with the right signature into a weave devices create local auth tokens handler
type WeaveDevicesCreateLocalAuthTokensHandlerFunc func(WeaveDevicesCreateLocalAuthTokensParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveDevicesCreateLocalAuthTokensHandlerFunc) Handle(params WeaveDevicesCreateLocalAuthTokensParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveDevicesCreateLocalAuthTokensHandler interface for that can handle valid weave devices create local auth tokens params
type WeaveDevicesCreateLocalAuthTokensHandler interface {
	Handle(WeaveDevicesCreateLocalAuthTokensParams, interface{}) middleware.Responder
}

// NewWeaveDevicesCreateLocalAuthTokens creates a new http.Handler for the weave devices create local auth tokens operation
func NewWeaveDevicesCreateLocalAuthTokens(ctx *middleware.Context, handler WeaveDevicesCreateLocalAuthTokensHandler) *WeaveDevicesCreateLocalAuthTokens {
	return &WeaveDevicesCreateLocalAuthTokens{Context: ctx, Handler: handler}
}

/*WeaveDevicesCreateLocalAuthTokens swagger:route POST /devices/createLocalAuthTokens devices weaveDevicesCreateLocalAuthTokens

Creates client and device local auth tokens to be used by a client locally.

*/
type WeaveDevicesCreateLocalAuthTokens struct {
	Context *middleware.Context
	Handler WeaveDevicesCreateLocalAuthTokensHandler
}

func (o *WeaveDevicesCreateLocalAuthTokens) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveDevicesCreateLocalAuthTokensParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
