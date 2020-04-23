package kong

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// WorkspaceService handles Workspaces in Kong.
type WorkspaceService service

// Create creates a Workspace in Kong.
func (s *WorkspaceService) Create(ctx context.Context,
	workspace *Workspace) (*Workspace, error) {

	if workspace == nil {
		return nil, errors.New("cannot create a nil workspace")
	}

	endpoint := "/workspaces"
	method := "POST"
	if workspace.ID != nil {
		endpoint = endpoint + "/" + *workspace.ID
		method = "PUT"
	}
	req, err := s.client.NewRequest(method, endpoint, nil, workspace)

	if err != nil {
		return nil, err
	}

	var createdWorkspace Workspace
	_, err = s.client.Do(ctx, req, &createdWorkspace)
	if err != nil {
		return nil, err
	}
	return &createdWorkspace, nil
}

// Get fetches a Workspace in Kong.
func (s *WorkspaceService) Get(ctx context.Context,
	nameOrID *string) (*Workspace, error) {

	if isEmptyString(nameOrID) {
		return nil, errors.New("nameOrID cannot be nil for Get operation")
	}

	endpoint := fmt.Sprintf("/workspaces/%v", *nameOrID)
	req, err := s.client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var Workspace Workspace
	_, err = s.client.Do(ctx, req, &Workspace)
	if err != nil {
		return nil, err
	}
	return &Workspace, nil
}

// Update updates a Workspace in Kong. Only updates to the
// `comment` field are supported. To rename a workspace use Create.
func (s *WorkspaceService) Update(ctx context.Context,
	workspace *Workspace) (*Workspace, error) {

	if workspace == nil {
		return nil, errors.New("cannot update a nil Workspace")
	}

	if isEmptyString(workspace.ID) {
		return nil, errors.New("ID cannot be nil for Update operation")
	}

	endpoint := fmt.Sprintf("/workspaces/%v", *workspace.ID)
	req, err := s.client.NewRequest("PATCH", endpoint, nil, workspace)
	if err != nil {
		return nil, err
	}

	var updatedWorkspace Workspace
	_, err = s.client.Do(ctx, req, &updatedWorkspace)
	if err != nil {
		return nil, err
	}
	return &updatedWorkspace, nil
}

// Delete deletes a Workspace in Kong
func (s *WorkspaceService) Delete(ctx context.Context,
	WorkspaceOrID *string) error {

	if isEmptyString(WorkspaceOrID) {
		return errors.New("WorkspaceOrID cannot be nil for Delete operation")
	}

	endpoint := fmt.Sprintf("/workspaces/%v", *WorkspaceOrID)
	req, err := s.client.NewRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	return err
}

// List fetches a list of all Workspaces in Kong.
func (s *WorkspaceService) List(ctx context.Context,
	opt *ListOpt) ([]*Workspace, *ListOpt, error) {

	data, next, err := s.client.list(ctx, "/workspaces/", opt)
	if err != nil {
		return nil, nil, err
	}
	var workspaces []*Workspace
	for _, object := range data {
		b, err := object.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		var workspace Workspace
		err = json.Unmarshal(b, &workspace)
		if err != nil {
			return nil, nil, err
		}
		workspaces = append(workspaces, &workspace)
	}

	return workspaces, next, nil
}

// ListAll fetches all workspaces in Kong.
func (s *WorkspaceService) ListAll(ctx context.Context) ([]*Workspace, error) {

	var workspaces, data []*Workspace
	var err error
	opt := &ListOpt{Size: pageSize}

	for opt != nil {
		data, opt, err = s.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		workspaces = append(workspaces, data...)
	}

	return workspaces, nil
}
