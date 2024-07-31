package application

import (
	"context"
	"errors"
	l "service-hf-orch-p5/external/logger"
	ps "service-hf-orch-p5/external/strings"
	"service-hf-orch-p5/internal/core/domain/entity"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
	vo "service-hf-orch-p5/internal/core/domain/entity/valueObject"

	httpi "service-hf-orch-p5/internal/core/domain/http"
	"service-hf-orch-p5/internal/core/domain/rpc"
	"fmt"
	"strings"
)

type Application interface {
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	SaveClient(client dto.RequestClient) (*dto.OutputClient, error)

	SaveOrder(order dto.RequestOrder) (*dto.OutputOrderApp, error)
	UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrderApp, error)
	GetOrders() ([]dto.OutputOrderApp, error)
	GetOrderByID(id int64) (*dto.OutputOrderApp, error)

	SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	DeleteProductByID(id string) error

	GetVoucherByID(id string) (*dto.OutputVoucher, error)
	SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}

type application struct {
	ctx        context.Context
	clientRPC  rpc.ClientRPC
	orderRPC   rpc.OrderRPC
	productRPC rpc.ProductRPC
	voucherRPC rpc.VoucherRPC
	paymentAPI httpi.PaymentAPI
}

func NewApplication(clientRPC rpc.ClientRPC, orderRPC rpc.OrderRPC, productRPC rpc.ProductRPC, voucherRPC rpc.VoucherRPC, paymentAPI httpi.PaymentAPI) Application {
	return application{
		clientRPC:  clientRPC,
		orderRPC:   orderRPC,
		productRPC: productRPC,
		voucherRPC: voucherRPC,
		paymentAPI: paymentAPI,
	}
}

// ========== Client ==========

func (app application) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	l.Infof("GetClientByIDApp: ", " | ", cpf)
	c, err := app.clientRPC.GetClientByCPF(cpf)

	if err != nil {
		l.Errorf("GetClientByIDApp error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByIDApp output: ", " | ", nil)
		return nil, nil
	}

	l.Infof("GetClientByIDApp output: ", " | ", ps.MarshalString(c))
	return c, err
}

