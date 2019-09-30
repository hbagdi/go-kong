// DO NOT EDIT: Auto generated

package kong

import (
	"context"
)

// ACLServiceInterface ...
type ACLServiceInterface interface {
	// Create adds a consumer to an ACL group in Kong
	// If an ID is specified, it will be used to
	// create the group association in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, consumerUsernameOrID *string, aclGroup *ACLGroup) (*ACLGroup, error)
	// Get fetches an ACL group for a consumer in Kong.
	Get(ctx context.Context, consumerUsernameOrID, groupOrID *string) (*ACLGroup, error)
	// Update updates an ACL group for a consumer in Kong
	Update(ctx context.Context, consumerUsernameOrID *string, aclGroup *ACLGroup) (*ACLGroup, error)
	// Delete deletes an ACL group association for a consumer in Kong
	Delete(ctx context.Context, consumerUsernameOrID, groupOrID *string) error
	// List fetches a list of all ACL group and consumer associations in Kong.
	// opt can be used to control pagination.
	List(ctx context.Context, opt *ListOpt) ([]*ACLGroup, *ListOpt, error)
	// ListAll fetches all all ACL group associations in Kong.
	// This method can take a while if there
	// a lot of ACLGroup associations are present.
	ListAll(ctx context.Context) ([]*ACLGroup, error)
	// ListForConsumer fetches a list of ACL groups
	// in Kong associated with a specific consumer.
	// opt can be used to control pagination.
	ListForConsumer(ctx context.Context, consumerUsernameOrID *string, opt *ListOpt) ([]*ACLGroup, *ListOpt, error)
}
