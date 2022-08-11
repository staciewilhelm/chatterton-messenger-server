package domain

// REST API Response struct
type APIResponse struct {
	Response APIMessageCode `json:"response"`
	Data     interface{}    `json:"data,omitempty"`
}

type APIMessageCode struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
