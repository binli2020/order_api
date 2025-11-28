package model

type OrderRequest struct {
	CouponCode string      `json:"couponCode"`
	Items      []OrderItem `json:"items"`
}
