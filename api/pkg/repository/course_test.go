package repository

import (
	"bytes"
	"testing"

	"github.com/Serj1c/datalearn/api/pkg/entity"
	"github.com/Serj1c/datalearn/api/pkg/middleware"
	"github.com/Serj1c/datalearn/api/pkg/util"
	"github.com/stretchr/testify/assert"
)

/* TODO tests to be reworked */

func TestCourseMissingTitleReturnsError(t *testing.T) {
	c := entity.Course{
		Description: "dummies love SQL",
	}
	v := middleware.NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestCourseMissingAuthorReturnsError(t *testing.T) {
	c := entity.Course{
		Title:       "SQL for dummies",
		Description: "dummies love SQL",
	}
	v := middleware.NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNotReturnError(t *testing.T) {
	c := entity.Course{
		Title:       "SQL for dummies",
		Description: "dummies love SQL",
		Author:      "Sergei Isaev",
	}

	v := middleware.NewValidation()
	err := v.Validate(c)
	assert.Len(t, err, 1)
}

func TestCourseToJSON(t *testing.T) {
	cs := []*entity.Course{
		{
			Title: "SQL for dummies",
		},
	}
	b := bytes.NewBufferString("")
	err := util.ToJSON(cs, b)
	assert.NoError(t, err)
}
