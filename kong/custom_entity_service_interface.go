// DO NOT EDIT: Auto generated

package kong

import (
	"context"

	"github.com/hbagdi/go-kong/kong/custom"
)

// CustomEntityServiceInterface ...
type CustomEntityServiceInterface interface {
	// Get fetches a custom entity. The primary key and all relations of the
	// entity must be populated in entity.
	Get(ctx context.Context, entity custom.Entity) (custom.Entity, error)
	// Create creates a custom entity based on entity.
	// All required fields must be present in entity.
	Create(ctx context.Context, entity custom.Entity) (custom.Entity, error)
	// Update updates a custom entity in Kong.
	Update(ctx context.Context, entity custom.Entity) (custom.Entity, error)
	// Delete deletes a custom entity in Kong.
	Delete(ctx context.Context, entity custom.Entity) error
	// List fetches all custom entities based on relations
	List(ctx context.Context, opt *ListOpt, entity custom.Entity) ([]custom.Entity, *ListOpt, error)
	// ListAll fetches all custom entities based on relations
	ListAll(ctx context.Context, entity custom.Entity) ([]custom.Entity, error)
}
