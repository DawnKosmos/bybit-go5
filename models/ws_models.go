package models

type WsPublicRequest struct {
	ReqId string   `json:"req_id,omitempty"`
	Op    string   `json:"op"`
	Args  []string `json:"args"`
}

type WsPublicResponse struct {
	ReqId   string `json:"req_id,omitempty"`
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	ConnId  string `json:"conn_id"`
	Op      string `json:"op"`
}

type WsPrivateRequest struct {
	ReqId string `json:"req_id,omitempty"`
	Op    string `json:"op"`
	Args  []any  `json:"args"`
}

type PongPrivate struct {
	ReqId  string   `json:"req_id,omitempty"`
	Op     string   `json:"op"`
	Args   []string `json:"args"`
	ConnId string   `json:"conn_id"`
}

type WsPing struct {
	ReqId string `json:"req_id,omitempty"`
	Op    string `json:"op"`
}

type WsPong struct {
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	ConnId  string `json:"conn_id"`
	Op      string `json:"op"`
}

type Test struct {
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Ts    int64  `json:"ts"`
	Data  struct {
		S   string        `json:"s"`
		B   []interface{} `json:"b"`
		A   []interface{} `json:"a"`
		U   int           `json:"u"`
		Seq int64         `json:"seq"`
	} `json:"data"`
}
