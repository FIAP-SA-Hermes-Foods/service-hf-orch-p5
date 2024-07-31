package http

import (
	"log"
	"testing"
)

// go test -v -count=1 -failfast -run ^TestValidRoute$
func TestValidRoute(t *testing.T) {
	v, err, m := validRoute("get hermes_foods/order", "hermes_foods/order", "GET")
	log.Printf("isValid true? -> %v | %v | %v", v, err, m)

	v, err, m = validRoute("get hermes_foods/order", "hermes_foods/order/some", "GET")
	log.Printf("isValid false? -> %v | %v | %v", v, err, m)

	v, err, m = validRoute("get hermes_foods/order/{id}", "hermes_foods/order/1", "GET")
	log.Printf("isValid true? -> %v | %v | %v", v, err, m)

	v, err, m = validRoute("get hermes_foods/order/{id}", "hermes_foods/order/1so", "GET")
	log.Printf("isValid false? -> %v | %v | %v", v, err, m)
}
