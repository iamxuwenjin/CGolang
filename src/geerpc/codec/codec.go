package codec

import "io"

type Header struct {
	ServiceMethod string // 服务名和方法名，通常与 Go 语言中的结构体和方法相映射
	Seq           uint64 // 请求的序号
	error         string
}

type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(interface{}) error
	Write(header *Header, args interface{}) error
}

type NewCodecFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
