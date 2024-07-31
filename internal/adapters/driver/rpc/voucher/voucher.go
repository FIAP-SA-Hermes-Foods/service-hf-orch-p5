package voucher

import (
	"context"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
	"service-hf-orch-p5/internal/core/domain/rpc"
	op "service-hf-orch-p5/orch_proto"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ rpc.VoucherRPC = (*voucherRPC)(nil)

type voucherRPC struct {
	ctx  context.Context
	host string
	port string
}

func NewVoucherRPC(ctx context.Context, host, port string) rpc.VoucherRPC {
	return voucherRPC{ctx: ctx, host: host, port: port}
}

func (v voucherRPC) GetVoucherByID(id string) (*dto.OutputVoucher, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", v.host, v.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetVoucherByIDRequest{
		Uuid: id,
	}

	cc := op.NewVoucherClient(conn)

	resp, err := cc.GetVoucherByID(v.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputVoucher{
		ID:         resp.Uuid,
		Code:       resp.Code,
		Percentage: resp.Percentage,
		CreatedAt:  resp.CreatedAt,
		ExpiresAt:  resp.ExpiresAt,
	}

	return &out, nil
}

func (v voucherRPC) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", v.host, v.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.CreateVoucherRequest{
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		ExpiresAt:  voucher.ExpiresAt,
	}

	cc := op.NewVoucherClient(conn)

	resp, err := cc.CreateVoucher(v.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputVoucher{
		ID:         resp.Uuid,
		Code:       resp.Code,
		Percentage: resp.Percentage,
		CreatedAt:  resp.CreatedAt,
		ExpiresAt:  resp.ExpiresAt,
	}

	return &out, nil
}

func (v voucherRPC) UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", v.host, v.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.UpdateVoucherByIDRequest{
		Uuid:       id,
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		CreatedAt:  voucher.CreatedAt,
		ExpiresAt:  voucher.ExpiresAt,
	}

	cc := op.NewVoucherClient(conn)

	resp, err := cc.UpdateVoucherByID(v.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputVoucher{
		ID:         resp.Uuid,
		Code:       resp.Code,
		Percentage: resp.Percentage,
		CreatedAt:  resp.CreatedAt,
		ExpiresAt:  resp.ExpiresAt,
	}

	return &out, nil
}
