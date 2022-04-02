package api

import (
	"context"

	pb "xinhari.com/hari/service/auth/api/proto"
	"xinhari.com/xinhari"
	"xinhari.com/xinhari/auth"
	"xinhari.com/xinhari/errors"
)

// Handler is an impementation of the auth api
type Handler struct {
	auth auth.Auth
}

// NewHandler returns an initialized Handler
func NewHandler(srv micro.Service) *Handler {
	return &Handler{auth: auth.DefaultAuth}
}

// Verify gets a token and verifies it with the auth package
func (h *Handler) Verify(ctx context.Context, req *pb.VerifyRequest, rsp *pb.VerifyResponse) error {
	if len(req.Token) == 0 {
		return errors.BadRequest("go.micro.api.auth", "token required")
	}

	_, err := h.auth.Inspect(req.Token)
	return err
}
