package order_pg

import (
	"database/sql"
	"h8-assignment-2/entity"
	"h8-assignment-2/repository/order_repository"
)

type orderPG struct {
	db *sql.DB
}

const (
	createOrderQuery = `
		INSERT INTO "orders"
		("ordered_at", "customer_name")
		VALUES ($1, $2)

		RETURNING "order_id"
	`

	createItemQuery = `
		INSERT INTO "items"
		("item_code", "description", "quantity", "order_id")
		VALUES ($1, $2, $3, $4)
	`

	getOrdersWithItemsQuery = `
		SELECT "o"."order_id", "o"."customer_name", "o"."ordered_at", "o"."created_at", "o"."updated_at",
		"i"."item_id", "i"."item_code", "i"."quantity", "i"."description", "i"."order_id", "i"."created_at", "i"."updated_at"
		from "orders" as "o"
		LEFT JOIN "items" as "i" ON "o"."order_id" = "i"."order_id"
		ORDER BY "o"."order_id" ASC
	`
)

func NewOrderPG(db *sql.DB) order_repository.Repository {
	return &orderPG{db: db}
}

func (orderPG *orderPG) ReadOrders() ([]order_repository.OrderItemMapped, error) {
	rows, err := orderPG.db.Query(getOrdersWithItemsQuery)

	if err != nil {
		return nil, err
	}

	orderItems := []order_repository.OrderItem{}

	for rows.Next() {
		var orderItem order_repository.OrderItem

		err = rows.Scan(
			&orderItem.Order.OrderId, &orderItem.Order.CustomerName, &orderItem.Order.OrderedAt, &orderItem.Order.CreatedAt, &orderItem.Order.UpdatedAt,
			&orderItem.Item.ItemId, &orderItem.Item.ItemCode, &orderItem.Item.Quantity, &orderItem.Item.Description, &orderItem.Item.OrderId, &orderItem.Item.CreatedAt, &orderItem.Item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}

	var result order_repository.OrderItemMapped

	return result.HandleMappingOrderWithItems(orderItems), nil

}

func (orderPG *orderPG) CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) error {

	tx, err := orderPG.db.Begin()

	if err != nil {
		return err
	}

	var orderId int

	orderRow := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)

	err = orderRow.Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, eachItem := range itemPayload {
		// ("item_code", "description", "quantity", "order_id")
		_, err := tx.Exec(createItemQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, orderId)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
