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

type LoginRequest struct {
	Body LoginRequestBody
}

type LoginRequestBody struct {
	Email string `json:"email" format:"email" minLength:"3" maxLength:"255"`
	Code  string `json:"code" pattern:"^[0-9]{5}$" patternDescription:"verification code must be 5-digit number"`
}

type LoginResponse struct {
	Body LoginResponseBody
}

type LoginResponseBody struct {
	Token string `json:"token"`
}
