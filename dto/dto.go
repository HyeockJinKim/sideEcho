package dto

type ErrorResponse struct {
	Message string `json:"message"`
}

type BuyRequest struct {
	Value uint64 `json:"value"`
}

type BuyResponse struct {
	Value uint64 `json:"value"`
}

type SellRequest struct {
	Value uint64 `json:"value"`
}

type SellResponse struct {
	Value uint64 `json:"value"`
}
