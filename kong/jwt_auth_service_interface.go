// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// JWTAuthServiceInterface ...
type JWTAuthServiceInterface interface {
	// Create creates a JWT credential in Kong
	// If an ID is specified, it will be used to
	// create a JWT in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumerUsernameOrID *string, jwtAuth *JWTAuth) (*JWTAuth, error)
	// Get fetches a JWT credential from Kong.
	Get(ctx context.Context, consumerUsernameOrID, keyOrID *string) (*JWTAuth, error)
	// Update updates a JWT credential in Kong
	Update(ctx context.Context, consumerUsernameOrID *string, jwtAuth *JWTAuth) (*JWTAuth, error)
	// Delete deletes a JWT credential in Kong
	Delete(ctx context.Context, consumerUsernameOrID, keyOrID *string) error
	// List fetches a list of JWT credentials in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*JWTAuth, *ListOpt, error)
	// ListAll fetches all JWT credentials in Kong.
	// This method can take a while if there
	// a lot of JWT credentials present.
	ListAll(ctx context.Context) ([]*JWTAuth, error)
	// ListForConsumer fetches a list of jwt credentials
	// in Kong associated with a specific consumer.
	// opt can be used to control pagination.
	ListForConsumer(ctx context.Context, consumerUsernameOrID *string, opt *ListOpt) ([]*JWTAuth, *ListOpt, error)
}
