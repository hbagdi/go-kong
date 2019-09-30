// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// KeyAuthServiceInterface ...
type KeyAuthServiceInterface interface {
	// Create creates a key-auth credential in Kong
	// If an ID is specified, it will be used to
	// create a key-auth in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumerUsernameOrID *string, keyAuth *KeyAuth) (*KeyAuth, error)
	// Get fetches a key-auth credential from Kong.
	Get(ctx context.Context, consumerUsernameOrID, keyOrID *string) (*KeyAuth, error)
	// Update updates a key-auth credential in Kong
	Update(ctx context.Context, consumerUsernameOrID *string, keyAuth *KeyAuth) (*KeyAuth, error)
	// Delete deletes a key-auth credential in Kong
	Delete(ctx context.Context, consumerUsernameOrID, keyOrID *string) error
	// List fetches a list of key-auth credentials in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*KeyAuth, *ListOpt, error)
	// ListAll fetches all key-auth credentials in Kong.
	// This method can take a while if there
	// a lot of key-auth credentials present.
	ListAll(ctx context.Context) ([]*KeyAuth, error)
	// ListForConsumer fetches a list of key-auth credentials
	// in Kong associated with a specific consumer.
	// opt can be used to control pagination.
	ListForConsumer(ctx context.Context, consumerUsernameOrID *string, opt *ListOpt) ([]*KeyAuth, *ListOpt, error)
}