func (app application) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {
	l.Infof("SaveClientApp: ", " | ", ps.MarshalString(client))
	clientWithCpf, err := app.GetClientByCPF(client.CPF)

	if err != nil {
		l.Errorf("SaveClientApp error: ", " | ", err)
		return nil, err
	}

	if clientWithCpf != nil {
		l.Infof("SaveClientApp output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because this cpf is already in use")
	}

	c, err := app.clientRPC.SaveClient(client)

	if err != nil {
		l.Errorf("SaveClientApp error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("SaveClientApp output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because it's null")
	}

	l.Infof("SaveClientApp output: ", " | ", ps.MarshalString(c))
	return c, nil
}

// ========== Order ==========

func (app application) UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrderApp, error) {
	l.Infof("UpdateOrderByIDApp: ", " | ", id, " | ", ps.MarshalString(order))
	oSvc, err := app.orderRPC.UpdateOrderByID(id, order)

	if err != nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	if oSvc == nil {
		l.Infof("UpdateOrderByIDApp output: ", " | ", nil)
		return nil, errors.New("order is null, is not possible to proceed with update order")
	}

	client, err := app.GetClientByCPF(oSvc.ClientUUID)

	if err != nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	if client == nil {
		l.Infof("UpdateOrderByIDApp output: ", " | ", nil)
		return nil, errors.New("client is null, is not possible to proceed with update order")
	}

	voucher, err := app.GetVoucherByID(oSvc.VoucherUUID)

	if err != nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	out := &dto.OutputOrderApp{
		ID:               oSvc.ID,
		Client:           client,
		Voucher:          voucher,
		Status:           oSvc.Status,
		VerificationCode: oSvc.VerificationCode,
		CreatedAt:        oSvc.CreatedAt,
	}

	l.Infof("UpdateOrderByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) GetOrders() ([]dto.OutputOrderApp, error) {
	l.Infof("GetOrdersApp: ", " | ")
	orderList := make([]dto.OutputOrderApp, 0)

	orders, err := app.orderRPC.GetOrders()

	if err != nil {
		l.Errorf("GetOrdersApp error: ", " | ", err)
		return nil, err
	}

	for i := range orders {

		client, err := app.GetClientByCPF(orders[i].ClientUUID)

		if err != nil {
			l.Errorf("GetOrdersApp error: ", " | ", err)
			return nil, err
		}

		orderItemList := make([]dto.OutputOrderItem, 0)

		for oItemIdx := range orders[i].Items {
			if orders[i].Items[oItemIdx].OrderID == orders[i].ID {

				orderItemList = append(orderItemList, orders[i].Items[oItemIdx])
			}
		}

		totalPrice := 0.0

		productList := make([]dto.OutputProduct, 0)

		var voucher = dto.OutputVoucher{}

		if len(orders[i].VoucherUUID) > 0 {

			v, err := app.GetVoucherByID(orders[i].VoucherUUID)

			if err != nil {
				l.Errorf("GetOrdersApp error: ", " | ", err)
				return nil, err
			}

			if v == nil {
				l.Infof("GetOrdersApp output: ", " | ", nil)
				return nil, errors.New("is not possible to save order because this voucher does not exist")
			}

			voucher = *v
		}

		for _, op := range orderItemList {
			if len(op.ProductUUID) > 0 {

				meal, errGetMeal := app.GetProductByCategory("meal")
				if errGetMeal != nil {
					l.Errorf("GetOrdersApp error: ", " | ", errGetMeal)
					return nil, errGetMeal
				}

				for i := range meal {
					if meal[i].UUID == op.ProductUUID {
						productList = append(productList, meal[i])
					}
				}

				drink, errGetDrink := app.GetProductByCategory("drink")
				if errGetDrink != nil {
					l.Errorf("GetOrdersApp error: ", " | ", errGetDrink)
					return nil, errGetDrink
				}

				for i := range drink {
					if drink[i].UUID == op.ProductUUID {
						productList = append(productList, drink[i])
					}
				}

				complement, errGetComplement := app.GetProductByCategory("complement")
				if errGetComplement != nil {
					l.Errorf("GetOrdersApp error: ", " | ", errGetComplement)
					return nil, errGetComplement
				}

				for i := range complement {
					if meal[i].UUID == op.ProductUUID {

						totalPrice = totalPrice + getTotalPrice(op.Quantity, meal[i].Price)
						productList = append(productList, meal[i])
					}
				}
			}
		}

		if voucher.Percentage > 0 {
			totalPrice = calculateDiscountByPercentage(voucher.Percentage, totalPrice)
		}

		order := dto.OutputOrderApp{
			ID: orders[i].ID,
			Client: &dto.OutputClient{
				UUID:      client.UUID,
				Name:      client.Name,
				CPF:       client.CPF,
				Email:     client.Email,
				PhoneNumber: client.PhoneNumber,
				Address:    client.Address,
				CreatedAt: client.CreatedAt,
			},
			Voucher:          &voucher,
			Products:         productList,
			TotalPrice:       totalPrice,
			Status:           orders[i].Status,
			VerificationCode: orders[i].VerificationCode,
			CreatedAt:        orders[i].CreatedAt,
		}

		if strings.ToLower(order.Status) != vo.FinishedStatusKey {
			orderList = append(orderList, order)
		}
	}

	l.Infof("GetOrdersApp output: ", " | ", ps.MarshalString(orderList))
	return orderList, nil
}

func (app application) GetOrderByID(id int64) (*dto.OutputOrderApp, error) {
	l.Infof("GetOrderByIDApp: ", " | ", id)

	orders, err := app.orderRPC.GetOrderByID(id)

	if err != nil {
		l.Errorf("GetOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	if orders == nil {
		l.Infof("GetOrderByIDApp output: ", " | ", nil)
		return nil, nil
	}

	client, err := app.GetClientByCPF(orders.ClientUUID)

	if err != nil {
		l.Errorf("GetOrdersApp error: ", " | ", err)
		return nil, err
	}

	orderItemList := make([]dto.OutputOrderItem, 0)

	for oItemIdx := range orders.Items {
		if orders.Items[oItemIdx].OrderID == orders.ID {

			orderItemList = append(orderItemList, orders.Items[oItemIdx])
		}
	}

	totalPrice := 0.0

	productList := make([]dto.OutputProduct, 0)

	var voucher = dto.OutputVoucher{}

	if len(orders.VoucherUUID) > 0 {

		v, err := app.GetVoucherByID(orders.VoucherUUID)

		if err != nil {
			l.Errorf("GetOrdersApp error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			l.Infof("GetOrdersApp output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because this voucher does not exist")
		}

		voucher = *v
	}

	for _, op := range orderItemList {
		if len(op.ProductUUID) > 0 {

			meal, errGetMeal := app.GetProductByCategory("meal")
			if errGetMeal != nil {
				l.Errorf("GetOrdersApp error: ", " | ", errGetMeal)
				return nil, errGetMeal
			}

			for i := range meal {
				if meal[i].UUID == op.ProductUUID {
					productList = append(productList, meal[i])
				}
			}

			drink, errGetDrink := app.GetProductByCategory("drink")
			if errGetDrink != nil {
				l.Errorf("GetOrdersApp error: ", " | ", errGetDrink)
				return nil, errGetDrink
			}

			for i := range drink {
				if drink[i].UUID == op.ProductUUID {
					productList = append(productList, drink[i])
				}
			}

			complement, errGetComplement := app.GetProductByCategory("complement")
			if errGetComplement != nil {
				l.Errorf("GetOrdersApp error: ", " | ", errGetComplement)
				return nil, errGetComplement
			}

			for i := range complement {
				if meal[i].UUID == op.ProductUUID {

					totalPrice = totalPrice + getTotalPrice(op.Quantity, meal[i].Price)
					productList = append(productList, meal[i])
				}
			}
		}
	}

	if voucher.Percentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Percentage, totalPrice)
	}

	out := dto.OutputOrderApp{
		ID: orders.ID,
		Client: &dto.OutputClient{
			UUID:      client.UUID,
			Name:      client.Name,
			CPF:       client.CPF,
			Email:     client.Email,
			PhoneNumber: client.PhoneNumber,
			Address:    client.Address,
			CreatedAt: client.CreatedAt,
		},
		Voucher:          &voucher,
		Products:         productList,
		TotalPrice:       totalPrice,
		Status:           orders.Status,
		VerificationCode: orders.VerificationCode,
		CreatedAt:        orders.CreatedAt,
	}

	l.Infof("GetOrderByIDApp output: ", " | ", ps.MarshalString(out))
	return &out, nil
}

func (app application) SaveOrder(order dto.RequestOrder) (*dto.OutputOrderApp, error) {
	l.Infof("SaveOrderApp: ", " | ", ps.MarshalString(order))

	c, err := app.GetClientByCPF(order.ClientUUID)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	inputDoPaymentAPI := dto.InputPaymentAPI{
		// Price: 0.0,
		Client: entity.Client{
			ID:   c.UUID,
			Name: c.Name,
			CPF: vo.Cpf{
				Value: c.CPF,
			},
			Email:     c.Email,
			CreatedAt: vo.CreatedAt{
				// Value: c.CreatedAt,
			},
		},
	}

	out, err := app.DoPaymentAPI(app.ctx, inputDoPaymentAPI)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	if out.Error != nil {
		l.Errorf("SaveOrderApp error: ", " | ", out.Error.Message, " | ", out.Error.Code)
		return nil, fmt.Errorf("error to do payment message: %s, code: %s", out.Error.Message, out.Error.Code)
	}

	order.Status = out.PaymentStatus

	o, err := app.orderRPC.SaveOrder(order)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		orderNullErr := errors.New("is not possible to save order because it's null")
		l.Infof("SaveOrderApp output: ", " | ", orderNullErr)
		return nil, orderNullErr
	}

	productList := make([]dto.OutputProduct, 0)

	var voucher = dto.OutputVoucher{}

	if len(order.VoucherUUID) > 0 {

		v, err := app.GetVoucherByID(order.VoucherUUID)

		if err != nil {
			l.Errorf("GetOrdersApp error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			l.Infof("GetOrdersApp output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because this voucher does not exist")
		}

		voucher = *v
	}

	orderItemList := make([]dto.OutputOrderItem, 0)

	for oItemIdx := range order.Items {
		if order.Items[oItemIdx].OrderID == order.ID {

			orderItemList = append(orderItemList, order.Items[oItemIdx])
		}
	}

	totalPrice := 0.0

	for _, op := range orderItemList {
		if len(op.ProductUUID) > 0 {

			meal, errGetMeal := app.GetProductByCategory("meal")
			if errGetMeal != nil {
				l.Errorf("GetOrdersApp error: ", " | ", errGetMeal)
				return nil, errGetMeal
			}

			for i := range meal {
				if meal[i].UUID == op.ProductUUID {
					productList = append(productList, meal[i])
				}
			}

			drink, errGetDrink := app.GetProductByCategory("drink")
			if errGetDrink != nil {
				l.Errorf("GetOrdersApp error: ", " | ", errGetDrink)
				return nil, errGetDrink
			}

			for i := range drink {
				if drink[i].UUID == op.ProductUUID {
					productList = append(productList, drink[i])
				}
			}

			complement, errGetComplement := app.GetProductByCategory("complement")
			if errGetComplement != nil {
				l.Errorf("GetOrdersApp error: ", " | ", errGetComplement)
				return nil, errGetComplement
			}

			for i := range complement {
				if meal[i].UUID == op.ProductUUID {

					totalPrice = totalPrice + getTotalPrice(op.Quantity, meal[i].Price)
					productList = append(productList, meal[i])
				}
			}
		}
	}

	if voucher.Percentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Percentage, totalPrice)
	}

	outApp := dto.OutputOrderApp{
		ID: order.ID,
		Client: &dto.OutputClient{
			UUID:      c.UUID,
			Name:      c.Name,
			CPF:       c.CPF,
			Email:     c.Email,
			PhoneNumber: c.PhoneNumber,
			Address:    c.Address,	
			CreatedAt: c.CreatedAt,
		},
		Voucher:          &voucher,
		Products:         productList,
		TotalPrice:       totalPrice,
		Status:           order.Status,
		VerificationCode: order.VerificationCode,
		CreatedAt:        order.CreatedAt,
	}

	l.Infof("SaveOrderApp output: ", " | ", ps.MarshalString(outApp))
	return &outApp, nil
}

