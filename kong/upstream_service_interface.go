// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// UpstreamServiceInterface ...
type UpstreamServiceInterface interface {
	// Create creates a Upstream in Kong.
	// If an ID is specified, it will be used to
	// create a upstream in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, upstream *Upstream) (*Upstream, error)
	// Get fetches a Upstream in Kong.
	Get(ctx context.Context, upstreamNameOrID *string) (*Upstream, error)
	// Update updates a Upstream in Kong
	Update(ctx context.Context, upstream *Upstream) (*Upstream, error)
	// Delete deletes a Upstream in Kong
	Delete(ctx context.Context, upstreamNameOrID *string) error
	// List fetches a list of Upstreams in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Upstream, *ListOpt, error)
	// ListAll fetches all Upstreams in Kong.
	// This method can take a while if there
	// a lot of Upstreams present.
	ListAll(ctx context.Context) ([]*Upstream, error)
}
