package types

// SMSRequest sdkjfn
type SMSRequest struct {
	Phone string `json:"numero_destino"`
	Msg   string `json:"mensagem"`
}
