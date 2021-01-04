package main

//func main() {
//	point := structure.Point{
//		X: 4,
//		Y: 5,
//	}
//	fmt.Println(point.Y)
//
//	player := new(structure.Player)
//	player.Hp = 66
//	player.Mp = 88
//	player.Name = "jojo"
//
//	fmt.Println(player)
//
//	res := structure.Mimi("jojo")
//	fmt.Println(res)
//
//	car := structure.Car{
//		Wheel: structure.Wheel{
//			Size: 16,
//		},
//		Engine: structure.Engine{
//			Power: 240,
//			Type:  8,
//		},
//	}
//	fmt.Printf("%v\n", car)
//
//	var head = new(structure.Node)
//	head.Data = 0
//	//var node = new(structure.Node)
//	//node.Data = 2
//	//head.Next = node
//
//	var tail *structure.Node
//
//	tail = head
//	//for i := 1; i < 10; i++{
//	//	var node = structure.Node{Data: i}
//	//	node.Next = tail
//	//	tail = &node
//	//}
//
//	for i := 1; i < 10; i++ {
//		var node = structure.Node{
//			Data: i,
//		}
//		(*tail).Next = &node
//		tail = &node
//	}
//	// 首尾相连的链表
//	//(*tail).Next = head
//
//	structure.ShowNode(head)
//
//	var preNode = new(structure.PreNode)
//	preNode.Data = 0
//
//	newTail := preNode
//
//	for i := 1; i < 10; i++ {
//		node := structure.PreNode{
//			Data: i,
//		}
//		node.Pre = newTail
//		(*newTail).Next = &node
//		newTail = &node
//	}
//	structure.ShowNode(head)
//
//	data := []byte("c语言中文网")
//	rd := bytes.NewReader(data)
//	r := bufio.NewReader(rd)
//	var buf [128]byte
//	n, err := r.Read(buf[:])
//	fmt.Println(string(buf[:n]), err)
//
//
//}
