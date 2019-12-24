package models

type BaseResp struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"msg"`
}

type ChatReq struct {
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
	ChatType int    `json:"type"`
	ChatData []byte `json:"chat_data"`
}

type ChatRsp struct {
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
	ChatType int    `json:"type"`
	ChatData []byte `json:"chat_data"`
}

type ChangeStatusReq struct {
	UserID string `json:"user_id"`
	Status int    `json:"status"`
}
