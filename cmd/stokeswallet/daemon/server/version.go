package server

import (
	"context"
	"github.com/stokesnetwork/stokes/cmd/stokeswallet/daemon/pb"
	"github.com/stokesnetwork/stokes/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
