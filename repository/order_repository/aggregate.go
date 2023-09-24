package order_repository

import (
	"h8-assignment-2/entity"
)

type OrderItemMapped struct {
	Order entity.Order
	Items []entity.Item
}

type OrderItem struct {
	Order entity.Order
	Item  entity.Item
}

func (oim *OrderItemMapped) HandleMappingOrderWithItems(orderItem []OrderItem) []OrderItemMapped {

	ordersItemsMapped := []OrderItemMapped{}

	for _, eachOrderItem := range orderItem {

		isOrderExist := false

		for i := range ordersItemsMapped {
			if eachOrderItem.Order.OrderId == ordersItemsMapped[i].Order.OrderId {
				isOrderExist = true
				ordersItemsMapped[i].Items = append(ordersItemsMapped[i].Items, eachOrderItem.Item)
				break
			}
		}

		if !isOrderExist {

			orderItemMapped := OrderItemMapped{
				Order: eachOrderItem.Order,
			}

			orderItemMapped.Items = append(orderItemMapped.Items, eachOrderItem.Item)

			ordersItemsMapped = append(ordersItemsMapped, orderItemMapped)
		}

	}

	return ordersItemsMapped
}
