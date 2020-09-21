package models

type General interface {
	String() string
}

// 单例工厂
var Singletonfactory = make(map[string]func() General)

// 多例工厂
var Multitonfactory = make(map[string]func() []General)

// 注册对应的工作台

func RegisterSingleton(name string, workbench func() General) {
	Singletonfactory[name] = workbench
}

func RegisterMultiton(name string, workbench func() []General) {
	Multitonfactory[name] = workbench
}

// 工作台生成单例
func Create(name string) General {
	if workbench, ok := Singletonfactory[name]; ok {
		return workbench()
	}
	panic("No workbench found in factory.")
}

// 工作台生成多例
func CreateBatch(name string) []General {
	if workbench, ok := Multitonfactory[name]; ok {
		return workbench()
	} else {
		panic("No workbench found in factory.")
	}
}
