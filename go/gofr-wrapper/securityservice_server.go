// versions:
// 	gofr-cli v0.6.0
// 	gofr.dev v1.37.0
// 	source: security-service.proto

package gofrwrapper

import (
	"gofr.dev/pkg/gofr"

	"github.com/stratifyr/security-service-proto/go/pb"
)

// Register the gRPC service in your app using the following code in your main.go:
//
// grpc.RegisterSecurityServiceServerWithGofr(app, &grpc.NewSecurityServiceGoFrServer())
//
// SecurityServiceGoFrServer defines the gRPC server implementation.
// Customize the struct with required dependencies and fields as needed.

type SecurityServiceGoFrServer struct {
	health *healthServer
}

func (s *SecurityServiceGoFrServer) GetMarketDays(ctx *gofr.Context) (any, error) {
	return &pb.GetMarketDaysResponse{}, nil
}
func (s *SecurityServiceGoFrServer) GetMetrics(ctx *gofr.Context) (any, error) {
	return &pb.GetMetricsResponse{}, nil
}
func (s *SecurityServiceGoFrServer) GetSecurities(ctx *gofr.Context) (any, error) {
	return &pb.GetSecuritiesResponse{}, nil
}
func (s *SecurityServiceGoFrServer) GetHistoricalData(ctx *gofr.Context) (any, error) {
	return &pb.GetHistoricalDataResponse{}, nil
}
func (s *SecurityServiceGoFrServer) UpsertSecurity(ctx *gofr.Context) (any, error) {
	return &pb.UpsertSecurityResponse{}, nil
}
func (s *SecurityServiceGoFrServer) UpsertSecurityStat(ctx *gofr.Context) (any, error) {
	return &pb.UpsertSecurityStatRequest{}, nil
}
func (s *SecurityServiceGoFrServer) GetMarketDataJobs(ctx *gofr.Context) (any, error) {
	return &pb.GetMarketDataJobsResponse{}, nil
}
func (s *SecurityServiceGoFrServer) GetIndices(ctx *gofr.Context) (any, error) {
	return &pb.GetIndicesResponse{}, nil
}
func (s *SecurityServiceGoFrServer) UpsertIndex(ctx *gofr.Context) (any, error) {
	return &pb.UpsertIndexResponse{}, nil
}
func (s *SecurityServiceGoFrServer) UpsertIndexStat(ctx *gofr.Context) (any, error) {
	return &pb.UpsertIndexStatResponse{}, nil
}
func (s *SecurityServiceGoFrServer) UpdateMarketDataJob(ctx *gofr.Context) (any, error) {
	return &pb.UpdateMarketDataJobResponse{}, nil
}
