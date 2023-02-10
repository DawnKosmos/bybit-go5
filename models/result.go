package models

type Response[T any] struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  T      `json:"result"`
	Time    int64  `json:"time"`
}

type EmptyResponse struct {
}

func (r *Response[T]) Return() (code int, msg string) {
	return r.RetCode, r.RetMsg
}

type ReturnCode interface {
	Return() (code int, msg string)
}
