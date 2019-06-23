package server

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"k8s.io/klog"
	"minion/pkg/config"
	"minion/pkg/model"
)

func Startup() {
	ts := getTask()
}

//取任务
func getTask() *model.TaskResp {
	ts := &model.TaskResp{}
	if resp, body, errs := gorequest.New().Get(config.APITask).
		Set("User-Agent", "minion").
		End(); len(errs) != 0 {
		klog.Fatal(errs)
	} else if resp.StatusCode != 200 {
		klog.Errorln("请求返回不正常")
	} else {
		if err := json.Unmarshal([]byte(body), ts); err != nil {
			klog.Errorln(err)
		}
	}
	return ts
}

//执行
func exec() {

}

//汇报
func callback() {

}
