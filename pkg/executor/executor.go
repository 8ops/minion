package executor

type IExecutor interface {
	Name() string   //执行者角色
	Period() int    //间隔时间
	Execute() error //执行
	Release()       //释放
}