// ========== Product ==========

func (app application) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	l.Infof("SaveProductApp: ", " | ", ps.MarshalString(product))

	pRepo, err := app.productRPC.SaveProduct(product)

	if err != nil {
		l.Errorf("SaveProductApp error: ", " | ", err)
		return nil, err
	}

	if pRepo == nil {
		l.Infof("SaveProductApp output: ", " | ", nil)
		return nil, errors.New("is not possible to save product because it's null")
	}

	out := &dto.OutputProduct{
		UUID:          pRepo.UUID,
		Name:          pRepo.Name,
		Category:      pRepo.Category,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt,
		DeactivatedAt: pRepo.DeactivatedAt,
	}

	l.Infof("SaveProductApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	l.Infof("GetProductByCategoryApp: ", " | ", category)
	productList := make([]dto.OutputProduct, 0)

	products, err := app.productRPC.GetProductByCategory(category)

	if err != nil {
		l.Errorf("GetProductByCategoryApp error: ", " | ", err)
		return nil, err
	}

	if products == nil {
		l.Infof("GetProductByCategoryApp output: ", " | ", nil)
		return nil, nil
	}

	for i := range products {
		product := dto.OutputProduct{
			UUID:          products[i].UUID,
			Name:          products[i].Name,
			Category:      products[i].Category,
			Image:         products[i].Image,
			Description:   products[i].Description,
			Price:         products[i].Price,
			CreatedAt:     products[i].CreatedAt,
			DeactivatedAt: products[i].CreatedAt,
		}
		productList = append(productList, product)
	}

	l.Infof("GetProductByCategoryApp output: ", " | ", productList)
	return productList, nil
}

