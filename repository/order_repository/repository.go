package order_repository

import "h8-assignment-2/entity"

type Repository interface {
	CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) error
	ReadOrders() ([]OrderItemMapped, error)
}
