package order_pg

import (
	"database/sql"
	"errors"
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
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

	getOrderById = `
		SELECT "order_id", "customer_name", "ordered_at", "created_at", "updated_at" FROM "orders"
		WHERE "order_id" = $1

	`

	updateOrderByIdQuery = `
		UPDATE "orders"
		SET "ordered_at" = $2,
		"customer_name" = $3
		WHERE "order_id" = $1
	`

	updateItemByCodeQuery = `
		UPDATE "items"
		SET "description" = $2,
		"quantity" = $3
		WHERE "item_code" = $1
	`
	deleteOrderQuery = `
	DELETE FROM "orders" WHERE "order_id" = $1
	`
)

func NewOrderPG(db *sql.DB) order_repository.Repository {
	return &orderPG{db: db}
}

//Read Order by ID
func (orderPG *orderPG) ReadOrderById(orderId int) (*entity.Order, errs.Error) {
	row := orderPG.db.QueryRow(getOrderById, orderId)

	var order entity.Order

	err := row.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("order not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &order, nil
}

//Fungsi Update order
func (orderPG *orderPG) UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) errs.Error {
	tx, err := orderPG.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(updateOrderByIdQuery, orderPayload.OrderId, orderPayload.OrderedAt, orderPayload.CustomerName)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	for _, eachItem := range itemPayload {
		_, err = tx.Exec(updateItemByCodeQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity)

		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError("something went wrong")
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

//Fungsi Read Order
func (orderPG *orderPG) ReadOrders() ([]order_repository.OrderItemMapped, errs.Error) {
	rows, err := orderPG.db.Query(getOrdersWithItemsQuery)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	orderItems := []order_repository.OrderItem{}

	for rows.Next() {
		var orderItem order_repository.OrderItem

		err = rows.Scan(
			&orderItem.Order.OrderId, &orderItem.Order.CustomerName, &orderItem.Order.OrderedAt, &orderItem.Order.CreatedAt, &orderItem.Order.UpdatedAt,
			&orderItem.Item.ItemId, &orderItem.Item.ItemCode, &orderItem.Item.Quantity, &orderItem.Item.Description, &orderItem.Item.OrderId, &orderItem.Item.CreatedAt, &orderItem.Item.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		orderItems = append(orderItems, orderItem)
	}

	var result order_repository.OrderItemMapped

	return result.HandleMappingOrderWithItems(orderItems), nil

}

//fungsi Create Order
func (orderPG *orderPG) CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) errs.Error {

	tx, err := orderPG.db.Begin()

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	var orderId int

	orderRow := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)

	err = orderRow.Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	for _, eachItem := range itemPayload {
		_, err := tx.Exec(createItemQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, orderId)

		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError("something went wrong")
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}


func (orderPG *orderPG) DeleteOrder(orderId int) errs.Error {
	_, err := orderPG.db.Exec(deleteOrderQuery, orderId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}