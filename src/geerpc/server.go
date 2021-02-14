package geerpc

import (
	"encoding/json"
	"geerpc/src/geerpc/codec"
	"io"
	"log"
	"net"
	"sync"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodeType    codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType:    codec.GobType,
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error", err)
			return
		}
		go server.ServerConn(conn)
	}
}

func (server *Server) ServerConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc Server: option error", err)
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc Server: invalid magic numver %x", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodeType]
	if f == nil {
		log.Printf("rpc Server: invalid codec type %s", opt.CodeType)
		return
	}
	server.serverCodec(f(conn))
}
func (server *Server) serverCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break // it's not possible to recover, so close the connection
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()
}

// invalidRequest is a placeholder for response argv when error occurs
var invalidRequest = struct{}{}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}
