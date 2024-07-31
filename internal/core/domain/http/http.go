package http

import (
	"context"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
)

type PaymentAPI interface {
	DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error)
}
