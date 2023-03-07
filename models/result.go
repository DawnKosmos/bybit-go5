package models

type ReturnCode interface {
	Return() (code int, msg string)
}

type Response[T any] struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  T      `json:"result"`
	Time    int64  `json:"time"`
}

func (r *Response[T]) Return() (code int, msg string) {
	return r.RetCode, r.RetMsg
}

type ResponseBatch[T any] struct {
	RetCode    int    `json:"retCode"`
	RetMsg     string `json:"retMsg"`
	Result     T      `json:"result"`
	RetExtInfo struct {
		List []struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"list"`
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

func (r *ResponseBatch[T]) Return() (code int, msg string) {
	msg = r.RetMsg
	for _, v := range r.RetExtInfo.List {
		msg += "|" + v.Msg
	}
	return r.RetCode, msg
}

type EmptyResponse struct {
}
