// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// ProjectsService handles communication with the project V2
// methods of the GitHub API.
//
// GitHub API docs: https://docs.github.com/rest/projects/projects
type ProjectsService service

// ProjectV2 represents a v2 project.
type ProjectV2 struct {
	ID               *int64     `json:"id,omitempty"`
	NodeID           *string    `json:"node_id,omitempty"`
	Owner            *User      `json:"owner,omitempty"`
	Creator          *User      `json:"creator,omitempty"`
	Title            *string    `json:"title,omitempty"`
	Description      *string    `json:"description,omitempty"`
	Public           *bool      `json:"public,omitempty"`
	ClosedAt         *Timestamp `json:"closed_at,omitempty"`
	CreatedAt        *Timestamp `json:"created_at,omitempty"`
	UpdatedAt        *Timestamp `json:"updated_at,omitempty"`
	DeletedAt        *Timestamp `json:"deleted_at,omitempty"`
	Number           *int       `json:"number,omitempty"`
	ShortDescription *string    `json:"short_description,omitempty"`
	DeletedBy        *User      `json:"deleted_by,omitempty"`

	// Fields migrated from the Project (classic) struct:
	URL                    *string `json:"url,omitempty"`
	HTMLURL                *string `json:"html_url,omitempty"`
	ColumnsURL             *string `json:"columns_url,omitempty"`
	OwnerURL               *string `json:"owner_url,omitempty"`
	Name                   *string `json:"name,omitempty"`
	Body                   *string `json:"body,omitempty"`
	State                  *string `json:"state,omitempty"`
	OrganizationPermission *string `json:"organization_permission,omitempty"`
	Private                *bool   `json:"private,omitempty"`
}

func (p ProjectV2) String() string { return Stringify(p) }

// ListProjectsPaginationOptions specifies optional parameters to list projects for user / organization.
//
// Note: Pagination is powered by before/after cursor-style pagination. After the initial call,
// inspect the returned *Response. Use resp.After as the opts.After value to request
// the next page, and resp.Before as the opts.Before value to request the previous
// page. Set either Before or After for a request; if both are
// supplied GitHub API will return an error. PerPage controls the number of items
// per page (max 100 per GitHub API docs).
type ListProjectsPaginationOptions struct {
	// A cursor, as given in the Link header. If specified, the query only searches for events before this cursor.
	Before *string `url:"before,omitempty"`

	// A cursor, as given in the Link header. If specified, the query only searches for events after this cursor.
	After *string `url:"after,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage *int `url:"per_page,omitempty"`
}

// ListProjectsOptions specifies optional parameters to list projects for user / organization.
type ListProjectsOptions struct {
	ListProjectsPaginationOptions

	// Q is an optional query string to limit results to projects of the specified type.
	Query *string `url:"q,omitempty"`
}

// ProjectV2FieldOption represents an option for a project field of type single_select or multi_select.
// It defines the available choices that can be selected for dropdown-style fields.
//
// GitHub API docs: https://docs.github.com/rest/projects/fields
type ProjectV2FieldOption struct {
	ID          *string `json:"id,omitempty"`          // The unique identifier for this option.
	Name        *string `json:"name,omitempty"`        // The display name of the option.
	Color       *string `json:"color,omitempty"`       // The color associated with this option (e.g., "blue", "red").
	Description *string `json:"description,omitempty"` // An optional description for this option.
}

// ProjectV2Field represents a field in a GitHub Projects V2 project.
// Fields define the structure and data types for project items.
//
// GitHub API docs: https://docs.github.com/rest/projects/fields
type ProjectV2Field struct {
	ID        *int64                  `json:"id,omitempty"`
	NodeID    *string                 `json:"node_id,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	DataType  *string                 `json:"dataType,omitempty"`
	URL       *string                 `json:"url,omitempty"`
	Options   []*ProjectV2FieldOption `json:"options,omitempty"`
	CreatedAt *Timestamp              `json:"created_at,omitempty"`
	UpdatedAt *Timestamp              `json:"updated_at,omitempty"`
}

