package schemas

type BadRequestResponse struct {
	Message string `json:"message"`
}

type NotFoundResponse struct {
	Message string `json:"message"`
}

type UnprocessableEntityResponse struct {
	Message []map[string]string `json:"message"`
}

type InternalServerErrorResponse struct {
	Error string `json:"error"`
}