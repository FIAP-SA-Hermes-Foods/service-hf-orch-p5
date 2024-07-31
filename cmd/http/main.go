package main

import (
	"context"
	l "service-hf-orch-p5/external/logger"
	httpExt "service-hf-orch-p5/internal/adapters/driver/http"
	clientrpc "service-hf-orch-p5/internal/adapters/driver/rpc/client"
	orderrpc "service-hf-orch-p5/internal/adapters/driver/rpc/order"
	productrpc "service-hf-orch-p5/internal/adapters/driver/rpc/product"
	voucherrpc "service-hf-orch-p5/internal/adapters/driver/rpc/voucher"
	"service-hf-orch-p5/internal/core/application"
	httpH "service-hf-orch-p5/internal/handler/http"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()

	ctx := context.Background()

	urlAPI := fmt.Sprintf("http://%s:%s/%s",
		os.Getenv("MERCADO_PAGO_API_HOST"),
		os.Getenv("MERCADO_PAGO_API_PORT"),
		os.Getenv("MERCADO_PAGO_API_URI"),
	)

	headersAPI := map[string]string{
		"Content-type": "application/json",
	}

	du, err := time.ParseDuration(os.Getenv("MERCADO_PAGO_API_TIMEOUT"))

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	paymentAPI := httpExt.NewMercadoPagoAPI(urlAPI, headersAPI, du)
	clientRPC := clientrpc.NewClientRPC(ctx, os.Getenv("HOST_CLIENT"), os.Getenv("PORT_CLIENT"))
	orderRPC := orderrpc.NewOrderRPC(ctx, os.Getenv("HOST_ORDER"), os.Getenv("PORT_ORDER"))
	productRPC := productrpc.NewProductRPC(ctx, os.Getenv("HOST_PRODUCT"), os.Getenv("PORT_PRODUCT"))
	voucherRPC := voucherrpc.NewVoucherRPC(ctx, os.Getenv("HOST_PRODUCT"), os.Getenv("PORT_PRODUCT"))

	app := application.NewApplication(clientRPC, orderRPC, productRPC, voucherRPC, paymentAPI)

	h := httpH.NewHandler(app)

	router.Handle("/hermes_foods/health", http.StripPrefix("/", httpH.Middleware(h.HealthCheck)))
	router.Handle("/hermes_foods/client/", http.StripPrefix("/", httpH.Middleware(h.HandlerClient)))
	router.Handle("/hermes_foods/client", http.StripPrefix("/", httpH.Middleware(h.HandlerClient)))

	router.Handle("/hermes_foods/order/", http.StripPrefix("/", httpH.Middleware(h.HandlerOrder)))
	router.Handle("/hermes_foods/order", http.StripPrefix("/", httpH.Middleware(h.HandlerOrder)))

	router.Handle("/hermes_foods/product/", http.StripPrefix("/", httpH.Middleware(h.HandlerProduct)))
	router.Handle("/hermes_foods/product", http.StripPrefix("/", httpH.Middleware(h.HandlerProduct)))

	router.Handle("/hermes_foods/voucher/", http.StripPrefix("/", httpH.Middleware(h.HandlerVoucher)))
	router.Handle("/hermes_foods/voucher", http.StripPrefix("/", httpH.Middleware(h.HandlerVoucher)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
