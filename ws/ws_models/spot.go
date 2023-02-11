package ws_models

type PongSpot struct {
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	ConnId  string `json:"conn_id"`
	Op      string `json:"op"`
}
