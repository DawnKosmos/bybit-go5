package ws_models

/*
position
execution
order
wallet
greeks
*/

type PongPrivate struct {
	ReqId  string   `json:"req_id"`
	Op     string   `json:"op"`
	Args   []string `json:"args"`
	ConnId string   `json:"conn_id"`
}
