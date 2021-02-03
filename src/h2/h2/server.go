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

func ProtocolInfo(w http.ResponseWriter, r *http.Request) {
	Log("%s %s", r.Method, r.URL.Path)
	_, _ = fmt.Fprintf(w, "HTTP协议版本：%s\n", r.Proto)
}

func StartServer() {
	Path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/protocol_info", ProtocolInfo)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	// net/http包默认支持http2的，而HTTP/2强制使用TLS的，所以在使用的时候必须指定证书
	// 这里使用自签证书支持h2服务端
	server := &http.Server{
		Addr: ":8080",
		// 长连接空闲保持时间
		IdleTimeout: 2 * time.Minute,
	}
	log.Fatal(server.ListenAndServeTLS(Path+"/src/h2/SSLKey/server.crt", Path+"/src/h2/SSLKey/server.key"))
}
