package handlers

import (
	"context"

	"github.com/chaos-io/chaos/logs"

	// this service api
	pb "github.com/liankui/git-contributions-draw/go/git-contributions-draw/v1"
)

type gitContributionsDrawServer struct {
	pb.UnimplementedGitContributionsDrawServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.GitContributionsDrawServer {
	return gitContributionsDrawServer{}
}

// GetDraw implements Interface.
func (s gitContributionsDrawServer) GetDraw(ctx context.Context, in *pb.GetDrawRequest) (*pb.GetDrawResponse, error) {
	logs.Infow("GetDraw", "request", in)

	resp := &pb.GetDrawResponse{
		// Years:
		// Contributions:
	}
	return resp, nil
}
