package auth

import (
	"context"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SendCode(ctx context.Context, input *SendCodeRequest) (*SendCodeResponse, error) {
	err := h.service.SendCode(ctx, input.Body.Email)
	if err != nil {
		return nil, err
	}

	return &SendCodeResponse{Body: SendCodeResponseBody{Message: "Send verification code successfully"}}, nil
}

func (h *Handler) Login(ctx context.Context, input *LoginRequest) (*LoginResponse, error) {
	token, err := h.service.LoginAndReturnToken(ctx, input.Body.Email, input.Body.Code)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Body: LoginResponseBody{Token: token}}, nil
}
