package main

//func main() {
//	// 实例化file
//	f := new(iface.Files)
//	// 声明一个DataWriter的接口
//	var writer iface.DataWriter
//	// 将接口赋值f，也就是*file类型
//	writer = f
//	_ = writer.WriterData("data")
//
//	s := new(iface.Socket)
//	iface.UsingWriter(s)
//	iface.UsingClose(s)
//
//	var service iface.Service = new(iface.GameService)
//	service.Start()
//	service.Log("Game end")
//
//	var x interface{}
//	x = "hello"
//	value := x.(string)
//	fmt.Println(value)
//	iface.GetType("1111")
//
//	names := iface.MyString{
//		"3. Triple Kill",
//		"5. Penta Kill",
//		"2. Double Kill",
//		"4. Quadra Kill",
//		"1. First Blood",
//	}
//	sort.Strings(names)
//
//	for _, v :=range names{
//		fmt.Println(v)
//	}
//
//}
