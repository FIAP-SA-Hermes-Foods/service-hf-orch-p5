package http

import (
	"service-hf-orch-p5/internal/core/application"
	"net/http"
)

type Handler interface {
	HandlerClient(rw http.ResponseWriter, req *http.Request)
	HandlerOrder(rw http.ResponseWriter, req *http.Request)
	HandlerProduct(rw http.ResponseWriter, req *http.Request)
	HandlerVoucher(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	app application.Application
}

func NewHandler(app application.Application) Handler {
	return handler{app: app}
}

func (h handler) HandlerClient(rw http.ResponseWriter, req *http.Request) {

	var routesVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/client/{cpf}": h.getClientByCPF,
		"post hermes_foods/client":      h.saveClient,
	}

	handler, err := router(req.Method, req.URL.Path, routesVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handler) HandlerProduct(rw http.ResponseWriter, req *http.Request) {

	var routesVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/product":         h.getProductByCategory,
		"post hermes_foods/product":        h.saveProduct,
		"put hermes_foods/product/{id}":    h.UpdateProductByUUID,
		"delete hermes_foods/product/{id}": h.deleteProductByUUID,
	}

	handler, err := router(req.Method, req.URL.Path, routesVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handler) HandlerOrder(rw http.ResponseWriter, req *http.Request) {

	var routesVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/order":        h.getOrders,
		"get hermes_foods/order/{id}":   h.getOrderByID,
		"post hermes_foods/order":       h.saveOrder,
		"patch hermes_foods/order/{id}": h.updateOrderByID,
	}

	handler, err := router(req.Method, req.URL.Path, routesVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handler) HandlerVoucher(rw http.ResponseWriter, req *http.Request) {

	var routesVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/voucher/{id}": h.getVoucherByID,
		"post hermes_foods/voucher":     h.saveVoucher,
		"put hermes_foods/voucher/{id}": h.updateVoucherByID,
	}

	handler, err := router(req.Method, req.URL.Path, routesVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}
