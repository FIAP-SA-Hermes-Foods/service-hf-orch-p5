package product

import (
	"context"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
	"service-hf-orch-p5/internal/core/domain/rpc"
	op "service-hf-orch-p5/orch_proto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ rpc.ProductRPC = (*productRPC)(nil)

type productRPC struct {
	ctx  context.Context
	host string
	port string
}

func NewProductRPC(ctx context.Context, host, port string) rpc.ProductRPC {
	return productRPC{ctx: ctx, host: host, port: port}
}

func (p productRPC) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.CreateProductRequest{
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         float32(product.Price),
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.CreateProduct(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputProduct{
		UUID:          resp.Uuid,
		Name:          resp.Name,
		Category:      resp.Category,
		Image:         resp.Image,
		Description:   resp.Description,
		Price:         float64(resp.Price),
		CreatedAt:     resp.CreatedAt,
		DeactivatedAt: resp.DeactivatedAt,
	}

	return &out, nil
}

func (p productRPC) UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.UpdateProductRequest{
		Uuid:          id,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         float32(product.Price),
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.UpdateProduct(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputProduct{
		UUID:          resp.Uuid,
		Name:          resp.Name,
		Category:      resp.Category,
		Image:         resp.Image,
		Description:   resp.Description,
		Price:         float64(resp.Price),
		CreatedAt:     resp.CreatedAt,
		DeactivatedAt: resp.DeactivatedAt,
	}

	return &out, nil
}

func (p productRPC) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetProductByCategoryRequest{
		Category: category,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.GetProductByCategory(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	out := make([]dto.OutputProduct, 0)

	for i := range resp.Items {
		var outItem = dto.OutputProduct{
			UUID:          resp.Items[i].Uuid,
			Name:          resp.Items[i].Name,
			Category:      resp.Items[i].Category,
			Image:         resp.Items[i].Image,
			Description:   resp.Items[i].Description,
			Price:         float64(resp.Items[i].Price),
			CreatedAt:     resp.Items[i].CreatedAt,
			DeactivatedAt: resp.Items[i].DeactivatedAt,
		}

		out = append(out, outItem)
	}

	return out, nil
}

func (p productRPC) DeleteProductByID(id string) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	defer conn.Close()

	input := op.DeleteProductByUUIDRequest{
		Uuid: id,
	}

	cc := op.NewProductClient(conn)

	if _, err := cc.DeleteProductByUUID(p.ctx, &input); err != nil {
		return err
	}

	return nil
}
