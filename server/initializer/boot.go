/*
	管理 应用加载启动生命周期
*/
package initializer

import (
	starter "server/initializer/starter"
)

type BootApplication struct {
}

func New() *BootApplication {
	b := &BootApplication{}
	return b
}

func (b *BootApplication) Start() {
	b.register()
	// 1、初始化 starter
	b.init()
	// 2、安装 starter
	b.setup()
	// 3、启动 starter
	b.start()
}

func (b *BootApplication) register() {
	starter.StarterRegister.Register(&starter.ConfigStarter{})
	starter.StarterRegister.Register(&starter.DBStarter{})
}

func (b *BootApplication) init() {
	for _, starter := range starter.StarterRegister.AllStarters() {
		starter.Init()
	}
}

func (b *BootApplication) setup() {
	for _, starter := range starter.StarterRegister.AllStarters() {
		starter.Setup()
	}
}

func (b *BootApplication) start() {
	allStarters := starter.StarterRegister.AllStarters()
	for index, s := range starter.StarterRegister.AllStarters() {
		if s.StartBlocking() {
			// 如果是最后一个 starter，是可以阻塞的，直接启动并阻塞
			if index+1 == len(allStarters) {
				s.Start()
				// 如果不是，使用goroutine，防止阻塞后续 starter
			} else {
				go s.Start()
			}
		} else {
			s.Start()

		}
	}
}
