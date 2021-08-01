package courses

import (
	"bytes"
	"testing"

	"github.com/Serj1c/datalearn/api/pkg/data"
	"github.com/stretchr/testify/assert"
)

/* TODO tests to be reworked */

func TestCourseMissingTitleReturnsError(t *testing.T) {
	c := Course{
		Description: "dummies love SQL",
	}
	v := data.NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestCourseMissingAuthorReturnsError(t *testing.T) {
	c := Course{
		Title:       "SQL for dummies",
		Description: "dummies love SQL",
	}
	v := data.NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNotReturnError(t *testing.T) {
	c := Course{
		Title:       "SQL for dummies",
		Description: "dummies love SQL",
		Author:      "Sergei Isaev",
	}

	v := data.NewValidation()
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
	err := data.ToJSON(cs, b)
	assert.NoError(t, err)
}
