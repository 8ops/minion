package model

type TaskResp struct {
	Code   int     `json:"code"`
	Policy *Policy `json:"policy"`
	Batch  []*Task `json:"batch"`
}

type Policy struct {
	Period       int    `json:"period"`
	Timeout      int    `json:"timeout"`
	TaskPoolSize int    `json:"taskPoolSize"`
	CallbackUrl  string `json:"callbackUrl"`
}

type Task struct {
	TaskId  string `json:"taskId"`
	UtilUrl string `json:"utilUrl"`
	Args    string `json:"args"`
}
