package auth

type SendCodeRequest struct {
	Body SendCodeRequestBody
}

type SendCodeRequestBody struct {
	Email string `json:"email" format:"email" minLength:"3" maxLength:"255"`
}

type SendCodeResponse struct {
	Body SendCodeResponseBody
}

type SendCodeResponseBody struct {
	Message string `json:"message"`
}
