// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// Oauth2ServiceInterface ...
type Oauth2ServiceInterface interface {
	// Create creates an oauth2 credential in Kong
	// If an ID is specified, it will be used to
	// create a oauth2 credential in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumerUsernameOrID *string, oauth2Cred *Oauth2Credential) (*Oauth2Credential, error)
	// Get fetches an oauth2 credential from Kong.
	Get(ctx context.Context, consumerUsernameOrID, clientIDorID *string) (*Oauth2Credential, error)
	// Update updates an oauth2 credential in Kong.
	Update(ctx context.Context, consumerUsernameOrID *string, oauth2Cred *Oauth2Credential) (*Oauth2Credential, error)
	// Delete deletes an oauth2 credential in Kong.
	Delete(ctx context.Context, consumerUsernameOrID, clientIDorID *string) error
	// List fetches a list of oauth2 credentials in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Oauth2Credential, *ListOpt, error)
	// ListAll fetches all oauth2 credentials in Kong.
	// This method can take a while if there
	// a lot of oauth2 credentials present.
	ListAll(ctx context.Context) ([]*Oauth2Credential, error)
	// ListForConsumer fetches a list of oauth2 credentials
	// in Kong associated with a specific consumer.
	// opt can be used to control pagination.
	ListForConsumer(ctx context.Context, consumerUsernameOrID *string, opt *ListOpt) ([]*Oauth2Credential, *ListOpt, error)
}
