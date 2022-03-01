package dto

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
