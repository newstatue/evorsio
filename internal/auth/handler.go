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

	return &SendCodeResponse{Body: SendCodeResponseBody{Message: "OK"}}, nil
}
