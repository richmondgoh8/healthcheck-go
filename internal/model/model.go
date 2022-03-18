package model

const (
	SvcSuccessMsg = "Service is Up & Running"
)

// swagger:response HealthResponse
type Response struct {
	Status int `json:"status"`
	//swagger:ignore
	Data interface{} `json:"data"`
}

type RandomResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

type PredictionChannel struct {
	Name string
	Age  int
}

func GenerateResponse(code int, data interface{}) Response {
	return Response{
		Status: code,
		Data:   data,
	}
}