// ListOrganizationProjects lists Projects V2 for an organization.
//
// GitHub API docs: https://docs.github.com/rest/projects/projects#list-projects-for-organization
//
//meta:operation GET /orgs/{org}/projectsV2
func (s *ProjectsService) ListOrganizationProjects(ctx context.Context, org string, opts *ListProjectsOptions) ([]*ProjectV2, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2", org)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var projects []*ProjectV2
	resp, err := s.client.Do(ctx, req, &projects)
	if err != nil {
		return nil, resp, err
	}
	return projects, resp, nil
}

// GetOrganizationProject gets a Projects V2 project for an organization by ID.
//
// GitHub API docs: https://docs.github.com/rest/projects/projects#get-project-for-organization
//
//meta:operation GET /orgs/{org}/projectsV2/{project_number}
func (s *ProjectsService) GetOrganizationProject(ctx context.Context, org string, projectNumber int) (*ProjectV2, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v", org, projectNumber)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(ProjectV2)
	resp, err := s.client.Do(ctx, req, project)
	if err != nil {
		return nil, resp, err
	}
	return project, resp, nil
}

// ListUserProjects lists Projects V2 for a user.
//
// GitHub API docs: https://docs.github.com/rest/projects/projects#list-projects-for-user
//
//meta:operation GET /users/{username}/projectsV2
func (s *ProjectsService) ListUserProjects(ctx context.Context, username string, opts *ListProjectsOptions) ([]*ProjectV2, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2", username)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var projects []*ProjectV2
	resp, err := s.client.Do(ctx, req, &projects)
	if err != nil {
		return nil, resp, err
	}
	return projects, resp, nil
}

// GetUserProject gets a Projects V2 project for a user by ID.
//
// GitHub API docs: https://docs.github.com/rest/projects/projects#get-project-for-user
//
//meta:operation GET /users/{username}/projectsV2/{project_number}
func (s *ProjectsService) GetUserProject(ctx context.Context, username string, projectNumber int) (*ProjectV2, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v", username, projectNumber)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(ProjectV2)
	resp, err := s.client.Do(ctx, req, project)
	if err != nil {
		return nil, resp, err
	}
	return project, resp, nil
}

// ListOrganizationProjectFields lists Projects V2 for an organization.
//
// GitHub API docs: https://docs.github.com/rest/projects/fields#list-project-fields-for-organization
//
//meta:operation GET /orgs/{org}/projectsV2/{project_number}/fields
func (s *ProjectsService) ListOrganizationProjectFields(ctx context.Context, org string, projectNumber int, opts *ListProjectsOptions) ([]*ProjectV2Field, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/fields", org, projectNumber)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var fields []*ProjectV2Field
	resp, err := s.client.Do(ctx, req, &fields)
	if err != nil {
		return nil, resp, err
	}
	return fields, resp, nil
}

// ListUserProjectFields lists Projects V2 for a user.
//
// GitHub API docs: https://docs.github.com/rest/projects/fields#list-project-fields-for-user
//
//meta:operation GET /users/{username}/projectsV2/{project_number}/fields
func (s *ProjectsService) ListUserProjectFields(ctx context.Context, user string, projectNumber int, opts *ListProjectsOptions) ([]*ProjectV2Field, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/fields", user, projectNumber)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var fields []*ProjectV2Field
	resp, err := s.client.Do(ctx, req, &fields)
	if err != nil {
		return nil, resp, err
	}
	return fields, resp, nil
}

// GetOrganizationProjectField gets a single project field from an organization owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/fields#get-project-field-for-organization
//
//meta:operation GET /orgs/{org}/projectsV2/{project_number}/fields/{field_id}
func (s *ProjectsService) GetOrganizationProjectField(ctx context.Context, org string, projectNumber int, fieldID int64) (*ProjectV2Field, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/fields/%v", org, projectNumber, fieldID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	field := new(ProjectV2Field)
	resp, err := s.client.Do(ctx, req, field)
	if err != nil {
		return nil, resp, err
	}
	return field, resp, nil
}

