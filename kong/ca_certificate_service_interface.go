// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// CACertificateServiceInterface ...
type CACertificateServiceInterface interface {
	// Create creates a CACertificate in Kong.
	// If an ID is specified, it will be used to
	// create a certificate in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, certificate *CACertificate) (*CACertificate, error)
	// Get fetches a CACertificate in Kong.
	Get(ctx context.Context, ID *string) (*CACertificate, error)
	// Update updates a CACertificate in Kong
	Update(ctx context.Context, certificate *CACertificate) (*CACertificate, error)
	// Delete deletes a CACertificate in Kong
	Delete(ctx context.Context, ID *string) error
	// List fetches a list of certificate in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*CACertificate, *ListOpt, error)
	// ListAll fetches all Certificates in Kong.
	// This method can take a while if there
	// a lot of Certificates present.
	ListAll(ctx context.Context) ([]*CACertificate, error)
}
