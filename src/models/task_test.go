package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOk(t *testing.T) {
	// Setup
	task := Task{
		Id:          1,
		Description: "Golang",
		Type:        "Dev",
	}

	errors, valid := task.Validate()

	// Assertions
	assert.Equal(t, valid, true)
	assert.Equal(t, len(errors), 0)
}

func TestValidateNot(t *testing.T) {
	// Setup
	task := Task{
		Id:          1,
		Description: "Go",
		Type:        "Dev",
	}

	errors, valid := task.Validate()

	// Assertions
	assert.Equal(t, valid, false)
	assert.Equal(t, len(errors), 1)
}
