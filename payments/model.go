package payments

type Payment struct {
	Amount float32 `json:"amount"`
	Type   string  `json:"type"`
}
