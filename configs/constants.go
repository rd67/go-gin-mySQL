package configs

type CommonResponseStruct struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

const (
	DefaultPageSize = 100
)