// GetUserProjectField gets a single project field from a user owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/fields#get-project-field-for-user
//
//meta:operation GET /users/{username}/projectsV2/{project_number}/fields/{field_id}
func (s *ProjectsService) GetUserProjectField(ctx context.Context, user string, projectNumber int, fieldID int64) (*ProjectV2Field, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/fields/%v", user, projectNumber, fieldID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	field := new(ProjectV2Field)
	resp, err := s.client.Do(ctx, req, field)
	if err != nil {
		return nil, resp, err
	}
	return field, resp, nil
}

// ListProjectItemsOptions specifies optional parameters when listing project items.
// Note: Pagination uses before/after cursor-style pagination similar to ListProjectsOptions.
// "Fields" can be used to restrict which field values are returned (by their numeric IDs).
type ListProjectItemsOptions struct {
	// Embed ListProjectsOptions to reuse pagination and query parameters.
	ListProjectsOptions
	// Fields restricts which field values are returned by numeric field IDs.
	Fields []int64 `url:"fields,omitempty,comma"`
}

// GetProjectItemOptions specifies optional parameters when getting a project item.
type GetProjectItemOptions struct {
	// Fields restricts which field values are returned by numeric field IDs.
	Fields []int64 `url:"fields,omitempty,comma"`
}

// AddProjectItemOptions represents the payload to add an item (issue or pull request)
// to a project. The Type must be either "Issue" or "PullRequest" (as per API docs) and
// ID is the numerical ID of that issue or pull request.
type AddProjectItemOptions struct {
	Type string `json:"type,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

// UpdateProjectItemOptions represents fields that can be modified for a project item.
// Currently the REST API allows archiving/unarchiving an item (archived boolean).
// This struct can be expanded in the future as the API grows.
type UpdateProjectItemOptions struct {
	// Archived indicates whether the item should be archived (true) or unarchived (false).
	Archived *bool `json:"archived,omitempty"`
	// Fields allows updating field values for the item. Each entry supplies a field ID and a value.
	Fields []*ProjectV2Field `json:"fields,omitempty"`
}

// ListOrganizationProjectItems lists items for an organization owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#list-items-for-an-organization-owned-project
//
//meta:operation GET /orgs/{org}/projectsV2/{project_number}/items
func (s *ProjectsService) ListOrganizationProjectItems(ctx context.Context, org string, projectNumber int, opts *ListProjectItemsOptions) ([]*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/items", org, projectNumber)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var items []*ProjectV2Item
	resp, err := s.client.Do(ctx, req, &items)
	if err != nil {
		return nil, resp, err
	}
	return items, resp, nil
}

// AddOrganizationProjectItem adds an issue or pull request item to an organization owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#add-item-to-organization-owned-project
//
//meta:operation POST /orgs/{org}/projectsV2/{project_number}/items
func (s *ProjectsService) AddOrganizationProjectItem(ctx context.Context, org string, projectNumber int, opts *AddProjectItemOptions) (*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/items", org, projectNumber)
	req, err := s.client.NewRequest("POST", u, opts)
	if err != nil {
		return nil, nil, err
	}

	item := new(ProjectV2Item)
	resp, err := s.client.Do(ctx, req, item)
	if err != nil {
		return nil, resp, err
	}
	return item, resp, nil
}

// GetOrganizationProjectItem gets a single item from an organization owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#get-an-item-for-an-organization-owned-project
//
//meta:operation GET /orgs/{org}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) GetOrganizationProjectItem(ctx context.Context, org string, projectNumber int, itemID int64, opts *GetProjectItemOptions) (*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/items/%v", org, projectNumber, itemID)
	req, err := s.client.NewRequest("GET", u, opts)
	if err != nil {
		return nil, nil, err
	}
	item := new(ProjectV2Item)
	resp, err := s.client.Do(ctx, req, item)
	if err != nil {
		return nil, resp, err
	}
	return item, resp, nil
}

// UpdateOrganizationProjectItem updates an item in an organization owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#update-project-item-for-organization
//
//meta:operation PATCH /orgs/{org}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) UpdateOrganizationProjectItem(ctx context.Context, org string, projectNumber int, itemID int64, opts *UpdateProjectItemOptions) (*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/items/%v", org, projectNumber, itemID)
	req, err := s.client.NewRequest("PATCH", u, opts)
	if err != nil {
		return nil, nil, err
	}
	item := new(ProjectV2Item)
	resp, err := s.client.Do(ctx, req, item)
	if err != nil {
		return nil, resp, err
	}
	return item, resp, nil
}

// DeleteOrganizationProjectItem deletes an item from an organization owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#delete-project-item-for-organization
//
//meta:operation DELETE /orgs/{org}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) DeleteOrganizationProjectItem(ctx context.Context, org string, projectNumber int, itemID int64) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/projectsV2/%v/items/%v", org, projectNumber, itemID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// ListUserProjectItems lists items for a user owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#list-items-for-a-user-owned-project
//
//meta:operation GET /users/{username}/projectsV2/{project_number}/items
func (s *ProjectsService) ListUserProjectItems(ctx context.Context, username string, projectNumber int, opts *ListProjectItemsOptions) ([]*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/items", username, projectNumber)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var items []*ProjectV2Item
	resp, err := s.client.Do(ctx, req, &items)
	if err != nil {
		return nil, resp, err
	}
	return items, resp, nil
}

// AddUserProjectItem adds an issue or pull request item to a user owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#add-item-to-user-owned-project
//
//meta:operation POST /users/{username}/projectsV2/{project_number}/items
func (s *ProjectsService) AddUserProjectItem(ctx context.Context, username string, projectNumber int, opts *AddProjectItemOptions) (*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/items", username, projectNumber)
	req, err := s.client.NewRequest("POST", u, opts)
	if err != nil {
		return nil, nil, err
	}
	item := new(ProjectV2Item)
	resp, err := s.client.Do(ctx, req, item)
	if err != nil {
		return nil, resp, err
	}
	return item, resp, nil
}

// GetUserProjectItem gets a single item from a user owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#get-an-item-for-a-user-owned-project
//
//meta:operation GET /users/{username}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) GetUserProjectItem(ctx context.Context, username string, projectNumber int, itemID int64, opts *GetProjectItemOptions) (*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/items/%v", username, projectNumber, itemID)
	req, err := s.client.NewRequest("GET", u, opts)
	if err != nil {
		return nil, nil, err
	}
	item := new(ProjectV2Item)
	resp, err := s.client.Do(ctx, req, item)
	if err != nil {
		return nil, resp, err
	}
	return item, resp, nil
}

// UpdateUserProjectItem updates an item in a user owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#update-project-item-for-user
//
//meta:operation PATCH /users/{username}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) UpdateUserProjectItem(ctx context.Context, username string, projectNumber int, itemID int64, opts *UpdateProjectItemOptions) (*ProjectV2Item, *Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/items/%v", username, projectNumber, itemID)
	req, err := s.client.NewRequest("PATCH", u, opts)
	if err != nil {
		return nil, nil, err
	}
	item := new(ProjectV2Item)
	resp, err := s.client.Do(ctx, req, item)
	if err != nil {
		return nil, resp, err
	}
	return item, resp, nil
}

// DeleteUserProjectItem deletes an item from a user owned project.
//
// GitHub API docs: https://docs.github.com/rest/projects/items#delete-project-item-for-user
//
//meta:operation DELETE /users/{username}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) DeleteUserProjectItem(ctx context.Context, username string, projectNumber int, itemID int64) (*Response, error) {
	u := fmt.Sprintf("users/%v/projectsV2/%v/items/%v", username, projectNumber, itemID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
