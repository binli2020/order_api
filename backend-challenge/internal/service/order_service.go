package service

// OrderService defines behavior related to orders.
type OrderService interface {
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}
