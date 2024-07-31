package rpc

import (
	"service-hf-orch-p5/internal/core/domain/entity/dto"
)

type ClientRPC interface {
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	SaveClient(client dto.RequestClient) (*dto.OutputClient, error)
}

type OrderRPC interface {
	SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error)
	UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error)
	GetOrders() ([]dto.OutputOrder, error)
	GetOrderByID(id int64) (*dto.OutputOrder, error)
}

type ProductRPC interface {
	SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	DeleteProductByID(id string) error
}

type VoucherRPC interface {
	GetVoucherByID(id string) (*dto.OutputVoucher, error)
	SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}
