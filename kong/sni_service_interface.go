// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// SNIServiceInterface ...
type SNIServiceInterface interface {
	// Create creates a SNI in Kong.
	// If an ID is specified, it will be used to
	// create a sni in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, sni *SNI) (*SNI, error)
	// Get fetches a SNI in Kong.
	Get(ctx context.Context, usernameOrID *string) (*SNI, error)
	// Update updates a SNI in Kong
	Update(ctx context.Context, sni *SNI) (*SNI, error)
	// Delete deletes a SNI in Kong
	Delete(ctx context.Context, usernameOrID *string) error
	// List fetches a list of SNIs in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*SNI, *ListOpt, error)
	// ListForCertificate fetches a list of SNIs
	// in Kong associated with certificateID.
	// opt can be used to control pagination.
	ListForCertificate(ctx context.Context, certificateID *string, opt *ListOpt) ([]*SNI, *ListOpt, error)
	// ListAll fetches all SNIs in Kong.
	// This method can take a while if there
	// a lot of SNIs present.
	ListAll(ctx context.Context) ([]*SNI, error)
}
