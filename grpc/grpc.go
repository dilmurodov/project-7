package grpc

import (
	"paynet/pb_installment_service/config"
	"paynet/pb_installment_service/genproto/installment_service"
	"paynet/pb_installment_service/grpc/client"
	"paynet/pb_installment_service/grpc/service"
	"paynet/pb_installment_service/pkg/logger"
	"paynet/pb_installment_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	installment_service.RegisterInstallmentServiceServer(grpcServer, service.NewInstallmentService(cfg, log, strg, svcs))
	installment_service.RegisterCustomerServiceServer(grpcServer, service.NewCustomerService(cfg, log, strg, svcs))
	// object_service.RegisterObjectServiceServer(grpcServer, service.NewObjectService(cfg, log, strg, svcs))
	// object_service.RegisterRegulationServiceServer(grpcServer, service.NewRegulationService(cfg, log, strg, svcs))
	// object_service.RegisterViolationServiceServer(grpcServer, service.NewViolationService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
