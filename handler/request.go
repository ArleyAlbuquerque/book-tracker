package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Creating books

type CreatingBookRequest struct {
	Name  string `json:"name"`
	Autor string `json:"autor"`
	Pages int64  `json:"pages"`
}

func (r *CreatingBookRequest) Validate() error {
	if r.Autor == "" && r.Name == "" && r.Pages <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Autor == "" {
		return errParamIsRequired("autor", "string")
	}
	if r.Pages <= 0 {
		return errParamIsRequired("id", "int64")
	}

	return nil
}

type UpdateBookRequest struct {
	Name  string `json:"name"`
	Autor string `json:"autor"`
	Pages int64  `json:"pages"`
}

func (r *UpdateBookRequest) Validate() error {
	//if anyfield has been provided, validation is truthy
	if r.Autor != "" || r.Name != "" || r.Pages > 0 {
		return nil
	}
	//if none of fields has been provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
