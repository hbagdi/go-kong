// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// CertificateServiceInterface ...
type CertificateServiceInterface interface {
	// Create creates a Certificate in Kong.
	// If an ID is specified, it will be used to
	// create a certificate in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, certificate *Certificate) (*Certificate, error)
	// Get fetches a Certificate in Kong.
	Get(ctx context.Context, usernameOrID *string) (*Certificate, error)
	// Update updates a Certificate in Kong
	Update(ctx context.Context, certificate *Certificate) (*Certificate, error)
	// Delete deletes a Certificate in Kong
	Delete(ctx context.Context, usernameOrID *string) error
	// List fetches a list of certificate in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Certificate, *ListOpt, error)
	// ListAll fetches all Certificates in Kong.
	// This method can take a while if there
	// a lot of Certificates present.
	ListAll(ctx context.Context) ([]*Certificate, error)
}
