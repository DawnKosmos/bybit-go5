package ws_models

type PongLinear struct {
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	ConnId  string `json:"conn_id"`
	ReqId   string `json:"req_id"`
	Op      string `json:"op"`
}
