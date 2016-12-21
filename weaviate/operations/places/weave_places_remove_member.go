package places


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeavePlacesRemoveMemberHandlerFunc turns a function with the right signature into a weave places remove member handler
type WeavePlacesRemoveMemberHandlerFunc func(WeavePlacesRemoveMemberParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeavePlacesRemoveMemberHandlerFunc) Handle(params WeavePlacesRemoveMemberParams) middleware.Responder {
	return fn(params)
}

// WeavePlacesRemoveMemberHandler interface for that can handle valid weave places remove member params
type WeavePlacesRemoveMemberHandler interface {
	Handle(WeavePlacesRemoveMemberParams) middleware.Responder
}

// NewWeavePlacesRemoveMember creates a new http.Handler for the weave places remove member operation
func NewWeavePlacesRemoveMember(ctx *middleware.Context, handler WeavePlacesRemoveMemberHandler) *WeavePlacesRemoveMember {
	return &WeavePlacesRemoveMember{Context: ctx, Handler: handler}
}

/*WeavePlacesRemoveMember swagger:route POST /places/{placeId}/removeMember places weavePlacesRemoveMember

Removes a member of a place. Does not affect device sharing, devices might still be shared with this member.

*/
type WeavePlacesRemoveMember struct {
	Context *middleware.Context
	Handler WeavePlacesRemoveMemberHandler
}

func (o *WeavePlacesRemoveMember) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeavePlacesRemoveMemberParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
