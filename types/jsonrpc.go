package types

type JsonrpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Result  string `json:"result"`
}
