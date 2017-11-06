package models

import (
	shared "github.com/FernandoCagale/go-api-shared/src/validation"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (t Task) Validate() (errors map[string][]shared.Validation, ok bool) {
	errors = make(map[string][]shared.Validation)

	if t.Description == "" {
		errors["description"] = append(errors["description"], shared.Validation{
			Type:    "required",
			Message: "field is required",
		})
	}

	if len(t.Description) > 40 {
		errors["description"] = append(errors["description"], shared.Validation{
			Type:    "lenght-max",
			Message: "field lenght max 40",
		})
	}

	if len(t.Description) < 5 {
		errors["description"] = append(errors["description"], shared.Validation{
			Type:    "lenght-min",
			Message: "field lenght min 5",
		})
	}

	if t.Type == "" {
		errors["type"] = append(errors["type"], shared.Validation{
			Type:    "required",
			Message: "field is required",
		})
	}

	return errors, len(errors) == 0
}
