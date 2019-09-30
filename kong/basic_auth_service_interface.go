// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// BasicAuthServiceInterface ...
type BasicAuthServiceInterface interface {
	// Create creates a basic-auth credential in Kong
	// If an ID is specified, it will be used to
	// create a basic-auth in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumerUsernameOrID *string, basicAuth *BasicAuth) (*BasicAuth, error)
	// Get fetches a basic-auth credential from Kong.
	Get(ctx context.Context, consumerUsernameOrID, usernameOrID *string) (*BasicAuth, error)
	// Update updates a basic-auth credential in Kong
	Update(ctx context.Context, consumerUsernameOrID *string, basicAuth *BasicAuth) (*BasicAuth, error)
	// Delete deletes a basic-auth credential in Kong
	Delete(ctx context.Context, consumerUsernameOrID, usernameOrID *string) error
	// List fetches a list of basic-auth credentials in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*BasicAuth, *ListOpt, error)
	// ListAll fetches all basic-auth credentials in Kong.
	// This method can take a while if there
	// a lot of basic-auth credentials present.
	ListAll(ctx context.Context) ([]*BasicAuth, error)
	// ListForConsumer fetches a list of basic-auth credentials
	// in Kong associated with a specific consumer.
	// opt can be used to control pagination.
	ListForConsumer(ctx context.Context, consumerUsernameOrID *string, opt *ListOpt) ([]*BasicAuth, *ListOpt, error)
}
