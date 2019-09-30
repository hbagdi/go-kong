// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// RouteServiceInterface ...
type RouteServiceInterface interface {
	// Create creates a Route in Kong
	// If an ID is specified, it will be used to
	// create a route in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, route *Route) (*Route, error)
	// CreateInService creates a route associated with serviceID
	CreateInService(ctx context.Context, serviceID *string, route *Route) (*Route, error)
	// Get fetches a Route in Kong.
	Get(ctx context.Context, nameOrID *string) (*Route, error)
	// Update updates a Route in Kong
	Update(ctx context.Context, route *Route) (*Route, error)
	// Delete deletes a Route in Kong
	Delete(ctx context.Context, nameOrID *string) error
	// List fetches a list of Routes in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Route, *ListOpt, error)
	// ListAll fetches all Routes in Kong.
	// This method can take a while if there
	// a lot of Routes present.
	ListAll(ctx context.Context) ([]*Route, error)
	// ListForService fetches a list of Routes in Kong associated with a service.
	// opt can be used to control pagination.
	ListForService(ctx context.Context, serviceNameOrID *string, opt *ListOpt) ([]*Route, *ListOpt, error)
}
