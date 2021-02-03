package h2

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// 展示了如何与自签证书h2服务端建立h2连接
func TestProtocolInfo(t *testing.T) {
	// certificate signed by unknown authority
	//resp, err := http.Get("https://localhost:8080/protocol_info")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 强制使用h2
		ForceAttemptHTTP2: true,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("https://localhost:8080/protocol_info")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
	_ = resp.Body.Close()
}

// 验证client复用性
func TestClientReuse(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 强制使用h2
		ForceAttemptHTTP2: true,
		// 允许最大连接数
		MaxIdleConns: 10,
		// 允许最大空闲连接数
		MaxConnsPerHost: 10,
		// 每个host允许最大空闲连接数
		MaxIdleConnsPerHost: 10,
		// 单个连接允许最大空闲时间
		IdleConnTimeout: 2 * time.Minute,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("https://localhost:8080/protocol_info")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
	_ = resp.Body.Close()

	// 阻塞一分钟，方便命令行查看连接状态，或者在此处debug
	// todo cmd: netstat -an |grep ESTABLISHED |grep 8080
	//time.Sleep(1*time.Minute)

	resp, err = client.Get("https://localhost:8080/protocol_info")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}
