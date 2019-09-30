// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// TargetServiceInterface ...
type TargetServiceInterface interface {
	// Create creates a Target in Kong under upstreamID.
	// If an ID is specified, it will be used to
	// create a target in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, upstreamNameOrID *string, target *Target) (*Target, error)
	// Delete deletes a Target in Kong
	Delete(ctx context.Context, upstreamNameOrID *string, targetOrID *string) error
	// List fetches a list of Targets in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, upstreamNameOrID *string, opt *ListOpt) ([]*Target, *ListOpt, error)
	// ListAll fetches all Targets in Kong for an upstream.
	ListAll(ctx context.Context, upstreamNameOrID *string) ([]*Target, error)
	// MarkHealthy marks target belonging to upstreamNameOrID as healthy in
	// Kong's load balancer.
	MarkHealthy(ctx context.Context, upstreamNameOrID *string, target *Target) error
	// MarkUnhealthy marks target belonging to upstreamNameOrID as unhealthy in
	// Kong's load balancer.
	MarkUnhealthy(ctx context.Context, upstreamNameOrID *string, target *Target) error
}
