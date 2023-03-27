package handler

import "fmt"

func throwNewEmptyError(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return throwNewEmptyError("role", "string")
	}
	if r.Company == "" {
		return throwNewEmptyError("company", "string")
	}
	if r.Location == "" {
		return throwNewEmptyError("location", "string")
	}
	if r.Link == "" {
		return throwNewEmptyError("link", "string")
	}
	if r.Remote == nil {
		return throwNewEmptyError("remote", "bool")
	}
	if r.Salary <= 0 {
		return throwNewEmptyError("salary", "int64")
	}

	return nil
}
