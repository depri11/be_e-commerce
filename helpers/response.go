package helpers

type Response struct {
	Status  int         `json:"status" default:"200"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Other   interface{} `json:"other"`
}
