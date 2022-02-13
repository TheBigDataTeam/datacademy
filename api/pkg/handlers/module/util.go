package module

import "github.com/Serj1c/datalearn/api/pkg/entity"

func entityFromCreateRequest(req *CreateRequest) entity.Module {
	return entity.Module{
		CourseId: req.CourseId,
		Title:    req.Title,
		Body:     req.Body,
		Badge:    req.Badge,
	}
}
