package response

func ErrBadRequest(err error) AppResponse {
	return AppResponse{
		Metadata: Metadata{
			Code:    "BAD_REQUEST",
			Message: "Bad Request",
			Errors:  []string{err.Error()},
		},
		Data: nil,
	}
}

func ErrInternalServer() AppResponse {
	return AppResponse{
		Metadata: Metadata{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Internal Server Error",
			Errors:  []string{"Please contact support"},
		},
		Data: nil,
	}
}
