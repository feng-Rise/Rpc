package codec

import (
	"io"
)

type Header struct {
	ServiceMethod string // 服务名+方法
	Seq           uint64 // 请求的序号，用来区分不同的请求。
	Error         string // 是错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中。
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
