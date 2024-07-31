package entity

import (
	vo "service-hf-orch-p5/internal/core/domain/entity/valueObject"
)

type (
	Order struct {
		ID               int64               `json:"id,omitempty"`
		ClientUUID       string              `json:"clientUuid,omitempty"`
		VoucherUUID      string              `json:"voucherUuid,omitempty"`
		Items            []OrderItems        `json:"items,omitempty"`
		Status           vo.Status           `json:"status,omitempty"`
		VerificationCode vo.VerificationCode `json:"verificationCode,omitempty"`
		CreatedAt        vo.CreatedAt        `json:"createdAt,omitempty"`
	}

	OrderItems struct {
		ProductUUID string  `json:"productUuid,omitempty"`
		Quantity    int64   `json:"quantity,omitempty"`
		TotalPrice  float64 `json:"totalPrice,omitempty"`
		Discount    float64 `json:"discount,omitempty"`
	}
)
