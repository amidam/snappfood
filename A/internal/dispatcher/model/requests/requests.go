package requests

type (
	GetOrder struct {
		ID    int    `json:"order_id"`
		Price int    `json:"price"`
		Title string `json:"title"`
	}
)
