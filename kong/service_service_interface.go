// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// SvcserviceInterface ...
type SvcserviceInterface interface {
	// Create creates an Service in Kong
	// If an ID is specified, it will be used to
	// create a service in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, service *Service) (*Service, error)
	// Get fetches an Service in Kong.
	Get(ctx context.Context, nameOrID *string) (*Service, error)
	// GetForRoute fetches a Service associated with routeID in Kong.
	GetForRoute(ctx context.Context, routeID *string) (*Service, error)
	// Update updates an Service in Kong
	Update(ctx context.Context, service *Service) (*Service, error)
	// Delete deletes an Service in Kong
	Delete(ctx context.Context, nameOrID *string) error
	// List fetches a list of Services in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Service, *ListOpt, error)
	// ListAll fetches all Services in Kong.
	// This method can take a while if there
	// a lot of Services present.
	ListAll(ctx context.Context) ([]*Service, error)
}
