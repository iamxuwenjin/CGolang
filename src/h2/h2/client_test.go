package h2

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
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

// 验证client 单个HTTP连接的可复用性
func TestClientReuse(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 强制使用h2
		ForceAttemptHTTP2: true,
		// 允许最大连接数
		MaxIdleConns: 1,
		// 允许最大空闲连接数
		MaxConnsPerHost: 1,
		// 每个host允许最大空闲连接数
		MaxIdleConnsPerHost: 1,
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
	_ = resp.Body.Close()

}

// 验证HTTP2的fully multiplexed 完全多路复用
func TestFullyMux(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 强制使用h2
		ForceAttemptHTTP2: true,
		// 允许最大连接数
		MaxIdleConns: 1,
		// 允许最大空闲连接数
		MaxConnsPerHost: 1,
		// 每个host允许最大空闲连接数
		MaxIdleConnsPerHost: 1,
		// 单个连接允许最大空闲时间
		IdleConnTimeout: 2 * time.Minute,
	}
	client := &http.Client{
		Transport: tr,
	}
	// 率先调用响应延迟的接口，HTTP2不存在线头阻塞，无需等待delay_resp接口响应
	go func() {
		defer wg.Done()
		resp, err := client.Get("https://localhost:8080/delay_resp")
		if err != nil {
			fmt.Println(err)
			return
		}
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(content))
		_ = resp.Body.Close()
	}()
	time.Sleep(time.Second)
	// 率先响应protocol_info接口
	go func() {
		defer wg.Done()
		resp, err := client.Get("https://localhost:8080/protocol_info")
		if err != nil {
			fmt.Println(err)
			return
		}
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(content))
		_ = resp.Body.Close()
	}()
	wg.Wait()
}

// 验证HTTP1.x的对头阻塞
func TestHTTPHeadBlock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 强制使用h2
		ForceAttemptHTTP2: true,
		// 允许最大连接数
		MaxIdleConns: 1,
		// 允许最大空闲连接数
		MaxConnsPerHost: 1,
		// 每个host允许最大空闲连接数
		MaxIdleConnsPerHost: 1,
		// 单个连接允许最大空闲时间
		IdleConnTimeout: 2 * time.Minute,
	}
	client := &http.Client{
		Transport: tr,
	}
	// 率先调用响应延迟的接口，HTTP1.x存在线头阻塞，需等待delay_resp接口响应
	go func() {
		defer wg.Done()
		resp, err := client.Get("http://localhost:8081/delay_resp")
		if err != nil {
			fmt.Println(err)
			return
		}
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(content))
		_ = resp.Body.Close()
	}()
	time.Sleep(time.Second)
	go func() {
		defer wg.Done()
		resp, err := client.Get("http://localhost:8081/protocol_info")
		if err != nil {
			fmt.Println(err)
			return
		}
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(content))
		_ = resp.Body.Close()
	}()
	wg.Wait()
}

func TestServerPush(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 强制使用h2
		ForceAttemptHTTP2: true,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("https://localhost:8080/server_push")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	// 很遗憾我们没有收到，来自服务端的推送，原因是目前Go http client 不支持服务端推送
	// 我们可以在Chrome验证这个问题
	fmt.Println(string(content))
	_ = resp.Body.Close()
}
