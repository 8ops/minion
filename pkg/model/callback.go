package model

type CallbackReq struct {
	Batch []*Status `json:"batch"`
}

type Status struct {
	TaskId  string     `json:"taskId"`
	Code    StatusCode `json:"code"`
	Console string     `json:"console"`
}

type StatusCode int
