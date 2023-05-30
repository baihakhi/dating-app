package response

type (
	MapResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

const (
	SUCCESS = "success"
	CREATED = "created"
)
