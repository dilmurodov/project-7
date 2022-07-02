package service

import (
	"context"
	"paynet/pb_installment_service/config"
	"paynet/pb_installment_service/genproto/installment_service"
	"paynet/pb_installment_service/grpc/client"
	"paynet/pb_installment_service/pkg/logger"
	"paynet/pb_installment_service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type customerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	installment_service.UnimplementedCustomerServiceServer
}

func NewCustomerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *customerService {
	return &customerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (c *customerService) CreateCustomer(ctx context.Context, req *installment_service.CreateCustomerRequest) (resp *installment_service.CreateCustomerResponse, err error) {

	c.log.Info("---CreateCustomer--->", logger.Any("req", req))

	resp, err = c.strg.Customer().CreateCustomer(ctx, req)

	if err != nil {
		c.log.Error("!!!CreateCustomer->Customer->CreateCustomer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) GetCustomer(ctx context.Context, req *installment_service.GetCustomerRequest) (resp *installment_service.GetCustomerResponse, err error) {

	c.log.Info("---GetCustomer------>", logger.Any("req", req))

	resp, err = c.strg.Customer().GetCustomer(ctx, req)

	if err != nil {
		c.log.Error("!!!GetCustomer->Customer->GetCustomer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) GetAllCustomers(ctx context.Context, req *installment_service.GetAllCustomersRequest) (resp *installment_service.GetAllCustomersResponse, err error) {

	c.log.Info("---GetAllCustomers------>", logger.Any("req", req))

	resp, err = c.strg.Customer().GetAllCustomers(ctx, req)

	if err != nil {
		c.log.Error("!!!GetAllCustomers->Customer->GetAllCustomers--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) DeleteCustomer(ctx context.Context, req *installment_service.DeleteCustomerRequest) (resp *empty.Empty, err error) {

	c.log.Info("---DeleteCustomer------>", logger.Any("req", req))

	resp, err = c.strg.Customer().DeleteCustomer(ctx, req)

	if err != nil {
		c.log.Error("!!!DeleteCustomer->Customer->DeleteCustomer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) AddCustomerAdditionalContacts(ctx context.Context, req *installment_service.AddCustomerAdditionalContactsRequest) (resp *empty.Empty, err error) {
	c.log.Info("---AddCustomerAddionalContacts---->", logger.Any("req", req))

	resp, err = c.strg.Customer().AddCustomerAdditionalContacts(ctx, req)

	if err != nil {
		c.log.Error("!!!AddCustomerAdditionalContacts--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) GetCustomerAdditionalContacts(ctx context.Context, req *installment_service.GetCustomerAdditionalContactsRequest) (resp *installment_service.GetCustomerAdditionalContactsResponse, err error) {
	c.log.Info("---GetCustomerAddionalContacts---->", logger.Any("req", req))

	resp, err = c.strg.Customer().GetCustomerAdditionalContacts(ctx, req)

	if err != nil {
		c.log.Error("!!!GetCustomerAdditionalContacts--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) AddCustomerAddresses(ctx context.Context, req *installment_service.AddCustomerAddressRequest) (resp *empty.Empty, err error) {

	c.log.Info("---AddCustomerAddressess---->", logger.Any("req", req))

	resp, err = c.strg.Customer().AddCustomerAddresses(ctx, req)

	if err != nil {
		c.log.Error("!!!AddCustomerAddressess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) GetCustomerAddresses(ctx context.Context, req *installment_service.GetCustomerAddressesRequest) (resp *installment_service.GetCustomerAddressesResponse, err error) {
	c.log.Info("---GetCustomerAddressess---->", logger.Any("req", req))

	resp, err = c.strg.Customer().GetCustomerAddresses(ctx, req)

	if err != nil {
		c.log.Error("!!!GetCustomerAddressess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) AddCustomerPersonalData(ctx context.Context, req *installment_service.AddCustomerPersonalDataRequest) (resp *empty.Empty, err error) {
	c.log.Info("---AddCustomerPersonalData---->", logger.Any("req", req))

	resp, err = c.strg.Customer().AddCustomerPersonalData(ctx, req)

	if err != nil {
		c.log.Error("!!!AddCustomerPersonalData--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) GetCustomerPersonalData(ctx context.Context, req *installment_service.GetCustomerPersonalDataRequest) (resp *installment_service.GetCustomerPersonalDataResponse, err error) {
	c.log.Info("---GetCustomerPersonalData---->", logger.Any("req", req))

	resp, err = c.strg.Customer().GetCustomerPersonalData(ctx, req)

	if err != nil {
		c.log.Error("!!!GetCustomerPersonalData--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) TestUpdateCustomerPersonalData(ctx context.Context, req *installment_service.UpdateCustomerPersonalDataRequest) (resp *installment_service.CustomerPersonalData, err error) {
	c.log.Info("---UpdateCustomerPersonalData---->", logger.Any("req", req))

	resp, err = c.strg.Customer().UpdateCustomerPersonalData(ctx, req)

	if err != nil {
		c.log.Error("!!!UpdateCustomerPersonalData--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (c *customerService) AddCustomerPassportData(ctx context.Context, req *installment_service.AddCustomerPassportDataRequest) (resp *installment_service.AddCustomerPassportDataResponse, err error) {
	c.log.Info("---AddCustomerPassportData---->", logger.Any("req", req))

	resp, err = c.strg.Customer().AddCustomerPassportData(ctx, req)

	if err != nil {
		c.log.Error("!!!AddCustomerPassportData--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *customerService) GetCustomerPassportData(ctx context.Context, req *installment_service.GetCustomerPassportDataRequest) (resp *installment_service.GetCustomerPassportDataResponse, err error) {
	c.log.Info("---GetCustomerPassportData---->", logger.Any("req", req))

	resp, err = c.strg.Customer().GetCustomerPassportData(ctx, req)

	if err != nil {
		c.log.Error("!!!GetCustomerPassportData--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}
