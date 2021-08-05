package module

type ProduceMsgReq struct {
	Msg string `json:"msg"`
}

type ProduceMsgResp struct {
	Cost int64 `json:"cost"`
}
