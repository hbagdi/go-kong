// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// PluginServiceInterface ...
type PluginServiceInterface interface {
	// Create creates a Plugin in Kong.
	// If an ID is specified, it will be used to
	// create a plugin in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, plugin *Plugin) (*Plugin, error)
	// Get fetches a Plugin in Kong.
	Get(ctx context.Context, usernameOrID *string) (*Plugin, error)
	// Update updates a Plugin in Kong
	Update(ctx context.Context, plugin *Plugin) (*Plugin, error)
	// Delete deletes a Plugin in Kong
	Delete(ctx context.Context, usernameOrID *string) error
	// List fetches a list of Plugins in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*Plugin, *ListOpt, error)
	// ListAll fetches all Plugins in Kong.
	// This method can take a while if there
	// a lot of Plugins present.
	ListAll(ctx context.Context) ([]*Plugin, error)
	// ListAllForConsumer fetches all Plugins in Kong enabled for a consumer.
	ListAllForConsumer(ctx context.Context, consumerIDorName *string) ([]*Plugin, error)
	// ListAllForService fetches all Plugins in Kong enabled for a service.
	ListAllForService(ctx context.Context, serviceIDorName *string) ([]*Plugin, error)
	// ListAllForRoute fetches all Plugins in Kong enabled for a service.
	ListAllForRoute(ctx context.Context, routeID *string) ([]*Plugin, error)
}
