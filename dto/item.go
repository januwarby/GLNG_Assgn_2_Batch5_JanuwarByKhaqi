package dto

import "time"

type NewItemRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type OrderWithItems struct {
	OrderId      int               `json:"orderId"`
	CustomerName string            `json:"customerName"`
	OrderedAt    time.Time         `json:"orderedAt"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Items        []GetItemResponse `json:"items"`
}

type GetOrdersResponse struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       []OrderWithItems `json:"data"`
}
