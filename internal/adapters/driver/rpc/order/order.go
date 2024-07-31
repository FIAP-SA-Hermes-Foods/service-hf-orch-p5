package order

import (
	"context"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
	"service-hf-orch-p5/internal/core/domain/rpc"
	op "service-hf-orch-p5/orch_proto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ rpc.OrderRPC = (*orderRPC)(nil)

type orderRPC struct {
	ctx  context.Context
	host string
	port string
}

func NewOrderRPC(ctx context.Context, host, port string) rpc.OrderRPC {
	return orderRPC{ctx: ctx, host: host, port: port}
}

func (o orderRPC) SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", o.host, o.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	items := make([]*op.Item, 0)

	for i := range order.Items {
		item := op.Item{
			ProductUuid: order.Items[i].ProductUUID,
			Quantity:    order.Items[i].Quantity,
		}

		items = append(items, &item)
	}

	input := op.CreateOrderRequest{
		ClientUuid:  order.ClientUUID,
		VoucherUuid: order.VoucherUUID,
		Items:       items,
	}

	cc := op.NewOrderClient(conn)

	resp, err := cc.CreateOrder(o.ctx, &input)

	if err != nil {
		return nil, err
	}

	outItems := make([]dto.OutputOrderItem, 0)

	for i := range resp.Items {
		item := dto.OutputOrderItem{
			OrderID:     resp.Items[i].OrderId,
			ProductUUID: resp.Items[i].ProductUuid,
			Quantity:    resp.Items[i].Quantity,
		}

		outItems = append(outItems, item)
	}

	var out = dto.OutputOrder{
		ID:               resp.Id,
		ClientUUID:       resp.ClientUuid,
		VoucherUUID:      resp.VoucherUuid,
		Items:            outItems,
		Status:           resp.Status,
		VerificationCode: resp.VerificationCode,
		CreatedAt:        resp.CreatedAt,
	}

	return &out, nil

}
func (o orderRPC) UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", o.host, o.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	items := make([]*op.Item, 0)

	for i := range order.Items {
		item := op.Item{
			ProductUuid: order.Items[i].ProductUUID,
			Quantity:    order.Items[i].Quantity,
		}

		items = append(items, &item)
	}

	input := op.UpdateOrderRequest{
		Id:               id,
		ClientUuid:       order.ClientUUID,
		VoucherUuid:      order.VoucherUUID,
		Items:            items,
		Status:           order.Status,
		VerificationCode: order.Status,
		CreatedAt:        order.CreatedAt,
	}

	cc := op.NewOrderClient(conn)

	resp, err := cc.UpdateOrder(o.ctx, &input)

	if err != nil {
		return nil, err
	}

	outItems := make([]dto.OutputOrderItem, 0)

	for i := range resp.Items {
		item := dto.OutputOrderItem{
			OrderID:     resp.Items[i].OrderId,
			ProductUUID: resp.Items[i].ProductUuid,
			Quantity:    resp.Items[i].Quantity,
		}

		outItems = append(outItems, item)
	}

	var out = dto.OutputOrder{
		ID:               resp.Id,
		ClientUUID:       resp.ClientUuid,
		VoucherUUID:      resp.VoucherUuid,
		Items:            outItems,
		Status:           resp.Status,
		VerificationCode: resp.VerificationCode,
		CreatedAt:        resp.CreatedAt,
	}

	return &out, nil

}

func (o orderRPC) GetOrders() ([]dto.OutputOrder, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", o.host, o.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetOrderRequest{}

	cc := op.NewOrderClient(conn)

	resp, err := cc.GetOrder(o.ctx, &input)

	if err != nil {
		return nil, err
	}

	out := make([]dto.OutputOrder, 0)
	for orderIdx := range resp.Orders {
		outItems := make([]dto.OutputOrderItem, 0)
		for i := range resp.Orders[orderIdx].Items {
			item := dto.OutputOrderItem{
				OrderID:     resp.Orders[orderIdx].Items[i].OrderId,
				ProductUUID: resp.Orders[orderIdx].Items[i].ProductUuid,
				Quantity:    resp.Orders[orderIdx].Items[i].Quantity,
			}

			outItems = append(outItems, item)
		}

		outItem := dto.OutputOrder{
			ID:               resp.Orders[orderIdx].Id,
			ClientUUID:       resp.Orders[orderIdx].ClientUuid,
			VoucherUUID:      resp.Orders[orderIdx].VoucherUuid,
			Items:            outItems,
			Status:           resp.Orders[orderIdx].Status,
			VerificationCode: resp.Orders[orderIdx].VerificationCode,
			CreatedAt:        resp.Orders[orderIdx].CreatedAt,
		}

		out = append(out, outItem)
	}

	return out, nil

}

func (o orderRPC) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", o.host, o.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetOrderByIDRequest{Id: id}

	cc := op.NewOrderClient(conn)

	resp, err := cc.GetOrderByID(o.ctx, &input)

	if err != nil {
		return nil, err
	}

	outItems := make([]dto.OutputOrderItem, 0)

	for i := range resp.Items {
		item := dto.OutputOrderItem{
			OrderID:     resp.Items[i].OrderId,
			ProductUUID: resp.Items[i].ProductUuid,
			Quantity:    resp.Items[i].Quantity,
		}

		outItems = append(outItems, item)
	}

	var out = dto.OutputOrder{
		ID:               resp.Id,
		ClientUUID:       resp.ClientUuid,
		VoucherUUID:      resp.VoucherUuid,
		Items:            outItems,
		Status:           resp.Status,
		VerificationCode: resp.VerificationCode,
		CreatedAt:        resp.CreatedAt,
	}

	return &out, nil

}
