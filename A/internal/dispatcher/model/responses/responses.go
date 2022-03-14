package responses

type (
	GetOrder struct {
		ID int `json:"order_id"`
		General
	}

	General struct {
		Description string `json:"description"`
		Status      int    `json:"http_status"`
		Code        int    `json:"code"`
	}
)
