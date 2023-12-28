package main

import (
	"github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider"
	// "github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider/mock"
	"github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider/grpc"
)

// func registerMock(s *provider.Store) {
// 	/* #nosec */
// 	s.Register("mock", func(cfg provider.InitConfig) (provider.Provider, error) { //nolint:errcheck
// 		return mock.NewMockProvider(
// 			cfg.ConfigPath,
// 			cfg.NodeName,
// 			cfg.OperatingSystem,
// 			cfg.InternalIP,
// 			cfg.DaemonPort,
// 		)
// 	})
// }


func registerGrpc(s *provider.Store) {
	/* #nosec */
	s.Register("grpc", func(cfg provider.InitConfig) (provider.Provider, error) { //nolint:errcheck
		return grpc.NewGrpcProvider(
			cfg.ConfigPath,
		)
	})
}
