package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourseMissingTitleReturnsError(t *testing.T) {
	c := Course{
		Description: "dummies love SQL",
	}
	v := NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestCourseMissingAuthorReturnsError(t *testing.T) {
	c := Course{
		Title:       "SQL for dummies",
		Description: "dummies love SQL",
	}
	v := NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNotReturnError(t *testing.T) {
	c := Course{
		Title:       "SQL for dummies",
		Description: "dummies love SQL",
		Author:      "Sergei Isaev",
	}

	v := NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestCourseToJSON(t *testing.T) {
	cs := []*Course{
		{
			Title: "SQL for dummies",
		},
	}
	b := bytes.NewBufferString("")
	err := ToJSON(cs, b)
	assert.NoError(t, err)
}
