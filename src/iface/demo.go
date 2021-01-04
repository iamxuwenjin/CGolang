package iface

import (
	"fmt"
	"io"
)

type DataWriter interface {
	WriterData(data interface{}) error
	CanWrite(data int) bool
}

type Files struct{}

func (d *Files) WriterData(data interface{}) error {
	fmt.Println("WriterData: ", data)
	return nil
}

func (d *Files) CanWrite(data int) bool {
	if data%2 == 0 {
		return true
	} else {
		return false
	}
}

type Socket struct{}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

func UsingWriter(write io.Writer) {
	fmt.Println("this is writer")
	_, _ = write.Write(nil)
}

func UsingClose(close io.Closer) {
	fmt.Println("this is closer")
	_ = close.Close()
}

type Service interface {
	Start()
	Log(string)
}

type Logger struct{}

func (g *Logger) Log(str string) {
	fmt.Println(str)
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {
	fmt.Println("Game start")

}

func GetType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")

	}
}

type MyString []string

func (m MyString) Len() int {
	return len(m)
}

func (m MyString) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyString) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
