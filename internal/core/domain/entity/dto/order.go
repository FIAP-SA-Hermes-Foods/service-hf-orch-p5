package dto

import (
	"service-hf-orch-p5/internal/core/domain/entity"
	vo "service-hf-orch-p5/internal/core/domain/entity/valueObject"
)

type OrderDB struct {
	ID               int64             `json:"id,omitempty"`
	ClientUUID       string            `json:"clientUuid,omitempty"`
	VoucherUUID      string            `json:"voucherUuid,omitempty"`
	Items            []OutputOrderItem `json:"items,omitempty"`
	Status           string            `json:"status,omitempty"`
	VerificationCode string            `json:"verificationCode,omitempty"`
	CreatedAt        string            `json:"createdAt,omitempty"`
}

type (
	RequestOrder struct {
		ID               int64             `json:"id,omitempty"`
		ClientUUID       string            `json:"clientUuid,omitempty"`
		VoucherUUID      string            `json:"voucherUuid,omitempty"`
		Items            []OutputOrderItem `json:"items,omitempty"`
		Status           string            `json:"status,omitempty"`
		VerificationCode string            `json:"verificationCode,omitempty"`
		CreatedAt        string            `json:"createdAt,omitempty"`
	}

	OutputOrder struct {
		ID               int64             `json:"id,omitempty"`
		ClientUUID       string            `json:"clientUuid,omitempty"`
		VoucherUUID      string            `json:"voucherUuid,omitempty"`
		Items            []OutputOrderItem `json:"items,omitempty"`
		Status           string            `json:"status,omitempty"`
		VerificationCode string            `json:"verificationCode,omitempty"`
		CreatedAt        string            `json:"createdAt,omitempty"`
	}

	OutputOrderApp struct {
		ID               int64             `json:"id,omitempty"`
		Client           *OutputClient     `json:"client,omitempty"`
		Voucher          *OutputVoucher    `json:"voucher,omitempty"`
		Products         []OutputProduct   `json:"products,omitempty"`
		Items            []OutputOrderItem `json:"items,omitempty"`
		Status           string            `json:"status,omitempty"`
		TotalPrice       float64           `json:"totalPrice,omitempty"`
		VerificationCode string            `json:"verificationCode,omitempty"`
		CreatedAt        string            `json:"createdAt,omitempty"`
	}
)

func (r RequestOrder) Order() entity.Order {
	items := make([]entity.OrderItems, 0)
	for i := range r.Items {
		item := entity.OrderItems{
			ProductUUID: r.Items[i].ProductUUID,
			Quantity:    r.Items[i].Quantity,
			TotalPrice:  r.Items[i].TotalPrice,
			Discount:    r.Items[i].Discount,
		}
		items = append(items, item)
	}
	return entity.Order{
		ClientUUID:  r.ClientUUID,
		VoucherUUID: r.VoucherUUID,
		Items:       items,
		Status: vo.Status{
			Value: r.Status,
		},
	}
}
