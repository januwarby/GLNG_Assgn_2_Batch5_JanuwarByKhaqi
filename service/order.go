package service

import (
	"h8-assignment-2/dto"
	"h8-assignment-2/entity"
	"h8-assignment-2/repository/order_repository"
	"net/http"
)

type orderService struct {
	OrderRepo order_repository.Repository
}

type OrderService interface {
	CreateOrder(newOrderRequest dto.NewOrderRequest) error
	GetOrders() (*dto.GetOrdersResponse, error)
}

func NewOrderService(orderRepo order_repository.Repository) OrderService {
	return &orderService{
		OrderRepo: orderRepo,
	}
}

func (os *orderService) GetOrders() (*dto.GetOrdersResponse, error) {
	orders, err := os.OrderRepo.ReadOrders()

	if err != nil {
		return nil, err
	}

	orderResult := []dto.OrderWithItems{}

	for _, eachOrder := range orders {
		order := dto.OrderWithItems{
			OrderId:      eachOrder.Order.OrderId,
			CustomerName: eachOrder.Order.CustomerName,
			OrderedAt:    eachOrder.Order.OrderedAt,
			CreatedAt:    eachOrder.Order.CreatedAt,
			UpdatedAt:    eachOrder.Order.UpdatedAt,
			Items:        []dto.GetItemResponse{},
		}

		for _, eachItem := range eachOrder.Items {
			item := dto.GetItemResponse{
				ItemId:      eachItem.ItemId,
				ItemCode:    eachItem.ItemCode,
				Quantity:    eachItem.Quantity,
				Description: eachItem.Description,
				OrderId:     eachItem.OrderId,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
			}

			order.Items = append(order.Items, item)
		}

		orderResult = append(orderResult, order)

	}

	response := dto.GetOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "orders successfully fetched",
		Data:       orderResult,
	}

	return &response, nil
}

func (os *orderService) CreateOrder(newOrderRequest dto.NewOrderRequest) error {

	orderPayload := entity.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	itemPayload := []entity.Item{}

	for _, eachItem := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}

		itemPayload = append(itemPayload, item)

	}

	err := os.OrderRepo.CreateOrder(orderPayload, itemPayload)

	if err != nil {
		return err
	}

	return nil

}
