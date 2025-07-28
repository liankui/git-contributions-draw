package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/chaos-io/chaos/logs"
	gitDraw "github.com/liankui/git-contributions-draw/go/git-contributions-draw"
	"github.com/liankui/git-contributions-draw/service/draw"

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
	if in == nil || in.Name == "" {
		logs.Errorw("failed to get request", "in", in)
		return nil, logs.NewErrorw("failed to get request", "in", in)
	}

	pattern := draw.GeneratePatternFromText(in.Name)
	for _, line := range pattern {
		fmt.Println(line)
	}

	year := time.Now().Year()
	thisYear := draw.GetYear(year)
	lastYear := draw.GetYear(year - 1)
	years := []*gitDraw.Year{thisYear, lastYear}

	var conts []*gitDraw.Contribution
	contributions := draw.GetContributions(year, pattern)
	lastContributions := draw.GetContributions(year-1, pattern)
	conts = append(conts, contributions...)
	conts = append(conts, lastContributions...)

	resp := &pb.GetDrawResponse{
		Years:         years,
		Contributions: conts,
	}
	return resp, nil
}
