// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// ConsumerServiceInterface ...
type ConsumerServiceInterface interface {
	// Create creates a Consumer in Kong.
	// If an ID is specified, it will be used to
	// create a consumer in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumer *Consumer) (*Consumer, error)
	// Get fetches a Consumer in Kong.
	Get(ctx context.Context, usernameOrID *string) (*Consumer, error)
	// GetByCustomID fetches a Consumer in Kong.
	GetByCustomID(ctx context.Context, customID *string) (*Consumer, error)
	// Update updates a Consumer in Kong
	Update(ctx context.Context, consumer *Consumer) (*Consumer, error)
	// Delete deletes a Consumer in Kong
	Delete(ctx context.Context, usernameOrID *string) error
	// List fetches a list of Consumers in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Consumer, *ListOpt, error)
	// ListAll fetches all Consumers in Kong.
	// This method can take a while if there
	// a lot of Consumers present.
	ListAll(ctx context.Context) ([]*Consumer, error)
}
