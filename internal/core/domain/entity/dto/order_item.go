package dto

type (
	OrderItemDB struct {
		ID          int64   `json:"id,omitempty"`
		OrderID     int64   `json:"orderId,omitempty"`
		ProductUUID string  `json:"productUuid,omitempty"`
		Quantity    int64   `json:"quantity,omitempty"`
		TotalPrice  float64 `json:"totalPrice,omitempty"`
		Discount    float64 `json:"discount,omitempty"`
		CreatedAt   string  ` json:"createdAt,omitempty"`
	}

	RequestOrderItem struct {
		ID          int64   `json:"id,omitempty"`
		OrderID     int64   `json:"orderId,omitempty"`
		ProductUUID string  `json:"productUuid,omitempty"`
		Quantity    int64   `json:"quantity,omitempty"`
		TotalPrice  float64 `json:"totalPrice,omitempty"`
		Discount    float64 `json:"discount,omitempty"`
		CreatedAt   string  `json:"createdAt,omitempty"`
	}

	OutputOrderItem struct {
		ID          int64   `json:"id,omitempty"`
		OrderID     int64   `json:"orderId,omitempty"`
		ProductUUID string  `json:"productUuid,omitempty"`
		Quantity    int64   `json:"quantity,omitempty"`
		TotalPrice  float64 `json:"totalPrice,omitempty"`
		Discount    float64 `json:"discount,omitempty"`
		CreatedAt   string  `json:"createdAt,omitempty"`
	}
)
