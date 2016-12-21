package acl_entries


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveACLEntriesUpdateHandlerFunc turns a function with the right signature into a weave acl entries update handler
type WeaveACLEntriesUpdateHandlerFunc func(WeaveACLEntriesUpdateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveACLEntriesUpdateHandlerFunc) Handle(params WeaveACLEntriesUpdateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveACLEntriesUpdateHandler interface for that can handle valid weave acl entries update params
type WeaveACLEntriesUpdateHandler interface {
	Handle(WeaveACLEntriesUpdateParams, interface{}) middleware.Responder
}

// NewWeaveACLEntriesUpdate creates a new http.Handler for the weave acl entries update operation
func NewWeaveACLEntriesUpdate(ctx *middleware.Context, handler WeaveACLEntriesUpdateHandler) *WeaveACLEntriesUpdate {
	return &WeaveACLEntriesUpdate{Context: ctx, Handler: handler}
}

/*WeaveACLEntriesUpdate swagger:route PUT /devices/{deviceId}/aclEntries/{aclEntryId} aclEntries weaveAclEntriesUpdate

Update an ACL entry.

*/
type WeaveACLEntriesUpdate struct {
	Context *middleware.Context
	Handler WeaveACLEntriesUpdateHandler
}

func (o *WeaveACLEntriesUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveACLEntriesUpdateParams()

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