func (app application) UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	l.Infof("UpdateProductByIDApp: ", " | ", id, " | ", ps.MarshalString(product))

	p, err := app.productRPC.UpdateProductByID(id, product)

	if err != nil {
		l.Errorf("UpdateProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		productNullErr := errors.New("is not possible to save product because it's null")
		l.Errorf("UpdateProductByIDApp output: ", " | ", productNullErr)
		return nil, productNullErr
	}

	out := &dto.OutputProduct{
		UUID:          p.UUID,
		Name:          p.Name,
		Category:      p.Category,
		Image:         p.Image,
		Description:   p.Description,
		Price:         p.Price,
		CreatedAt:     p.CreatedAt,
		DeactivatedAt: p.DeactivatedAt,
	}

	l.Infof("UpdateProductByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) DeleteProductByID(id string) error {
	l.Infof("DeleteProductByIDApp: ", " | ", id)

	l.Infof("DeleteProductByIDApp output: ", " | ", nil)
	return app.productRPC.DeleteProductByID(id)
}

// ========== Voucher ==========

func (app application) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	l.Infof("SaveVoucherApp: ", " | ", ps.MarshalString(voucher))

	rVoucher, err := app.voucherRPC.SaveVoucher(voucher)

	if err != nil {
		l.Errorf("SaveVoucherApp error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		voucherNullErr := errors.New("is not possible to save voucher because it's null")
		l.Errorf("SaveVoucherApp output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	vOut := dto.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Percentage: rVoucher.Percentage,
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
	}

	l.Infof("SaveVoucherApp output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func (app application) GetVoucherByID(id string) (*dto.OutputVoucher, error) {
	l.Infof("GetVoucherByIDApp: ", " | ", id)

	rVoucher, err := app.voucherRPC.GetVoucherByID(id)

	if err != nil {
		l.Errorf("GetVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		voucherNotFoundErr := fmt.Errorf("voucher not found with the %d id", id)
		l.Errorf("GetVoucherByIDApp output: ", " | ", voucherNotFoundErr)
		return nil, voucherNotFoundErr
	}

	vOut := dto.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Percentage: rVoucher.Percentage,
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
	}

	l.Infof("GetVoucherByIDApp output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func (app application) UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	l.Infof("UpdateVoucherByIDApp: ", " | ", id, " | ", ps.MarshalString(voucher))

	rVoucher, err := app.voucherRPC.UpdateVoucherByID(id, voucher)

	if err != nil {
		l.Errorf("UpdateVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		voucherNullErr := errors.New("is not possible to update voucher because it's null")
		l.Infof("UpdateVoucherByIDApp output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	vOut := dto.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Percentage: rVoucher.Percentage,
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
	}

	l.Infof("UpdateVoucherByIDApp output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func getTotalPrice(quantity int64, productPrice float64) float64 {
	return productPrice * float64(quantity)
}

func calculateDiscountByPercentage(percentage int64, value float64) float64 {
	if percentage == 0 {
		return value
	}
	return value - (value * (float64(percentage) / 100))
}

func (app application) DoPaymentAPI(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error) {
	return app.paymentAPI.DoPayment(ctx, input)
}
