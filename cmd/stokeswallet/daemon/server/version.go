package server

import (
	"context"
	"github.com/Sam-Stokes/stokes/cmd/kaspawallet/daemon/pb"
	"github.com/Sam-Stokes/stokes/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
