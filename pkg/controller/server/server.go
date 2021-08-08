package controller

import (
	"context"

	"github.com/logand22/service-mesh-from-scratch/pkg/controller"
)

type Server struct {
	control *controller.Controller
}

func NewServer(ctx context.Context, control *controller.Controller) *Server {
	return &Server{control}
}
