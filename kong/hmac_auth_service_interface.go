// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// HMACAuthServiceInterface ...
type HMACAuthServiceInterface interface {
	// Create creates a hmac-auth credential in Kong
	// If an ID is specified, it will be used to
	// create a hmac-auth in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumerUsernameOrID *string, hmacAuth *HMACAuth) (*HMACAuth, error)
	// Get fetches a hmac-auth credential from Kong.
	Get(ctx context.Context, consumerUsernameOrID, usernameOrID *string) (*HMACAuth, error)
	// Update updates a hmac-auth credential in Kong
	Update(ctx context.Context, consumerUsernameOrID *string, hmacAuth *HMACAuth) (*HMACAuth, error)
	// Delete deletes a hmac-auth credential in Kong
	Delete(ctx context.Context, consumerUsernameOrID, usernameOrID *string) error
	// List fetches a list of hmac-auth credentials in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*HMACAuth, *ListOpt, error)
	// ListAll fetches all hmac-auth credentials in Kong.
	// This method can take a while if there
	// a lot of hmac-auth credentials present.
	ListAll(ctx context.Context) ([]*HMACAuth, error)
	// ListForConsumer fetches a list of hmac-auth credentials
	// in Kong associated with a specific consumer.
	// opt can be used to control pagination.
	ListForConsumer(ctx context.Context, consumerUsernameOrID *string, opt *ListOpt) ([]*HMACAuth, *ListOpt, error)
}
