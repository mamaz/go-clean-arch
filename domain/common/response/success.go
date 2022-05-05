package response

import "github.com/go-chi/render"

var successMeta = Metadata{
	Code:    "SUCCESS",
	Message: "Success",
	Errors:  nil,
}

func SuccessResponse(data render.Renderer) AppResponse {
	return AppResponse{
		Metadata: successMeta,
		Data:     data,
	}
}

func SuccessResponseList(data []render.Renderer) AppResponseList {
	return AppResponseList{
		Metadata: successMeta,
		Data:     data,
	}
}
