package h2

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Log(format string, v ...interface{}) {
	log.Printf("[Server] %s", fmt.Sprintf(format, v...))

}

// 查看HTTP协议版本
func ProtocolInfo(w http.ResponseWriter, r *http.Request) {
	Log("%s %s", r.Method, r.URL.Path)
	_, _ = fmt.Fprintf(w, "HTTP协议版本：%s\n", r.Proto)
}

// 延迟2S响应, fake响应慢的接口
func DelayResp(w http.ResponseWriter, r *http.Request) {
	Log("%s %s", r.Method, r.URL.Path)
	time.Sleep(2 * time.Second)
	_, _ = fmt.Fprintf(w, "延迟响应")
}

func StartServer() {
	Path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/protocol_info", ProtocolInfo)
	http.HandleFunc("/delay_resp", DelayResp)

	// net/http包默认支持http2的，而HTTP/2强制使用TLS的，所以在使用的时候必须指定证书
	// 这里使用自签证书支持h2服务端
	server := &http.Server{
		Addr: ":8080",
		// 长连接空闲保持时间
		IdleTimeout: 2 * time.Minute,
	}

	go func() {
		log.Println("HTTP2 server has started")
		log.Fatal(server.ListenAndServeTLS(Path+"/src/h2/SSLKey/server.crt", Path+"/src/h2/SSLKey/server.key"))
	}()

	// 使用默认配置，启动一个HTTP服务器，用于测试1.*版本特性
	log.Println("HTTP1.x server has started")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
