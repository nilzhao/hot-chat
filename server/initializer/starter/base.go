package starter

type IStarter interface {
	// 启动器名称
	Name() string
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

func (s *BaseStarter) Name() string { return "未知" }

func (s *BaseStarter) Init() {}

func (s *BaseStarter) Setup() {}

func (s *BaseStarter) Start() {}

func (s *BaseStarter) StartBlocking() bool {
	return false
}

func (s *BaseStarter) Stop() {}
