package starter

type IStarter interface {
	// 1. 系统启动,初始化一些资源
	Init()
	// 2. 系统资源安装
	Setup()
	// 3. 启动基础资源
	Start()
	// 启动器是否阻塞
	StartBlocking() bool
	// 4. 资源停止和销毁
	Stop()
}

// 启动器注册器
type starterRegister struct {
	starters []IStarter
}

// 注册启动器
func (r *starterRegister) Register(s IStarter) {
	r.starters = append(r.starters, s)
}

func (r *starterRegister) AllStarters() []IStarter {
	return r.starters
}

var StarterRegister *starterRegister = new(starterRegister)

func Register(s IStarter) {
	StarterRegister.Register(s)
}

type BaseStarter struct{}

func (b *BaseStarter) Init() {}

func (b *BaseStarter) Setup() {}

func (b *BaseStarter) Start() {}

func (b *BaseStarter) StartBlocking() bool {
	return false
}

func (b *BaseStarter) Stop() {}
