package structure

// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))

func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	funcList := eventByName[name]

	// 在切片中添加函数
	funcList = append(funcList, callback)

	// 重新存储世界列表
	eventByName[name] = funcList

	//eventByName[name] = append(eventByName[name], callback)
}

func CallEvent(name string, param interface{}) {
	funcList := eventByName[name]

	for _, callback := range funcList {
		callback(param)
	}
}
