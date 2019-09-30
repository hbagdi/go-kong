// DO NOT EDIT: Auto generated

package kong

import (
	"context"
	"encoding/json"
)

// credentialServiceInterface ...
type credentialServiceInterface interface {
	// Create creates a credential in Kong of type credType.
	// If an ID is specified in the credential, it will be used to
	// create a credential in Kong, otherwise an ID
	// is auto-generated.
	Create(ctx context.Context, credType string, consumerUsernameOrID *string, credential interface{}) (json.RawMessage, error)
	// Get fetches a credential of credType with credIdentifier from Kong.
	Get(ctx context.Context, credType string, consumerUsernameOrID *string, credIdentifier *string) (json.RawMessage, error)
	// Update updates credential in Kong
	Update(ctx context.Context, credType string, consumerUsernameOrID *string, credential interface{}) (json.RawMessage, error)
	// Delete deletes a credential in Kong
	Delete(ctx context.Context, credType string, consumerUsernameOrID, credIdentifier *string) error
}
